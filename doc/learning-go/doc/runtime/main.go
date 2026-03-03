package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("func CPUProfile() []byte // 获取 cpu 信息")
	fmt.Println("func GOMAXPROCS(n int) int // 设置当前 Go 程序可用的 CPU 的核心数并返回之前的设置, n < 1 则不做任何修改")
	fmt.Println("func GOROOT() string // 返回 Go 的根目录")
	fmt.Println("func Goexit() // 终止调用它的go程. 其它go程不会受影响. Goexit会在终止该go程前执行所有defer的函数")
	fmt.Println("func GC() // 执行垃圾回收器")
	fmt.Println("func Gosched() // 使当前go程放弃处理器, 以让其它go程运行. 它不会挂起当前go程, 因此当前go程未来会恢复执行")
	fmt.Println("func NumCPU() int // 返回本地机器的逻辑CPU个数")
	fmt.Println("func NumCgoCall() int64 // 返回当前进程执行的cgo调用次数")
	fmt.Println("func NumGoroutine() int // 返回当前存在的Go程数")
	fmt.Println("func Version() string // 获取 Go 的版本")
	fmt.Println("func Breakpoint() // 执行一个断点陷阱")
	fmt.Println("------------------------------------")

	fmt.Println("CPUProfile =", runtime.CPUProfile())
	fmt.Println("NumCPU =", runtime.NumCPU())
	fmt.Println("NumGoroutine =", runtime.NumGoroutine())
	fmt.Println("GOROOT =", runtime.GOROOT())
	fmt.Println("NumCgoCall =", runtime.NumCgoCall())
	fmt.Println("Version =", runtime.Version())
}
