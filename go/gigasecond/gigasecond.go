// Package gigasecond contains a single method AddGigasecond that performs time addition
package gigasecond

import "time"

// AddGigasecond takes time input and returns the time after 10^9 seconds have passed
func AddGigasecond(t time.Time) time.Time {
	duration, _ := time.ParseDuration("1000000000s")
	return t.Add(time.Duration(duration))
}
