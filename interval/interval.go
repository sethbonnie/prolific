package interval

import (
	"errors"
	"time"
)

type Interval interface {
	IsActive(time.Time) bool
}

var (
	ErrNonPositiveInt = errors.New("None positive int")
	ErrEmptyInterval  = errors.New("Empty interval")
)

const dateFormat = "2006/01/02"
