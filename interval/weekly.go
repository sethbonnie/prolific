package interval

import (
	"time"
)

type weekdays map[time.Weekday]bool

type weeklyInterval struct {
	startDate time.Time
	every     int // every n weeks
	on        weekdays
}

func Weekly(start time.Time, every int, days []time.Weekday) Interval {
	on := make(weekdays)
	for _, d := range days {
		on[d] = true
	}
	return weeklyInterval{start, every, on}
}

func (w weeklyInterval) IsActive(t time.Time) bool {
	if !w.on[t.Weekday()] {
		return false
	}

	d := w.startDate
	date := t.Format(dateFormat)

	for d.Before(t) {
		for wd := range w.on {
			if findNext(wd, d).Format(dateFormat) == date {
				return true
			}
		}

		d = d.AddDate(0, 0, w.every*7)
	}

	return d.Format(dateFormat) == date
}

func Sundays(from time.Time) Interval {
	return Weekly(findNext(time.Sunday, from), 1, []time.Weekday{time.Sunday})
}

func Mondays(from time.Time) Interval {
	return Weekly(findNext(time.Monday, from), 1, []time.Weekday{time.Monday})
}

func Tuesdays(from time.Time) Interval {
	return Weekly(findNext(time.Tuesday, from), 1, []time.Weekday{time.Tuesday})
}

func Wednesdays(from time.Time) Interval {
	return Weekly(findNext(time.Wednesday, from), 1, []time.Weekday{time.Wednesday})
}

func Thursdays(from time.Time) Interval {
	return Weekly(findNext(time.Thursday, from), 1, []time.Weekday{time.Thursday})
}

func Fridays(from time.Time) Interval {
	return Weekly(findNext(time.Friday, from), 1, []time.Weekday{time.Friday})
}

func Saturdays(from time.Time) Interval {
	return Weekly(findNext(time.Saturday, from), 1, []time.Weekday{time.Saturday})
}

func Weekends(from time.Time) Interval {
	d := from.Weekday()
	if d != time.Saturday && d != time.Sunday {
		from = findNext(time.Saturday, from)
	}
	return weeklyInterval{
		startDate: findNext(time.Saturday, from),
		every:     1,
		on: weekdays{
			time.Saturday: true,
			time.Sunday:   true,
		},
	}
}

func Weekdays(from time.Time) Interval {
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
