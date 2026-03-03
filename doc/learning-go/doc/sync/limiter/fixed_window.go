package limiter

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	limit         int
	windowSize    time.Duration
	count         int
	lastResetTime time.Time
	mu            sync.Mutex
}

func NewFixedWindowLimiter(limit int, windowSize time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:         limit,
		windowSize:    windowSize,
		lastResetTime: time.Now(),
	}
}

// 检查当前时间是否已经超过了时间窗口的大小
// 如果超过了时间窗口, 则重置计数器并将 lastResetTime 更新为当前时间
// 如果当前时间窗口内的请求书小于 limit, 则增加计数器并返回 true 表示请求被允许否则返回 false
func (fwl *FixedWindowLimiter) Allow() bool {
	fwl.mu.Lock()
	defer fwl.mu.Unlock()

	now := time.Now()
	if now.Sub(fwl.lastResetTime) >= fwl.windowSize {
		fwl.count = 0
		fwl.lastResetTime = now
	}

	if fwl.count < fwl.limit {
		fwl.count++
		return true
	}
	return false
}
