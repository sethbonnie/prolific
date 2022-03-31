package frequency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDaily(t *testing.T) {
	from := time.Now().UTC()
	testCases := []struct {
		name      string
		from      time.Time
		every     int
		until     time.Time
		date      time.Time
		expError  bool
		expActive bool
	}{
		{
			name:     "NegativeEvery",
			from:     from,
			every:    -1,
			expError: true,
		},
		{
			name:     "Every=0",
			from:     from,
			every:    0,
			expError: true,
		},
		{
			name:     "UntilBeforefrom",
			from:     from,
			every:    1,
			until:    from.AddDate(0, 0, -1),
			expError: true,
		},
		{
			name:      "DailyCheckToday",
			from:      from,
			every:     1,
			date:      from,
			expActive: true,
		},
		{
			name:      "DailyCheckTomorrow",
			from:      from,
			every:     1,
			date:      from.AddDate(0, 0, 1),
			expActive: true,
		},
		{
			name:      "EveryOtherDayCheckTomorrow",
			from:      from,
			every:     2,
			date:      from.AddDate(0, 0, 1),
			expActive: false,
		},
		{
			name:      "For1WeekCheck5Days",
			from:      from,
			every:     1,
			until:     from.AddDate(0, 0, 7),
			date:      from.AddDate(0, 0, 5),
			expActive: true,
		},
		{
			name:      "For1WeekCheck2Weeks",
			from:      from,
			every:     1,
			until:     from.AddDate(0, 0, 7),
			date:      from.AddDate(0, 0, 10),
			expActive: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := Daily(tc.from, tc.every, tc.until)
			if tc.expError {
				require.NotNil(t, err)
				return
			}
			require.Equal(t, tc.expActive, d.IsActive(tc.date))
		})
	}
}

func TestDailyForWeeks(t *testing.T) {
	from := time.Now().UTC()
	testCases := []struct {
		name     string
		from     time.Time
		weeks    int
		every    int
		date     time.Time
		expected bool
	}{
		{
			name:     "1WeekAfter8Days",
			from:     from,
			weeks:    1,
			every:    1,
			date:     from.AddDate(0, 0, 8),
			expected: false,
		},
		{
			name:     "1WeekAfter5Days",
			from:     from,
			weeks:    1,
			every:    1,
			date:     from.AddDate(0, 0, 5),
			expected: true,
		},
		{
			name:     "10WeeksAfter5Weeks",
			from:     from,
			weeks:    10,
			every:    1,
			date:     from.AddDate(0, 0, 5*7),
			expected: true,
		},
		{
			name:     "10WeeksAfter16Weeks",
			from:     from,
			weeks:    10,
			every:    1,
			date:     from.AddDate(0, 0, 16*7),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, _ := DailyForWeeks(tc.from, tc.every, tc.weeks)
			require.Equal(t, tc.expected, frequency.IsActive(tc.date))
		})
	}
}

func TestDailyForMonths(t *testing.T) {
	from := time.Now().UTC()
	testCases := []struct {
		name     string
		from     time.Time
		months   int
		every    int
		date     time.Time
		expected bool
	}{
		{
			name:     "1MonthAfter8Days",
			from:     from,
			months:   1,
			every:    1,
			date:     from.AddDate(0, 0, 8),
			expected: true,
		},
		{
			name:     "1MonthAfter5Days",
			from:     from,
			months:   1,
			every:    1,
			date:     from.AddDate(0, 0, 5),
			expected: true,
		},
		{
			name:     "1MonthAfter1Month",
			from:     from,
			months:   1,
			every:    1,
			date:     from.AddDate(0, 1, 0),
			expected: true,
		},
		{
			name:     "1MonthAfter1MonthAnd1Day",
			from:     from,
			months:   1,
			every:    1,
			date:     from.AddDate(0, 1, 1),
			expected: false,
		},
		{
			name:     "10MonthsAfter5Months",
			from:     from,
			months:   10,
			every:    1,
			date:     from.AddDate(0, 5, 0),
			expected: true,
		},
		{
			name:     "10MonthsAfter24Months",
			from:     from,
			months:   10,
			every:    1,
			date:     from.AddDate(0, 24, 0),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, _ := DailyForMonths(tc.from, tc.every, tc.months)
			require.Equal(t, tc.expected, frequency.IsActive(tc.date))
		})
	}
}

func TestDailyForYears(t *testing.T) {
	from := time.Now().UTC()
	testCases := []struct {
		name     string
		from     time.Time
		years    int
		every    int
		date     time.Time
		expected bool
	}{
		{
			name:     "1YearAfter1Day",
			from:     from,
			years:    1,
			every:    1,
			date:     from.AddDate(0, 0, 1),
			expected: true,
		},
		{
			name:     "1YearAfter1Month",
			from:     from,
			years:    1,
			every:    1,
			date:     from.AddDate(0, 1, 0),
			expected: true,
		},
		{
			name:     "1YearAfter1Year",
			from:     from,
			years:    1,
			every:    1,
			date:     from.AddDate(1, 0, 0),
			expected: true,
		},
		{
			name:     "1YearAfter1YearAnd1Day",
			from:     from,
			years:    1,
			every:    1,
			date:     from.AddDate(1, 0, 1),
			expected: false,
		},
		{
			name:     "1YearAfter5Years",
			from:     from,
			years:    1,
			every:    1,
			date:     from.AddDate(5, 0, 0),
			expected: false,
		},
		{
			name:     "2YearsAfter18Months",
			from:     from,
			years:    2,
			every:    1,
			date:     from.AddDate(0, 18, 0),
			expected: true,
		},
		{
			name:     "2YearsAfter25Months",
			from:     from,
			years:    2,
			every:    1,
			date:     from.AddDate(0, 25, 0),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			frequency, _ := DailyForYears(tc.from, tc.every, tc.years)
			require.Equal(t, tc.expected, frequency.IsActive(tc.date))
		})
	}
}
