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
