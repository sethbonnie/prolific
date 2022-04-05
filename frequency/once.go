package frequency

import "time"

type once struct {
	date time.Time
}

func (o once) Check(t time.Time) bool {
	return o.date.Year() == t.Year() && o.date.YearDay() == t.YearDay()
}

func Once(t time.Time) Frequency {
	return once{date: t}
}
