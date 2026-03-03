package main

import (
	"fmt"
	// "os"
	// "runtime/trace"
	"strings"
	"time"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("默认创建的 channel 是阻塞非缓冲的, 读写都是即时阻塞.")
	fmt.Println("缓冲 channel 自带一块缓冲区, 可以暂时存储数据, 如果缓冲区满了就会发生阻塞.")
	fmt.Println(tab, "当数据被发送到 channel 时会发生阻塞, 直到有其他的 goroutine 从该 channel 中读取数据.")
	fmt.Println(tab, "当从 channel 读取数据时, 读取也会被阻塞, 直到其他 goroutine 将数据写入该 channel")
	fmt.Println(tab, "for range 遍历 channel 时, 如果 channel 已经关闭, 则会正常遍历数据, 遍历完成后就会自动退出遍历")
	fmt.Println("--------------------")

	fmt.Println("1. 向   已关闭 的 chan 中 写入 数据时, 不管有无 缓冲区, 全部 panic")
	fmt.Println("2. 关闭 已关闭 的 chan(包括 nil chan), 不管有无 缓冲区, 全部 panic")
	fmt.Println("3. 写入/读取 nil chan 时, 永远挂起")
	fmt.Println("------------------------------------")
	time.Sleep(10 * time.Second)

	fmt.Println("单向通道")
	fmt.Println(tab, "chan<- int // 只写通道")
	fmt.Println(tab, "<-chan int // 只读通道")
	fmt.Println("------------------------------------")

	fmt.Println("通道接收数据的方法:")
	fmt.Println(tab, "1. 阻塞接收数据")
	fmt.Println(tab, tab, "data := <-ch")
	fmt.Println(tab, "2. 非阻塞接收数据")
	fmt.Println(tab, tab, "data, ok := <-ch // data 表示接收到的数据, 未接收到数据时, data 为通道类型的零值, ok 表示是否接收到数据")
	fmt.Println(tab, "3. 接收并忽略任意数据")
	fmt.Println(tab, tab, "<-ch")
	fmt.Println(tab, "4. 循环接收")
	fmt.Println(tab, tab, "for data := range ch { }")
	fmt.Println("------------------------------------")
	time.Sleep(5 * time.Second)

	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- 20
	allChan <- 30
	// allChan <- 40 // 向通道发送数据, 通道元素总数超过通道容量限制会报错 deadlock
	fmt.Printf("allChan := make(chan interface{}, 3) 的类型为 %T 长度为%d 容量为%d 地址为%p 值为 %v\n", allChan, len(allChan), cap(allChan), &allChan, allChan)
	<-allChan
	num := <-allChan
	<-allChan
	// <-allChan  // 从通道中取数据, 通道内没有数据后再次读取通道将会报错 dealock
	fmt.Printf("num 的类型为 %T 的值为 %v\n", num, num)
	fmt.Println("------------------------------------")

	fmt.Println("直接创建单向 channel 没有任何意义, 通常做法是创建双向 channel, 然后以单向 channel 的方式进行函数传递")
	fmt.Println("带缓冲的只写管道: var chan2 = make(chan<- int, 3)")
	var chan2 = make(chan<- int, 3)
	chan2 <- 20
	// <-chan2 // cannot receive from send-only channel chan2
	fmt.Println("----------------")
	fmt.Println("带缓冲的只读管道: var chan3 = make(<-chan int, 3)")
	// var chan3 = make(<-chan int, 3)
	// <-chan3
	// chan3 <- 20 //invalid operation: cannot send to receive-ony channel chan3
	fmt.Println("------------------------------------")

	selectNote()

	rangeNote()

	// fmt.Println("------parallelCalc() trace------")
	// // 生成 runtime 日志
	// // go tool trace parallel_calc_trace.out 浏览器打开 运行时日志
	// file, err := os.Create("parallel_calc_trace.out")
	// if err != nil {
	// 	fmt.Println("os.Create err =", err)
	// 	return
	// }
	// defer func() {
	// 	if err := file.Close(); err != nil {
	// 		fmt.Println("failed to close trace file: ", err)
	// 	}
	// }()

	// if err := trace.Start(file); err != nil {
	// 	fmt.Println("trace.Start() err =", err)
	// }
	// defer trace.Stop()
	// parallelCalc()
	// fmt.Println("------parallelCalc() trace------")

	sendChan()
}
