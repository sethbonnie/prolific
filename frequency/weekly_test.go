package frequency

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFindNext(t *testing.T) {
	testCases := []struct {
		from     time.Time
		weekday  time.Weekday
		expected time.Time
	}{
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Monday,
			expected: time.Date(2022, 4, 4, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Tuesday,
			expected: time.Date(2022, 4, 5, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Wednesday,
			expected: time.Date(2022, 4, 6, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Thursday,
			expected: time.Date(2022, 4, 7, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Friday,
			expected: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Saturday,
			expected: time.Date(2022, 4, 2, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
		{
			// Friday
			from:     time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			weekday:  time.Sunday,
			expected: time.Date(2022, 4, 3, 0, 0, 0, 0, time.Now().UTC().Location()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.weekday.String(), func(t *testing.T) {
			actual := findNext(tc.weekday, tc.from)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestWeeklyFrequency(t *testing.T) {
	aFriday := time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location())
	testCases := []struct {
		name      string
		startDate time.Time
		every     int
		on        []time.Weekday
		date      time.Time
		expect    bool
	}{
		{
			name:      "NextWeek",
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday},
			// The next Friday
			date:   aFriday.AddDate(0, 0, 7),
			expect: true,
		},
		{
			name: "Every2WeeksCheckNextWeek",
			// Friday
			startDate: aFriday,
			// Every other week
			every:  2,
			on:     []time.Weekday{time.Friday},
			date:   aFriday.AddDate(0, 0, 7),
			expect: false,
		},
		{
			name: "NextDay",
			// Friday
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday},
			// The next Saturday
			date:   aFriday.AddDate(0, 0, 1),
			expect: false,
		},
		{
			name: "EveryFridayAndSaturdayCheckSaturday",
			// Friday
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday, time.Saturday},
			// The next Saturday
			date:   aFriday.AddDate(0, 0, 1),
			expect: true,
		},
		{
			name:      "PreviousWeek",
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday},
			date:      aFriday.AddDate(0, 0, -7),
			expect:    false,
		},
		{
			name: "DateFurtherOutMatchingWeekday",
			// Friday
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday},
			// A Friday
			date:   aFriday.AddDate(0, 0, 7*18),
			expect: true,
		},
		{
			name: "DateFurtherOutMismatchingWeekday",
			// Friday
			startDate: aFriday,
			every:     1,
			on:        []time.Weekday{time.Friday},
			// A Friday
			date:   aFriday.AddDate(0, 0, (7*18)+1),
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, err := Weekly(tc.startDate, tc.every, tc.on)
			require.Nil(t, err)
			require.Equal(t, tc.expect, frequency.Check(tc.date))
		})
	}
}

func TestWeekdayHelpers(t *testing.T) {
	fns := map[time.Weekday]func(time.Time) (Frequency, error){
		time.Sunday:    Sundays,
		time.Monday:    Mondays,
		time.Tuesday:   Tuesdays,
		time.Wednesday: Wednesdays,
		time.Thursday:  Thursdays,
		time.Friday:    Fridays,
		time.Saturday:  Saturdays,
	}

	type testCase struct {
		name     string
		fn       func(time.Time) (Frequency, error)
		start    time.Time
		date     time.Time
		expected bool
	}

	testCases := []testCase{}
	diffs := map[int]bool{
		0:  true,
		1:  false,
		7:  true,
		8:  false,
		9:  false,
		25: false,
		35: true,
	}

	for weekday, fn := range fns {
		for days, match := range diffs {
			start := findNext(weekday, time.Now())
			testCases = append(testCases, testCase{
				name:     fmt.Sprintf("%s+%d", weekday, days),
				fn:       fn,
				start:    start,
				date:     start.AddDate(0, 0, days),
				expected: match,
			})
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			frequency, err := tc.fn(tc.start)
			require.Nil(t, err)
			require.Equal(t, tc.expected, frequency.Check(tc.date))
		})
	}
}

func TestWeekdays(t *testing.T) {
	type testCase struct {
		name     string
		from     time.Time
		date     time.Time
		expected bool
	}
	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	testCases := []testCase{}

	for _, wd := range weekdays {
		from := time.Now()
		var expected bool
		if wd != time.Sunday && wd != time.Saturday {
			expected = true
		}
		testCases = append(testCases, testCase{
			name:     wd.String(),
			from:     from,
			date:     findNext(wd, from),
			expected: expected,
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, err := Weekdays(tc.from)
			require.Nil(t, err)
			require.Equal(t, tc.expected, frequency.Check(tc.date))
		})
	}
}

func TestWeekends(t *testing.T) {
	type testCase struct {
		name     string
		from     time.Time
		date     time.Time
		expected bool
	}
	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	testCases := []testCase{}

	for _, wd := range weekdays {
		from := time.Now()
		var expected bool
		if wd == time.Sunday || wd == time.Saturday {
			expected = true
		}
		testCases = append(testCases, testCase{
			name:     wd.String(),
			from:     from,
			date:     findNext(wd, from),
			expected: expected,
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, err := Weekends(tc.from)
			require.Nil(t, err)
			require.Equal(t, tc.expected, frequency.Check(tc.date))
		})
	}
}
