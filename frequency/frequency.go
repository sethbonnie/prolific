package frequency

import (
	"errors"
	"time"
)

type Frequency interface {
	IsActive(time.Time) bool
}

var (
	ErrNonPositiveInt = errors.New("None positive int")
	ErrEmptyFrequency = errors.New("Empty frequency")
)

const dateFormat = "2006/01/02"
