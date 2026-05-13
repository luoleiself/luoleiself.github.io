package main

import (
	"fmt"
	"sync"
	"time"
)

func rangeNote() {
	fmt.Println("------------rangeNote()------------")

	fmt.Println(`
  var ch = make(chan int, 10)
  defer close(ch)
  go func() {
    for i := 1; i <= 30; i++ {
      ch <- i
    }
  }()
  go func() {
    for c := range ch {
      fmt.Println("goroutine use \033[1;32mrange\033[0m chan get val \033[1;32m", c, "\033[0m")
    }
  }()
  time.Sleep(time.Second * 2)`)
	fmt.Println("----------------")
	var ch = make(chan int, 10)
	defer close(ch)
	go func() {
		for i := 1; i <= 30; i++ {
			ch <- i
		}
	}()
	go func() {
		for c := range ch {
			fmt.Println("goroutine use \033[1;32mrange\033[0m chan get val \033[1;32m", c, "\033[0m")
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("----------------")

	fmt.Println("使用信号向量限制 goroutine 的数量")
	// 大量使用 goroutine 会拖慢运行速度
	// for _, item := range bigList {
	// 	go process(item)
	// }
	// 使用信号向量限制 goroutine 的数量
	// sem := make(chan struct{}, 100)
	// for _, item := range bigList {
	// 	sem <- struct{}{}
	// 	go func(i Item) {
	// 		defer func() { <-sem }()
	// 		process(i)
	// 	}(item)
	// }

	fmt.Println("------------------------------------")
}

// 合并函数(组件), 把多个 channel 中的数据发送到一个 channel 中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// 把一个 channel 中的数据发送到 out 中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	// 指定线程等待的计数
	wg.Add(len(ins))

	// 扇入, 需要启动多个 goroutine 用于处理多个 channel 中的数据
	for _, cs := range ins {
		go p(cs)
	}

	wg.Wait()
	close(out)

	return out
}
