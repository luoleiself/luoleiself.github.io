package limiter

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	t.Run("TokenBucket", func(t *testing.T) {
		tb := NewTokenBucket(5, time.Second)

		for i := 0; i < 10; i++ {
			if tb.Allow() {
				t.Log("Request", i, "allowed")
			} else {
				t.Log("Request", i, "denied")
			}
			time.Sleep(time.Millisecond * 200)
		}
	})
}
