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

package console

import (
	"fmt"
	"log"
	"time"

	"github.com/dinkur/dinkur/pkg/dinkur"
)

const firstWeekday = time.Monday
const lastWeekday = time.Sunday

// FormatDuration returns a formatted time.Duration in the format of
// h:mm:ss.
func FormatDuration(d time.Duration) string {
	negative := d < 0
	if negative {
		d = -d
	}
	var (
		totalSeconds = int(d.Seconds())
		hours        = totalSeconds / 60 / 60
		minutes      = totalSeconds / 60 % 60
		seconds      = totalSeconds % 60
	)
	if negative {
		return fmt.Sprintf("-%d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}

type group interface {
	fmt.Stringer

	new(_g groupable) group

	check(t time.Time) bool
	count() int
	addEntry(entry dinkur.Entry)
	getEntries() []dinkur.Entry
	getGroupBy() groupable
	start() time.Time
}

func newDate(year int, month time.Month, day int) date {
	return date{year, month, day}
}

type date struct {
	year  int
	month time.Month
	day   int
}

func (d date) String() string {
	return fmt.Sprintf("%s-%d", d.month.String()[:3], d.day)
}

func (day) new(t time.Time) groupable {
	return day{t.Day(), int(t.Weekday())}
}

type day struct {
	day     int
	weekDay int
}

func (d day) String() string {
	return fmt.Sprintf("%02d", d.day)
}

func (week) new(t time.Time) groupable {
	year, _week := t.ISOWeek()
	return week{year, _week}
}

type week struct {
	year int
	week int
}

func (w week) String() string {
	return fmt.Sprintf("%02d", w.week)
}

func (month) new(t time.Time) groupable {
	year, _month := t.Year(), t.Month()
	return month{year, _month}
}

type month struct {
	year  int
	month time.Month
}

func (m month) String() string {
	return fmt.Sprintf("%s", m.month.String()[:3])
}

func (year) new(t time.Time) groupable {
	_year := t.Year()
	return year{_year}
}

type year struct {
	year int
}

func (y year) String() string {
	return fmt.Sprintf("%d", y.year)
}

type groupable interface {
	fmt.Stringer
	new(t time.Time) groupable
}

type entryGroup struct {
	groupBy groupable

	entries []dinkur.Entry
}

func (g *entryGroup) new(_g groupable) group {
	return &entryGroup{
		groupBy: _g,
	}
}

func (g *entryGroup) String() string {
	return g.groupBy.String()
}

func (g *entryGroup) addEntry(entry dinkur.Entry) {
	g.entries = append(g.entries, entry)
}

func (g *entryGroup) check(t time.Time) bool {
	return g.groupBy.new(t) == g.groupBy
}

func (g *entryGroup) count() int {
	return len(g.entries)
}

func (g *entryGroup) getEntries() []dinkur.Entry {
	return g.entries
}

func (g *entryGroup) getGroupBy() groupable {
	return g.groupBy
}

func (g *entryGroup) start() time.Time {
	return g.entries[0].Start
}

// groupEntries assumes the slice is already sorted on entry.Start
func groupEntries(g group, entries []dinkur.Entry) []group {
	if len(entries) == 0 {
		return nil
	}
	var groups []group
	for _, t := range entries {
		if !g.check(t.Start) {
			if g.count() > 0 {
				groups = append(groups, g)
			}
			g = g.new(g.getGroupBy().new(t.Start))
		}
		g.addEntry(t)
	}
	if g.count() > 0 {
		groups = append(groups, g)
	}
	return groups
}

type entrySum struct {
	start    time.Time
	end      *time.Time
	duration time.Duration
}

func sumEntries(entries []dinkur.Entry) entrySum {
	if len(entries) == 0 {
		return entrySum{}
	}
	sum := entrySum{start: entries[0].Start}
	var anyNilEnd bool
	var prev = entries[0]
	for _, t := range entries {
		if prev.Start.After(t.Start) {
			log.Fatalf("not sorted properly\n\tCurrent: (%v -> %v)\n\tPrev:    (%v -> %v)", t.Start, t.End, prev.Start, prev.End)
		}
		prev = t
		sum.duration += t.Elapsed()
		if t.End == nil {
			anyNilEnd = true
		} else if sum.end == nil || t.End.After(*sum.end) {
			sum.end = t.End
		}
	}
	if anyNilEnd {
		sum.end = nil
	}
	return sum
}

func uintWidth(i uint) int {
	switch {
	case i < 1e1:
		return 1
	case i < 1e2:
		return 2
	case i < 1e3:
		return 3
	case i < 1e4:
		return 4
	case i < 1e5:
		return 5
	case i < 1e6:
		return 6
	case i < 1e7:
		return 7
	case i < 1e8:
		return 8
	case i < 1e9:
		return 9
	case i < 1e10:
		return 10
	case i < 1e11:
		return 11
	case i < 1e12:
		return 12
	case i < 1e13:
		return 13
	case i < 1e14:
		return 14
	case i < 1e15:
		return 15
	case i < 1e16:
		return 16
	case i < 1e17:
		return 17
	case i < 1e18:
		return 18
	case i < 1e19:
		return 19
	default:
		return 20
	}
}

func timesEqual(a, b time.Time) bool {
	return a.UnixMilli() == b.UnixMilli()
}

func timesPtrsEqual(a, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return timesEqual(*a, *b)
}

func getFirstWeekdayOfISOWeekInMonth(t time.Time) time.Weekday {
	m := t.Month()
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	for t.Month() != m {
		t = t.AddDate(0, 0, 1)
	}
	return t.Weekday()
}

func totalWorkDaysInYear(t time.Time) int {
	workDays := 0
	tm := time.Date(t.Year(), 1, 1, 12, 0, 0, 0, time.UTC)
	for tm.Year() == t.Year() {
		if tm.Weekday() != time.Saturday && tm.Weekday() != time.Sunday {
			workDays++
		}
		tm = tm.Add(24 * time.Hour)
	}
	return workDays
}

func totalWorkDaysInMonth(t time.Time) int {
	workDays := 0
	tm := time.Date(t.Year(), t.Month(), 1, 12, 0, 0, 0, time.UTC)
	for tm.Year() == t.Year() && tm.Month() == t.Month() {
		if tm.Weekday() != time.Saturday && tm.Weekday() != time.Sunday {
			workDays++
		}
		tm = tm.Add(24 * time.Hour)
	}
	return workDays
}

func totalWorkDaysInWeek(t time.Time) int {
	return 5
}

func countWorkDaysYearSoFar(start, end time.Time) int {
	if start.Year() != end.Year() {
		log.Fatalf("tried to count over year boundary, %v - %v", start, end)
	}
	workDaysInYear := totalWorkDaysInYear(start)
	workDays := countWorkDaysSpecial(start, end, func(s, _ time.Time) bool {
		return s.Before(time.Now())
	})
	if workDays > workDaysInYear {
		log.Fatalf("there should never be more than %d work days in a year, but we got %d", workDaysInYear, workDays)
	}
	return workDays
}

func countWorkDaysMonthSoFar(start, end time.Time) int {
	if start.Month() != end.Month() {
		log.Fatalf("tried to count over month boundary, %v - %v", start, end)
	}
	workDaysInMonth := totalWorkDaysInMonth(start)
	workDays := countWorkDaysSpecial(start, end, func(s, _ time.Time) bool {
		return s.Before(time.Now())
	})
	if workDays > workDaysInMonth {
		log.Fatalf("there should never be more than %d work days in %v/%v, but we got %d", workDaysInMonth, workDays, start.Year(), start.Month())
	}
	return workDays
}

func countWorkDaysWeekSoFar(start, end time.Time) int {
	startYear, startWeek := start.ISOWeek()
	endYear, endWeek := end.ISOWeek()
	if startYear != endYear || startWeek != endWeek {
		log.Fatalf("tried to count over week boundary, %v - %v", start, end)
	}
	workDaysInWeek := totalWorkDaysInWeek(start)
	workDays := countWorkDaysSpecial(start, end, func(s, _ time.Time) bool {
		return s.Before(time.Now())
	})
	if workDays > workDaysInWeek {
		log.Fatalf("there should never be more than %d work days in a week, but we got %d", workDaysInWeek, workDays)
	}
	return workDays
}

func countWorkDaysSpecial(start, end time.Time, f func(s, e time.Time) bool) int {
	if !start.Before(end) && !start.Equal(end) {
		log.Fatalf("start should be before or equal to end, but (start) %v > (end) %v", start, end)
	}
	workDays := 0
	dt := start
	for {
		if (dt.Equal(end) || dt.Before(end)) && (f == nil || f(dt, end)) {
			if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
				workDays++
			}
			dt = dt.Add(time.Hour * 24)
		} else {
			break
		}
	}

	return workDays
}

func countWorkdaysInSameMonthWeek(start time.Time) int {
	firstDay := getISOTimeFirstDayInSameMonthWeek(start)
	lastDay := getISOTimeLastDayInSameMonthWeek(start)
	return countWorkDaysWeekSoFar(firstDay, lastDay)
}

func getISOTimeFirstDayInSameMonthWeek(t time.Time) time.Time {
	start := t
	for t.Weekday() != firstWeekday && t.Month() == start.Month() {
		t = t.Add(time.Hour * -24)
	}
	if t.Month() != start.Month() {
		t = t.Add(time.Hour * 24)
	}
	if t.Month() != start.Month() {
		panic("failed, month is not equal")
	}
	return t
}

func getISOTimeLastDayInSameMonthWeek(t time.Time) time.Time {
	start := t
	t = getISOTimeFirstDayInSameMonthWeek(t)
	for t.Weekday() != lastWeekday && t.Month() == start.Month() {
		t = t.Add(time.Hour * 24)
	}

	if t.Month() != start.Month() {
		t = t.Add(time.Hour * -24)
	}

	if t.Month() != start.Month() {
		panic("failed, month is not equal")
	}
	return t
}
