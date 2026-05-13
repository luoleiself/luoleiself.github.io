package limiter

import (
	"testing"
	"time"
)

func TestFixedWindowLimiter(t *testing.T) {
	t.Run("FixedWindowLimiter", func(t *testing.T) {
		// 创建一个每秒钟最多允许 100 个请求的限流器
		// 循环 1000 次, 每次尝试发送请求, 并根据 Allow 方法的返回值统计相应的结果
		// 每次循环之间暂停 3 毫秒
		limiter := NewFixedWindowLimiter(100, time.Second)
		allowd := make([]int, 0)
		denied := make([]int, 0)
		for i := 0; i < 1000; i++ {
			if limiter.Allow() {
				allowd = append(allowd, i)
			} else {
				denied = append(denied, i)
			}
			time.Sleep(time.Millisecond * 3)
		}
		t.Logf("Request Allowd %v\n", allowd)
		t.Logf("Request Denied %v\n", denied)
	})
}
