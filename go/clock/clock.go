package clock

import (
	"fmt"
)

// Clock is a clock type
type Clock struct {
	h int
	m int
}

const (
	maxHours = 24
	maxMins = 60
)

// New instantiates a new CLock type
func New(hour, minute int) Clock {
	var c Clock
	c.h = hour
	c.m = minute
	return c.FormatTime()
}

// String returns clock in string format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add will add minutes to a clock time
func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return c.FormatTime()
}

// Subtract will subtract minutes from clock time
func (c Clock) Subtract(minutes int) Clock {
	c.m -= minutes
	return c.FormatTime()
}

// FormatTime ensures that the hours are less than 24 and minutes less than 60
func (c Clock) FormatTime() Clock {

	for c.m >= maxMins {
		c.m -= maxMins
		c.h++
	}

	for c.m < 0 {
		c.m += maxMins
		c.h--
	}

	for c.h < 0 {
		c.h += maxHours
	}

	c.m %= maxMins
	c.h %= maxHours


	return c
}