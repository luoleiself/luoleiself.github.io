package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func selectNote() {
	fmt.Println("------------selectNote()------------")
	fmt.Println("golang 中的 select 就是用来监听和 channel 有关的 IO 操作, 当 IO 操作发生时, 触发相应的动作. select 只能应用于 channel 的操作, 既可以用于 channel 的数据接收, 也可以用于 channel 的数据发送。如果 select 的多个分支都满足条件, 则会随机的选取其中一个满足条件的分支执行")
	fmt.Println("select 的用法与 switch 语言非常类似, 由 select 开始一个新的选择块, 每个选择条件由 case 语句来描述, 最大的区别是每个 case 语句里必须是一个 IO 操作")
	fmt.Println(`select {
    case n := <- chan1:
        // 如果 chan1 成功读到数据, 则进行该 case 处理语句
    case chan2 <- 1:
        // 如果成功向 chan2 写入数据, 则进行该 case 处理语句
    default:
        // 如果上面都没有成功, 则进入 default 处理流程
  }`)
	fmt.Println("------------------------------------")

	fmt.Println("伪随机数: 计算机底层生成一个数值, 其根源仍然是程序根据某种算法得到的数值, 认为操控的计算就一定有规律可循, 只是这个规律不是肉眼可见的, 而不是真正的无规则可循的")
	fmt.Println(`
  // 设置随机数种子
  rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
  var len = 20
  var ch = make(chan int)
	wg.Add(2)
  defer close(ch)
  go RandChan1(ch, &wg, &len)
  go RandChan2(ch, &wg, &len)
  wg.Wait()

  func RandChan1(ch chan int, wg *sync.WaitGroup, len *int) {
		defer wg.Done()
    for {
        n := rand.Intn(127) // 生成一个 0 - 127 之间的伪随机数
        select {
        case val := <-ch:
            fmt.Println("RandChan \033[1;32m1\033[0m \033[1;31mread\033[0m \033[1;32m", val, "\033[0m")
        case ch <- n:
            fmt.Println("RandChan \033[1;32m1\033[0m \033[1;34mwrite\033[0m \033[1;32m", n, "\033[0m")
            *len -= 1
        }
        if *len == 0 {
            break
        }
    }
  }`)
	fmt.Println("----------------")
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var len = 20
	var ch = make(chan int)
	wg.Add(2)
	defer close(ch)
	go RandChan1(ch, &wg, &len)
	go RandChan2(ch, &wg, &len)
	wg.Wait()
}

func RandChan1(ch chan int, wg *sync.WaitGroup, len *int) {
	for {
		n := rand.Intn(127) // 生成一个 0 - 127 之间的伪随机数
		select {
		case val := <-ch:
			fmt.Println("RandChan \033[1;32m1\033[0m \033[1;31mread\033[0m \033[1;32m", val, "\033[0m")
		case ch <- n:
			fmt.Println("RandChan \033[1;32m1\033[0m \033[1;34mwrite\033[0m \033[1;32m", n, "\033[0m")
			*len -= 1
		}
		if *len == 0 {
			wg.Done()
			break
		}
	}
}

func RandChan2(ch chan int, wg *sync.WaitGroup, len *int) {
	for {
		n := rand.Intn(127) // 生成一个 0 - 127 之间的伪随机数
		select {
		case val := <-ch:
			fmt.Println("RandChan \033[1;35m2\033[0m \033[1;31mread\033[0m \033[1;35m", val, "\033[0m")
		case ch <- n:
			fmt.Println("RandChan \033[1;35m2\033[0m \033[1;34mwrite\033[0m \033[1;35m", n, "\033[0m")
			*len -= 1
		}
		if *len == 0 {
			wg.Done()
			break
		}
	}
}
