package main

import (
	"fmt"
	"sync"
	"time"
)

func sendChan() {
	fmt.Println("-----------sendChan()---------------")
	fmt.Println(`
  func sendChan() {
    var ch = make(chan int, 3)
    var len = 40
    var wg sync.WaitGroup
    defer close(ch)            // 延迟当前的调用栈即将返回时执行 关闭 管道, 多个 defer 逆序执行
    go receiveChan(ch, &wg, 1) // 启动协程读取管道数据
    go receiveChan(ch, &wg, 2) // 启动协程读取管道数据
    go receiveChan(ch, &wg, 3) // 启动协程读取管道数据
    go receiveChan(ch, &wg, 4) // 启动协程读取管道数据
    for i := 1; i <= len; i++ {
      wg.Add(1)
      ch <- i
      fmt.Println("\033[1;33mch send val\033[0m \033[1;36m", i, "\033[0m")
    }
    wg.Wait()
  }

  // 声明一个 channel 方法, 第一个参数为 只读 channel
  func receiveChan(ch <-chan int, wg *sync.WaitGroup, num int) {
    for {
      select {
      case v := <-ch:
        fmt.Printf("\033[1;32mgoroutine\033[0m \033[1;44;31m%\d\033[0m \033[4;47;30mreceive\033[0m \033[0;32mval\033[0m \033[0;35m%\v\033[0m\n", num, v)
        time.Sleep(time.Millisecond * 300) // 协程等待200毫秒
        wg.Done()
      }
      // if v, ok := <-ch; ok {
      //   fmt.Printf("goroutine %\d receive %\v\n", num, v)
      //   time.Sleep(time.Millisecond * 300) // 协程等待200毫秒
      //   wg.Done()
      // }
    }
  }`)
	fmt.Println("----------------")
	var ch = make(chan int, 3)
	var len = 40
	var wg sync.WaitGroup
	defer close(ch)            // 延迟当前的调用栈即将返回时执行 关闭 管道, 多个 defer 逆序执行
	go receiveChan(ch, &wg, 1) // 启动协程读取管道数据
	go receiveChan(ch, &wg, 2) // 启动协程读取管道数据
	go receiveChan(ch, &wg, 3) // 启动协程读取管道数据
	go receiveChan(ch, &wg, 4) // 启动协程读取管道数据
	for i := 1; i <= len; i++ {
		wg.Add(1)
		ch <- i
		fmt.Println("\033[1;33mch send val\033[0m \033[1;36m", i, "\033[0m")
	}
	wg.Wait()
}

// 声明一个 channel 方法, 第一个参数为 只读 channel
func receiveChan(ch <-chan int, wg *sync.WaitGroup, num int) {
	// for {
	// 	select {
	// 	case v := <-ch:
	// 		fmt.Printf("\033[1;32mgoroutine\033[0m \033[1;44;31m%d\033[0m \033[4;47;30mreceive\033[0m \033[0;32mval\033[0m \033[0;35m%v\033[0m\n", num, v)
	// 		time.Sleep(time.Millisecond * 300) // 协程等待200毫秒
	// 		wg.Done()
	// 	}
	// 	// if v, ok := <-ch; ok {
	// 	// 	fmt.Printf("\033[1;32mgoroutine\033[0m \033[1;44;31m%d\033[0m \033[4;47;30mreceive\033[0m \033[0;32mval\033[0m \033[0;35m%v\033[0m\n", num, v)
	// 	// 	time.Sleep(time.Millisecond * 300) // 协程等待200毫秒
	// 	// 	wg.Done()
	// 	// }
	// }
	for v := range ch {
		fmt.Printf("\033[1;32mgoroutine\033[0m \033[1;44;31m%d\033[0m \033[4;47;30mreceive\033[0m \033[0;32mval\033[0m \033[0;35m%v\033[0m\n", num, v)
		time.Sleep(time.Millisecond * 300) // 协程等待200毫秒
		wg.Done()
	}
}
