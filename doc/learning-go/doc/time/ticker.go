package main

import (
	"fmt"
	"sync"
	"time"
)

func TickerNote() {
	fmt.Println("--------------ticker 周期计时器---------------")
	fmt.Println("func Tick(d Duration) <-chan Time // Tick是NewTicker的封装, 只提供对Ticker的通道的访问. 如果不需要关闭Ticker, 本函数就很方便")
	fmt.Println("func NewTicker(d Duration) *Ticker // 返回一个新的 Ticker, 该 Ticker 包含一个通道字段, 并会每隔时间段d就向该通道发送当时的时间.")
	fmt.Println(tab, "它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者. 如果d<=0会panic. 关闭该Ticker可以释放相关资源")
	fmt.Println(tab, "func (t *Ticker) Reset(d Duration) // Reset停止周期计时器并重置为指定的持续时间, 下一个 Tick 将在新时段过后到达")
	fmt.Println(tab, "func (t *Ticker) Stop() // Stop关闭一个Ticker. 在关闭后, 将不会发送更多的tick信息. Stop不会关闭通道t.C, 以避免从该通道的读取不正确的成功")
	fmt.Println("-----------------------------")

	var wg sync.WaitGroup
	wg.Add(40)

	fmt.Println(`方式1: time.NewTicker(d Duration) *Ticker
t1 := time.NewTicker(time.Millisecond * 800) // 创建一个 800 毫秒触发的周期计时器
done := make(chan bool)
go func() {
    done <- true
}()
num := 1
for {
    select {
    case <-done:
        fmt.Println("Done!")
    case t := <-t1.C:
      fmt.Printf("Mode 1 current time %\v\n\n", t) // 每 800 毫秒触发一次
      num += 1
      if num == 6 {
        fmt.Printf("当输出信息超过 %\d 次则重置计时器持续时间\n", num)
        t1.Reset(time.Millisecond * 100) // 重置计时器持续时间
      }
    }
}`)
	fmt.Println("---------------")
	go func() {
		t1 := time.NewTicker(time.Millisecond * 800) // 创建一个 800 毫秒触发的周期计时器
		done := make(chan bool)
		go func() {
			done <- true
		}()
		num := 1
		for {
			select {
			case <-done:
				fmt.Println("Done!")
			case t := <-t1.C:
				fmt.Printf("Mode 1 current time %v\n\n", t) // 每 800 毫秒触发一次
				num += 1
				if num == 6 {
					fmt.Printf("当输出信息超过 %d 次则重置计时器持续时间\n", num)
					t1.Reset(time.Millisecond * 100) // 重置计时器持续时间
				}
				wg.Done()
			}
		}
	}()

	go func() {
		fmt.Println(`方式2: <-time.Tick(d Duration) <-chan Time
// go-staticcheck, should use for range instead of for { select {} }
for {
    select {
    case t := <-time.Tick(time.Millisecond * 300):
        fmt.Printf("Mode 2 current time %\v\n\n", t) // 每 300 毫秒触发一次
        wg.Done()
    }
}

tickChan := time.Tick(time.Millisecond * 300)
for t := range tickChan {
	fmt.Printf("Mode 2 current time \%\v\n\n", t) // 每 300 毫秒触发一次
	wg.Done()
}`)
		fmt.Println("---------------")
		// go-staticcheck, should use for range instead of for { select {} }
		// for {
		// 	select {
		// 	case t := <-time.Tick(time.Millisecond * 300):
		// 		fmt.Printf("Mode 2 current time %v\n\n", t) // 每 300 毫秒触发一次
		// 		wg.Done()
		// 	}
		// }
		tickChan := time.Tick(time.Millisecond * 300)
		for t := range tickChan {
			fmt.Printf("Mode 2 current time %v\n\n", t) // 每 300 毫秒触发一次
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("-----------------------------")
}
