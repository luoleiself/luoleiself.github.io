package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func WaitGroupNote() {
	fmt.Println("-------------WaitGroupNote()-------------")
	fmt.Println("同步等待组")
	var wg sync.WaitGroup
	// fmt.Println("wg ", wg)
	wg.Add(3) // 设定应等待的线程的数量
	// rand.Seed(time.Now().UnixNano()) // 根据纳秒产生随机数 // 1.20 deprecated
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("random int is %d\n", random.Int())
	fmt.Printf("random int32 is %d\n", random.Int31())
	fmt.Printf("random int64 is %d\n", random.Int63())
	fmt.Printf("random unit32 is %d\n", random.Uint32())
	fmt.Printf("random float64 is %f\n", random.Float64())
	fmt.Printf("random Perm is %v\n", random.Perm(12))
	go printNum(&wg, 1) // 启动 goroutine
	go printNum(&wg, 2) // 启动 goroutine
	go printNum(&wg, 3) // 启动 goroutine
	wg.Wait()           // 进入阻塞状态, 当计数为 0 时解除阻塞
	defer fmt.Println("main over...")
}

func printNum(wg *sync.WaitGroup, num int) {
	for i := 1; i <= 3; i++ {
		// 在每个 goroutine 前面添加多个制表符方便观看打印结果
		pre := strings.Repeat("\t", num-1)
		fmt.Printf("%s 第 %d 号子 goroutine, %d \n", pre, num, i)
		time.Sleep(time.Second)
	}
	wg.Done() // 计数减1
}
