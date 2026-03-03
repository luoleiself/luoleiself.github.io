package main

import (
	"fmt"
)

func panicNote() {
	fmt.Println("------------panicNote()-------------")
	fmt.Println("Go 语言的类型系统只能捕获编译时的错误, 有些运行时的错误(数组访问越界, 空指针异常等)无法捕获")
	fmt.Println("程序运行时 panic 导致程序退出之前立即逆序执行该 goroutine 中被 defer 的函数 ")
	fmt.Println("-------------------------")

	fmt.Println("内置函数 panic 手动触发: panic(\"crash\") // panic 函数后面的代码不会被执行")
	fmt.Println("内置函数 recover 可以让进入 panic 流程中的 goroutine 恢复过来, recover 内置函数只能运行在 defer 中")
	fmt.Println("panic 和 recover 的组合有如下特性")
	fmt.Println(tab, "有 panic 没 recover, 程序宕机")
	fmt.Println(tab, "有 panic 也有 recover, 程序不会宕机, 执行完对应的 defer 后, 从宕机点退出当前函数后继续执行")
	fmt.Println("-------------------------")

	defer func() {
		fmt.Println("defer func1...")
	}()

	// recover
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover func exec... ", err)
		}
	}()

	// panic("crash") // 手动触发 panic

	// var slice []int // panic 函数后面的代码不会被执行
	// fmt.Println(slice[4])

	defer func() {
		fmt.Println("defer func2...")
	}()
}
