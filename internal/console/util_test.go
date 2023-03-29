// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// SPDX-FileCopyrightText: 2023 Alexander Fougner
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
	"testing"
	"time"
)

func TestGetISOTimeFirstDayOfSameMonthWeek(t *testing.T) {
	testCases := []struct {
		start time.Time
		want  time.Time
	}{
		{
			start: time.Date(2023, 2, 28, 12, 0, 0, 0, time.UTC),
			want:  time.Date(2023, 2, 26, 12, 0, 0, 0, time.UTC),
		},
		{
			start: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC),
			want:  time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC),
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCase idx %d", idx), func(t *testing.T) {
			got := getISOTimeFirstDayInSameMonthWeek(tc.start)
			if got != tc.want {
				t.Errorf("got %v; wanted %v", got, tc.want)
			}
		})
	}
}

func TestGetISOTimeLastDayOfSameMonthWeek(t *testing.T) {
	testCases := []struct {
		start time.Time
		want  time.Time
	}{
		{
			start: time.Date(2023, 2, 27, 12, 0, 0, 0, time.UTC),
			want:  time.Date(2023, 2, 28, 12, 0, 0, 0, time.UTC),
		},
		{
			start: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC),
			want:  time.Date(2023, 3, 4, 12, 0, 0, 0, time.UTC),
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCase idx %d", idx), func(t *testing.T) {
			got := getISOTimeLastDayInSameMonthWeek(tc.start)
			if got != tc.want {
				t.Errorf("got %v; wanted %v", got, tc.want)
			}
		})
	}
}

func TestCountWorkdaysInSameMonthWeek(t *testing.T) {
	testCases := []struct {
		start time.Time
		want  int
	}{
		{
			start: time.Date(2023, 2, 27, 12, 0, 0, 0, time.UTC),
			want:  2,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCase idx %d", idx), func(t *testing.T) {
			fmt.Printf("count workdays idx %d\n", idx)
			got := countWorkdaysInSameMonthWeek(tc.start)
			if got != tc.want {
				t.Errorf("got %v; wanted %v", got, tc.want)
			}
		})
	}
}
