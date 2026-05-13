package main

import (
	"fmt"
	"sync"
	"time"
)

func CondNote() {
	fmt.Println("-------------CondNote()-------------")

	var m1 sync.Mutex
	var c1 = sync.NewCond(&m1)
	for i := 1; i <= 10; i++ {
		go func(x int) {
			c1.L.Lock()                           // 加锁
			defer c1.L.Unlock()                   // 延迟解锁
			c1.Wait()                             // 自行解锁 c.L 并阻塞当前线程等待被唤醒
			fmt.Println("goroutine ", x, "awake") // c 被唤醒时执行
		}(i)
	}
	time.Sleep(time.Second * 2) // 等待 2 秒后, 唤醒 1 个 goroutine
	fmt.Println("等待 2 秒后, 唤醒 1 个 goroutine")
	c1.Signal()
	time.Sleep(time.Second * 3) // 等待 3 秒后, 唤醒 1 个 goroutine
	fmt.Println("等待 3 秒后, 唤醒 1 个 goroutine")
	c1.Signal()
	time.Sleep(time.Second * 4)
	fmt.Println("等待 4 秒后, 广播唤醒剩余 goroutine")
	c1.Broadcast()
	time.Sleep(time.Second)
	fmt.Println("-----------------")

	var cg1 = make(chan bool, 1)
	var cg2 = make(chan bool, 1)
	defer close(cg1) // 延时释放通道
	defer close(cg2) // 延时释放通道
	var c2 = sync.NewCond(&sync.Mutex{})
	go func() {
		c2.L.Lock()         // 加锁
		defer c2.L.Unlock() // 延迟解锁

		c2.Wait() // 自行解锁 c.L 并阻塞当前线程等待被唤醒

		fmt.Println("goroutine 1 awake")
		cg1 <- true // 向通道写入数据
	}()
	go func() {
		c2.L.Lock()         // 加锁
		defer c2.L.Unlock() // 延迟解锁

		c2.Wait() // 自行解锁 c.L 并阻塞当前线程等待被唤醒

		fmt.Println("goroutine 2 awake")
		cg2 <- true // 向通道写入数据
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("等待 2 秒后, 广播唤醒所有 goroutine")
	c2.Broadcast() // 广播唤醒所有等待 c 的 goroutine
	<-cg1          // 阻塞通道等待写入数据
	<-cg2          // 阻塞通道等待写入数据
	fmt.Println("-----------------------------------")
}
