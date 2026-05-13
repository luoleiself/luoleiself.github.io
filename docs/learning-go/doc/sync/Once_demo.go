package main

import (
	"fmt"
	"sync"
	"time"
)

func OnceNote() {
	fmt.Println("-------------OnceNote()-------------")
	var once sync.Once
	// fmt.Println(once)
	once.Do(func() {
		fmt.Println("1. once.Do(func(){}) 方法当且仅当第一次被调用时才执行函数f")
		time.Sleep(2 * time.Second)
	})
	fmt.Println("once.Do 后的log 1")
	once.Do(func() {
		fmt.Println("2. once.Do(func(){}) 方法当且仅当第一次被调用时才执行函数f")
	})
	once.Do(func() {
		fmt.Println("3. once.Do(func(){}) 方法当且仅当第一次被调用时才执行函数f")
	})
}
