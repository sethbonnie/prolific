package frequency

import (
	"fmt"
	"time"
)

type weekdays map[time.Weekday]bool

type weeklyFrequency struct {
	startDate time.Time
	every     int // every n weeks
	on        weekdays
}

func (w weeklyFrequency) Check(t time.Time) bool {
	if !w.on[t.Weekday()] {
		return false
	}

	daysDiff := t.Sub(w.startDate).Hours() / 24

	if daysDiff < 0 {
		return false
	}

	if int(daysDiff)%(7*w.every) == 0 {
		return true
	}

	for wd := range w.on {
		next := findNext(wd, w.startDate)
		daysDiff = t.Sub(next).Hours() / 24

		if daysDiff < 0 {
			return false
		}

		if int(daysDiff)%(7*w.every) == 0 {
			return true
		}
	}

	return false
}

func Weekly(start time.Time, every int, days []time.Weekday) (Frequency, error) {
	if len(days) == 0 {
		return weeklyFrequency{}, ErrEmptyFrequency
	}

	if every < 1 {
		return weeklyFrequency{}, fmt.Errorf("%w: every %v", ErrNonPositiveInt, every)
	}

	on := make(weekdays)
	for _, d := range days {
		on[d] = true
	}
	return weeklyFrequency{start, every, on}, nil
}

func Sundays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Sunday, from), 1, []time.Weekday{time.Sunday})
}

func Mondays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Monday, from), 1, []time.Weekday{time.Monday})
}

func Tuesdays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Tuesday, from), 1, []time.Weekday{time.Tuesday})
}

func Wednesdays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Wednesday, from), 1, []time.Weekday{time.Wednesday})
}

func Thursdays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Thursday, from), 1, []time.Weekday{time.Thursday})
}

func Fridays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Friday, from), 1, []time.Weekday{time.Friday})
}

func Saturdays(from time.Time) (Frequency, error) {
	return Weekly(findNext(time.Saturday, from), 1, []time.Weekday{time.Saturday})
}

func Weekdays(from time.Time) (Frequency, error) {
	d := from.Weekday()
	if d == time.Saturday || d == time.Sunday {
		from = findNext(time.Monday, from)
	}
	return Weekly(
		from,
		1,
		[]time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
	)
}

func Weekends(from time.Time) (Frequency, error) {
	d := from.Weekday()
	if d != time.Saturday && d != time.Sunday {
		from = findNext(time.Saturday, from)
	}
	return Weekly(from, 1, []time.Weekday{time.Sunday, time.Saturday})
}

func findNext(next time.Weekday, start time.Time) time.Time {
	from := start.Weekday()
	if from == next {
		return start
	}

	var days int
	if next > from {
		// Just add the difference between the days
		days = int(next - from)
	} else {
		// Figure out how many days they are apart (from - next), subtract from a week (7 - diff)
		diff := (7 - (from - next))
		days = int(diff)
	}
	return start.AddDate(0, 0, days)
}
