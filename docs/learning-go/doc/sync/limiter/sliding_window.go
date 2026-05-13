package limiter

import (
	"sort"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	limit      int
	windowSize time.Duration
	timeStamps []time.Time
	mu         sync.Mutex
}

func NewSlidingWindowLimiter(limit int, windowSize time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:      limit,
		windowSize: windowSize,
	}
}

// 获取当前时间戳, 并移除过期的时间戳
// 如果当前时间戳的数量小于 limit, 即将当前时间戳添加到切片中, 并返回 true, 表示请求被允许否则返回 false
func (swl *SlidingWindowLimiter) Allow() bool {
	swl.mu.Lock()
	defer swl.mu.Unlock()

	now := time.Now()
	swl.removeExpiredTimeStamps(now)

	if len(swl.timeStamps) < swl.limit {
		swl.timeStamps = append(swl.timeStamps, now)
		return true
	}
	return false
}

// 移除所有过期的时间戳, 即早于当前时间减去窗口大小的时间戳
// 使用 sort.Search 查找第一个不早于 now.Add(-swl.windowSize) 的时间戳的索引, 并截取切片
func (swl *SlidingWindowLimiter) removeExpiredTimeStamps(now time.Time) {
	index := sort.Search(len(swl.timeStamps), func(i int) bool {
		return swl.timeStamps[i].After(now.Add(-swl.windowSize))
	})
	if index < len(swl.timeStamps) {
		swl.timeStamps = swl.timeStamps[index:]
	}
}
