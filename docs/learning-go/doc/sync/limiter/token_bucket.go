package limiter

import "time"

type TokenBucket struct {
	capacity int           // 桶的最大容量
	tokens   int           // 当前桶中的令牌数量
	rate     time.Duration // 令牌生成速率
	lastTime time.Time     // 上次令牌生成时间
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	now := time.Now()
	elapsed := now.Sub(tb.lastTime)
	tokensToAdd := int(elapsed / tb.rate)
	tb.lastTime = now

	// 如果有令牌生成, 更新桶中的令牌数量
	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
	}

	// 如果有令牌, 则消费一个令牌并返回 true
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	// 否则返回 false, 拒绝请求
	return false
}
