package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	t.Run("limiter rate 1000 unit 1 second", func(t *testing.T) {
		lc := Exec(1000, time.Second)
		result := <-lc

		fmt.Println("被限流的请求为", result)
	})
}
