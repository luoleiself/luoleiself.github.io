package limiter

import (
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	t.Run("LeakyBucket", func(t *testing.T) {
		bucket := NewLeakyBucket(10, 1)
		bucket.Start()
		for i := 0; i < 200; i++ {
			go func(reqID int) {
				if bucket.Allow(1) {
					t.Logf("Request %d Allowd\n", reqID)
					time.Sleep(time.Millisecond * 100)
				} else {
					t.Logf("Request %d Denied\n", reqID)
				}
			}(i)
			time.Sleep(time.Millisecond * time.Duration(5+1%300))
		}
		bucket.Stop()
	})
}
