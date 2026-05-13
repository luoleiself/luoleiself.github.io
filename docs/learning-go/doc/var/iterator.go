package main

import "fmt"

// iterator 函数签名
// func(func() bool)
// func(func(K) bool)
// func(func(K, V) bool)

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// 迭代器组合和泛型
type Iterator[A any] func(func(A) bool)
type IntGen struct {
	current int
}

func (g IntGen) Generate() Iterator[int] {
	return func(yield func(int) bool) {
		for {
			if !yield(g.current) {
				return
			}
			g.current++
		}
	}
}

func RangeOverFunc() {
	var g = IntGen{}
	// for range 的 body 被转换为了 iterator 的 yield 的实现
	for v := range g.Generate() {
		fmt.Println(v)
	}
}
