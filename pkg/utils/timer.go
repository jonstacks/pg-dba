package utils

import "time"

// Time times the given function
func Time(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
