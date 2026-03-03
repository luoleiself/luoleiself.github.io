package limiter

import (
	"sync"
	"time"
)

type Limter struct {
	rate  int           // 每分钟请求限制
	unit  time.Duration // 时间单位
	mutex sync.Mutex    // 并发控制器
	count int           // 当前时间窗口内的请求计数
	reset time.Time     // 使劲按窗口重置时间
}

func NewLimiter(rate int, unit time.Duration) *Limter {
	return &Limter{
		rate:  rate,
		unit:  unit,
		mutex: sync.Mutex{},
		count: 0,
		reset: time.Now(),
	}
}

func (l *Limter) Allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 检查是否需要重置时间窗口
	if time.Since(l.reset) >= l.unit {
		l.count = 0
		l.reset = time.Now()
	}

	// 判断请求是否超过限制
	if l.count >= l.rate {
		return false
	}

	// 请求计数加一
	l.count++
	return true
}

func Exec(rate int, unit time.Duration) <-chan int {
	limiter := NewLimiter(rate, unit)
	lc := make(chan int)
	go func() {
		defer close(lc)
		rate = rate * 2
		for i := 0; i <= rate; i++ {
			if !limiter.Allow() {
				lc <- i + 1
				break
			}
		}
	}()
	return lc
}
