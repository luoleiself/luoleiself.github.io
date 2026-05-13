package main

import (
	"runtime"
	"testing"
)

func TestGoroutine(t *testing.T) {
	t.Run("goroutine", func(t *testing.T) {
		t.Logf("runtime.GOMAXPROCS %d\n", runtime.NumCPU())
		t.Logf("runtime.GOOS %s\n", runtime.GOOS)
		t.Logf("runtime.GOARCH %s\n", runtime.GOARCH)
		t.Logf("runtime.GOROOT %s\n", runtime.GOROOT())
	})
}
