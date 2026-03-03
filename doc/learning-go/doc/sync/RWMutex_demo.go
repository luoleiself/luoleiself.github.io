package main

import (
	"fmt"
	"sync"
	"time"
)

func RWMutexNote() {
	fmt.Println("-------------RWMutexNote()-------------")
	fmt.Println("读写互斥锁: 写操作是互斥的, 读和写是互斥的, 读和读不互斥")
	fmt.Println(tab, "同时只能有一个 goroutine 能够获得写锁定")
	fmt.Println(tab, "同时可以有任意多个 goroutine 获得读锁定")
	fmt.Println(tab, "同时只能存在写锁定或读锁定(读和写互斥)")
	fmt.Println("RWMutex 类型的锁和 goroutine 无关, 可以由不同的 goroutine 加读取锁/写入锁和解读取锁/写入锁")
	var rwm sync.RWMutex
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d 尝试读锁定. \n", i)
			rwm.RLock()
			fmt.Printf("goroutine %d 已经读锁定了。\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d 读解锁 \n", i)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("main... 尝试写锁定\n")
	rwm.Lock()
	fmt.Printf("main... 已经写锁定了\n")
	time.Sleep(2 * time.Second)
	rwm.Unlock()
	fmt.Printf("main... 写解锁\n")
}
