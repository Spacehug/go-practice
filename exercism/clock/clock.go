package clock

import (
	"fmt"
)

type Clock uint16

const minutesInDay int = 1440

func New(hours, minutes int) Clock {
	mt := (hours * 60 + minutes) % minutesInDay
	if mt < 0 {mt += minutesInDay}
	return Clock(mt)
}

func (c Clock) Add(a int) Clock {
	return New(0, int(c) + a)
}

func (c Clock) Subtract(a int) Clock {
	return New(0, int(c) - a)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c / 60, c % 60)
}

