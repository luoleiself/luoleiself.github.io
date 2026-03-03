package main

import (
	"fmt"
	"sync"
	"time"
)

func MutexNote() {
	fmt.Println("-------------MutexNote()-------------")
	fmt.Println("互斥锁: Mutex 类型的锁和 goroutine 无关, 可以由不同的 goroutine 加锁和解锁")
	var mutex sync.Mutex
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d 尝试锁定. \n", i)
			mutex.Lock()
			fmt.Printf("goroutine %d 已经锁定了。\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d 解锁 \n", i)
			mutex.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("main... 尝试锁定\n")
	mutex.Lock()
	fmt.Printf("main... 已经锁定了\n")
	time.Sleep(3 * time.Second)
	mutex.Unlock()
	fmt.Printf("main... 解锁\n")
}
