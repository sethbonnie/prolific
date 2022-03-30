package interval

import (
	"time"
)

type Interval interface {
	IsActive(time.Time) bool
}
