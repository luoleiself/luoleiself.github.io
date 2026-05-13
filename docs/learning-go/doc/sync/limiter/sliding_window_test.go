package limiter

import (
	"testing"
	"time"
)

func TestSlidingWindowLimiter(t *testing.T) {
	t.Run("SlidingWindowLimiter", func(t *testing.T) {
		limiter := NewSlidingWindowLimiter(100, time.Second)
		allowed := make([]int, 0)
		denied := make([]int, 0)
		for i := 0; i < 1000; i++ {
			if limiter.Allow() {
				allowed = append(allowed, i)
			} else {
				denied = append(denied, i)
			}
			time.Sleep(time.Millisecond * 3)
		}
		t.Logf("Request Allowed %v\n", allowed)
		t.Logf("Request Denied %v\n", denied)
	})
}
