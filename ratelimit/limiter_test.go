package ratelimit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFixedWindowLimiter(t *testing.T) {

	fwLimiter := NewFixedWindowLimiter(3 * time.Second, 100)

	ch := make(chan int, 1000)
	total := 1000

	for i := 0; i < total; i++ {
		go func(n int) {
			if !fwLimiter.Limit() {
				ch <- 0
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	assert.Equal(t, 100, len(ch))
}
