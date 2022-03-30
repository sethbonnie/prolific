package interval

import (
	"time"
)

type Interval interface {
	IsActive(time.Time) bool
}

const dateFormat = "2006/01/02"
