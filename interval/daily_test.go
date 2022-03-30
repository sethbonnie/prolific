package interval

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
