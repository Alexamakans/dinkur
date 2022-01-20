// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package afkdetect contains code to detect if the user has gone AFK or
// returned from AFK.
package afkdetect

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dinkur/dinkur/internal/obs"
	"github.com/iver-wharf/wharf-core/pkg/logger"
)

// Errors specific to AFK-detectors.
var (
	ErrObserverIsNil = errors.New("observer is nil")
)

var afkPollIntervalDur = 3 * time.Second
var afkThresholdDur = 5 * time.Minute
var log = logger.NewScoped("AFK")

// Detector is an AFK-detector.
type Detector interface {
	// Start makes this detector start listening for OS-specific events to then
	// trigger AFK-started or AFK-stopped events.
	//
	// You need to stop the detector to remove any dangling Goroutines.
	StartDetecting() error
	// StopDetecting makes this detector stop listening for OS-specific events by
	// cleaning up its Goroutines and hooks.
	StopDetecting() error

	StartedObs() obs.Observer[Started]
	StoppedObs() obs.Observer[Stopped]
}

// Started contains event data for when user has gone AFK.
type Started struct{}

// Stopped contains event data for when user is no longer AFK (after being AFK).
type Stopped struct {
	AFKSince time.Time
}

type detectorHookRegisterer interface {
	Register(*detector) (detectorHook, error)
}

type detectorHook interface {
	Unregister() error
	Tick() error
}

var detectorHooks []detectorHookRegisterer

// New creates a new AFK-detector.
func New() Detector {
	return &detector{
		startedObs: obs.New[Started](),
		stoppedObs: obs.New[Stopped](),
	}
}

type detector struct {
	isAFKMutex sync.RWMutex
	afkSince   *time.Time
	startedObs obs.Observer[Started]
	stoppedObs obs.Observer[Stopped]

	hooks          []detectorHook
	startStopMutex sync.Mutex
	ticker         *time.Ticker
	tickChanStop   chan struct{}
}

func (d *detector) isAFK() bool {
	d.isAFKMutex.RLock()
	isAFK := d.afkSince != nil
	d.isAFKMutex.RUnlock()
	return isAFK
}

func (d *detector) setIsAFK(isAFK bool) (time.Time, bool) {
	if d.isAFK() == isAFK {
		// no update needed
		return time.Time{}, false
	}
	d.isAFKMutex.Lock()
	defer d.isAFKMutex.Unlock()
	if isAFK {
		now := time.Now()
		d.afkSince = &now
		return now, true
	}
	prevTimeSince := *d.afkSince
	d.afkSince = nil
	return prevTimeSince, true
}

func (d *detector) markAsAFK() {
	if _, changed := d.setIsAFK(true); !changed {
		return
	}
	log.Debug().Message("User is now AFK.")
	d.startedObs.Pub(Started{})
}

func (d *detector) markAsNoLongerAFK() {
	t, changed := d.setIsAFK(false)
	if !changed {
		return
	}
	log.Debug().Message("User is no longer AFK.")
	d.stoppedObs.Pub(Stopped{
		AFKSince: t,
	})
}

func (d *detector) StartDetecting() error {
	if len(detectorHooks) == 0 {
		log.Warn().Message("No AFK-detectors available for this OS.")
		return nil
	}
	d.startStopMutex.Lock()
	defer d.startStopMutex.Unlock()
	d.hooks = nil
	for _, reg := range detectorHooks {
		hook, err := reg.Register(d)
		if err != nil {
			d.StopDetecting()
			return err
		}
		if hook != nil {
			d.hooks = append(d.hooks, hook)
		}
	}
	d.ticker = time.NewTicker(afkPollIntervalDur)
	go d.timerTickListener(d.ticker)
	return nil
}

func (d *detector) StopDetecting() error {
	d.startStopMutex.Lock()
	defer d.startStopMutex.Unlock()
	for _, hook := range d.hooks {
		if err := hook.Unregister(); err != nil {
			log.Error().WithError(err).Messagef("Failed to unregister %T.", hook)
		}
	}
	d.hooks = nil
	if d.ticker != nil {
		d.ticker.Stop()
		d.ticker = nil
	}
	if d.tickChanStop != nil {
		d.tickChanStop <- struct{}{}
		d.tickChanStop = nil
	}
	unsubStartErr := d.startedObs.UnsubAll()
	unsubStopErr := d.stoppedObs.UnsubAll()
	if unsubStartErr != nil && unsubStopErr != nil {
		return fmt.Errorf("unsub all afk-start and stop subs: %w; %v", unsubStartErr, unsubStopErr)
	} else if unsubStartErr != nil {
		return fmt.Errorf("unsub all afk-start subs: %w", unsubStartErr)
	} else if unsubStopErr != nil {
		return fmt.Errorf("unsub all afk-stop subs: %w", unsubStopErr)
	}
	return nil
}

func (d *detector) timerTickListener(ticker *time.Ticker) {
	for {
		select {
		case <-d.tickChanStop:
			ticker.Stop()
			return
		case <-ticker.C:
			for _, hook := range d.hooks {
				if err := hook.Tick(); err != nil {
					log.Warn().WithError(err).
						Messagef("Failed to tick AFK hook %T.", hook)
				}
			}
		}
	}
}

func (d *detector) StartedObs() obs.Observer[Started] {
	return d.startedObs
}

func (d *detector) StoppedObs() obs.Observer[Stopped] {
	return d.stoppedObs
}
