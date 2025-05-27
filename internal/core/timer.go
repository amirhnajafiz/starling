package core

import "time"

// timer is a simple timer implementation that can be started and stopped.
type timer struct {
	// duration is the time duration for the timer
	duration int64
	// isRunning indicates whether the timer is currently running
	isRunning bool
	// tickChannel is used to signal when the timer ticks
	tickChannel chan struct{}
}

// newTimer creates a new timer with the specified duration.
func newTimer(duration int64) *timer {
	return &timer{
		duration:    duration,
		isRunning:   false,
		tickChannel: make(chan struct{}),
	}
}

// start starts the timer.
func (t *timer) start() {
	if t.isRunning {
		return
	}
	t.isRunning = true

	// using time.Ticker to handle the timer ticks
	go func() {
		ticker := time.NewTicker(time.Duration(t.duration) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			if t.isRunning {
				t.tickChannel <- struct{}{} // signal that the timer has ticked
			}
		}
	}()
}

// stop stops the timer.
func (t *timer) stop() {
	if !t.isRunning {
		return // timer is not running
	}
	t.isRunning = false

	// close the tick channel to stop receiving ticks
	close(t.tickChannel)
	t.tickChannel = make(chan struct{}) // reinitialize the channel for future use
}
