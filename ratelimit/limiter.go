package ratelimit

import (
	"sync/atomic"
	"time"
	"unsafe"
)

type Limiter interface {
	Limit() bool
}

type limiterState struct {
	counter    int64
	lastWindow time.Time
}

type FixedWindowLimiter struct {
	ls       unsafe.Pointer
	interval time.Duration
	limit    int64
}

func NewFixedWindowLimiter(interval time.Duration, limit int64) *FixedWindowLimiter {
	ls := &limiterState{
		counter:    0,
		lastWindow: time.Time{},
	}
	return &FixedWindowLimiter{
		ls: unsafe.Pointer(ls),
		interval: interval,
		limit:    limit,
	}
}

func (fw *FixedWindowLimiter) Limit() bool {
	now := time.Now()
	swapped := false

	for !swapped {
		lastState := (* limiterState)(fw.ls)
		counter := lastState.counter

		if counter >= fw.limit {
			return true
		}

		if lastState.lastWindow.IsZero() || now.Sub(lastState.lastWindow) >= fw.interval {
			newState := &limiterState{
				counter:    1,
				lastWindow: now,
			}
			swapped = atomic.CompareAndSwapPointer(&fw.ls, fw.ls, unsafe.Pointer(newState))
			continue
		}

		atomic.AddInt64(&lastState.counter, 1)
		return false
	}

	return false
}
