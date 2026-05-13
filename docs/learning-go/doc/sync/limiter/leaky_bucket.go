package limiter

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	water    int
	leakRate int
	mutex    sync.Mutex
	ticker   *time.Ticker
	stopCh   chan struct{}
}

func NewLeakyBucket(capacity, leakRate int) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		stopCh:   make(chan struct{}),
	}
}

func (lb *LeakyBucket) Start() {
	lb.ticker = time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-lb.ticker.C:
				lb.leak()
			case <-lb.stopCh:
				lb.ticker.Stop()
				return
			}
		}
	}()
}

func (lb *LeakyBucket) Stop() {
	close(lb.stopCh)
}
func (lb *LeakyBucket) leak() {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	if lb.water > 0 {
		lb.water -= lb.leakRate
		if lb.water < 0 {
			lb.water = 0
		}
	}
}
func (lb *LeakyBucket) Allow(amount int) bool {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	if lb.water+amount <= lb.capacity {
		lb.water += amount
		return true
	}
	return false
}
