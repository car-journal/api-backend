package clock

import (
	"time"
)

type ClockInterface interface {
	Now() time.Time
}

type Clock struct{}

func New() ClockInterface {
	return &Clock{}
}

func (t *Clock) Now() time.Time {
	return time.Now()
}
