package main

import (
	"sync"
	"testing"
)

var mutex sync.Mutex
var rwmutex sync.RWMutex
var data int

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		_ = data
		mutex.Unlock()
	}
}

func BenchmarkRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rwmutex.RLock()
		_ = data
		rwmutex.RUnlock()
	}
}
