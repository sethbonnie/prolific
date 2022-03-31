package interval

import (
	"fmt"
	"time"
)

type dailyInterval struct {
	startDate time.Time
	every     int // every n days
	until     time.Time
}

func (d dailyInterval) IsActive(t time.Time) bool {
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

func Daily(from time.Time, every int, until time.Time) (Interval, error) {
	if every < 1 {
		return dailyInterval{}, fmt.Errorf("%w: every %v", ErrNonPositiveInt, every)
	}

	d := dailyInterval{
		startDate: from,
		every:     every,
	}

	if !until.IsZero() {
		if !from.Before(until) {
			return dailyInterval{}, fmt.Errorf("from (%v) must be before until (%v)", every, until)
		}
		d.until = until
	}

	return d, nil
}
