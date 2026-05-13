package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"strings"
	"sync"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("func Start(w io.Writer) error // 启动 trace")
	fmt.Println("func Stop() // 停止 trace")
	fmt.Println("func IsEnabled() bool // 判断是否开启 trace")
	fmt.Println("func Log(ctx context.Context, category, message string) // 触发具有给定类别和消息的一次性事件")
	fmt.Println("func Logf(ctx context.Context, category, format string, args ...interface{})")
	fmt.Println("------------------------------------")

	fmt.Println("STW 是gc中的两个\"停止世界\"的阶段. 在这两个阶段中, goroutine会停止运行")
	fmt.Println("GC(idle) 指没有标记内存时的goroutine")
	fmt.Println("MARK ASSIST 在分配内存过程中重新标记内存(mark the memory)的goroutine")
	fmt.Println("WEEP 垃圾清理")
	fmt.Println("GXX runtime.gcBgMarkWorker 是帮助标记内存的专用后台goroutine")
}

func traceNote() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	// 2. trace绑定文件句柄
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// 下面就是你的监控的程序
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			_ = make([]int, 0, 20000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			_ = make([]int, 0, 10000)
		}
	}()

	wg.Wait()
}
