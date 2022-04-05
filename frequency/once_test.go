package frequency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {
	testCases := []struct {
		name     string
		on       time.Time
		date     time.Time
		expected bool
	}{
		{
			name:     "SameDay",
			on:       time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			date:     time.Date(2022, 4, 1, 22, 0, 0, 0, time.Now().UTC().Location()),
			expected: true,
		},
		{
			name:     "NextDay",
			on:       time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			date:     time.Date(2022, 4, 2, 0, 0, 0, 0, time.Now().UTC().Location()),
			expected: false,
		},
		{
			name:     "PreviousDay",
			on:       time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			date:     time.Date(2022, 3, 31, 23, 59, 59, 0, time.Now().UTC().Location()),
			expected: false,
		},
		{
			name:     "NextYear",
			on:       time.Date(2022, 4, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			date:     time.Date(2023, 4, 1, 22, 0, 0, 0, time.Now().UTC().Location()),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := Once(tc.on)
			assert.Equal(t, tc.expected, f.Check(tc.date))
		})
	}
}
