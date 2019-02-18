/*
 * The MIT License
 *
 * Copyright 2018 kinsey40.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * File:   clock.go
 * Author: kinsey40
 *
 * Created on 13 January 2019, 11:05
 *
 * Clock wraps functionality from the time module that is useful for tracking
 * the progress of the Pbar object.
 *
 */

package clock

import (
	"fmt"
	"time"
)

// Clock enables various operations relating to time to be performed
// easily.
type Clock interface {
	Now() time.Time
	Subtract(time.Time) time.Duration
	SetStart(time.Time)
	Start() time.Time
	Seconds(time.Duration) float64
	Remaining(float64) time.Duration
	Format(time.Duration) string
}

// clock implements a real-time clock by wrapping functions from the
// time module. It also contains a start time relating to when the
// Pbar object was initialized.
type clock struct {
	StartTime time.Time
}

// NewClock returns an instance of a real-time clock.
func NewClock() Clock {
	c := new(clock)

	return c
}

// Now returns the current time (from the time module).
func (c *clock) Now() time.Time {
	return time.Now()
}

// Subtract finds the difference between a passed in time
// and the start time.
func (c *clock) Subtract(now time.Time) time.Duration {
	return now.Sub(c.StartTime)
}

// SetStart enables the StartTime value to be set in the clock
// object.
func (c *clock) SetStart(t time.Time) {
	c.StartTime = t
}

// Start returns the StartTime for the clock object
func (c *clock) Start() time.Time {
	return c.StartTime
}

// Seconds returns the number of seconds in a time.Duration object
func (c *clock) Seconds(d time.Duration) float64 {
	return d.Seconds()
}

// Remaining returns a time.Duration object equating to the fraction
// of progress that has been performed.
func (c *clock) Remaining(fraction float64) time.Duration {
	return time.Duration(fraction) * time.Second
}

// Format enables a time.Duration object to be formatted into a string
// format that can be easily integrated into the progress bar.
func (c *clock) Format(d time.Duration) string {
	secs := (d % time.Minute) / time.Second
	mins := (d % time.Hour) / time.Minute
	hours := d / time.Hour

	if hours == 0 {
		return fmt.Sprintf("%02dm:%02ds", mins, secs)
	}

	return fmt.Sprintf("%02dh:%02dm:%02ds", hours, mins, secs)
}