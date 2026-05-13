package main

import (
	"fmt"
	"runtime"
	"time"
)

type Vector []int64

func (v Vector) DoAll() (sum int64) {
	NCPU := runtime.NumCPU()
	c := make(chan int64, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.DoSum(i*len(v)/NCPU, (i+1)*len(v)/NCPU, c)
	}

	for i := 0; i < NCPU; i++ {
		sum += <-c
	}

	return
}
func (v Vector) DoSum(i, n int, c chan int64) {
	var sum int64 = 0
	for ; i < n; i++ {
		sum += v[i]
	}
	c <- sum
}

func parallelCalc() {
	fmt.Println("--------多核并发计算数值累加结果--------")
	runtime.GOMAXPROCS(2) // 设置可用 CPU 的最大核心数
	fmt.Println("NumCPU =", runtime.NumCPU())
	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
	var start = time.Now().Unix()
	eg := make(Vector, 0, 10)
	for i := 0; i < 100000000; i++ {
		eg = append(eg, int64(i))
	}
	var sum = eg.DoAll() // 累加和
	var end = time.Now().Unix()
	fmt.Println("sum=", sum)
	fmt.Println(len(eg), "start", start, "now", end, "dur:", end-start)
}
