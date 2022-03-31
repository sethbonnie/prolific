package frequency

import (
	"fmt"
	"time"
)

type dailyFrequency struct {
	startDate time.Time
	every     int // every n days
	until     time.Time
}

func (d dailyFrequency) IsActive(t time.Time) bool {
	if t.Before(d.startDate) {
		return false
	}

	if !d.until.IsZero() && t.After(d.until) {
		return false
	}

	diff := t.Sub(d.startDate).Hours() / 24

	if int(diff)%d.every == 0 {
		return true
	}
	return false
}

func Daily(from time.Time, every int, until time.Time) (Frequency, error) {
	if every < 1 {
		return dailyFrequency{}, fmt.Errorf("%w: every %v", ErrNonPositiveInt, every)
	}

	d := dailyFrequency{
		startDate: from,
		every:     every,
	}

	if !until.IsZero() {
		if !from.Before(until) {
			return dailyFrequency{}, fmt.Errorf("from (%v) must be before until (%v)", every, until)
		}
		d.until = until
	}

	return d, nil
}

func DailyForWeeks(from time.Time, every, weeks int) (Frequency, error) {
	until := from.AddDate(0, 0, 7*weeks)
	return Daily(from, every, until)
}

func DailyForMonths(from time.Time, every, months int) (Frequency, error) {
	until := from.AddDate(0, months, 0)
	return Daily(from, every, until)
}

func DailyForYears(from time.Time, every, years int) (Frequency, error) {
	until := from.AddDate(years, 0, 0)
	return Daily(from, every, until)
}
