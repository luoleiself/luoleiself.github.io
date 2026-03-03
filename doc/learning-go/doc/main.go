package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("Golang")
	// var lp *fake
	// fred(nil)
	// fred(lp)
	// fmt.Println("------------------")
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: 1 | readme ----打印 readme 内容----")
		fmt.Println("input: 2 | sizeof ----打印 unsafe.Sizeof() 内容----")
		fmt.Println("input: 3 | memAlignment ----打印 内存对齐 内容----")
		fmt.Println("input: 4 | gc ----打印 gc.go 内容----")
		fmt.Println("input: 5 | goroutine ----打印 goroutine.go 内容----")
		fmt.Println("-------------------------------")
		var flag bool
		fmt.Print("请输入显示内容的指令: ")
		_, err := fmt.Scanln(&prompt) // 返回成功扫描的条目个数和遇到的任何错误
		if err != nil {
			fmt.Println("fmt.Scanln(&prompt) err = ", err)
			break
		}
		switch prompt {
		case "1", "readme":
			readme()
		case "2", "sizeof":
			unsafeSizeof()
		case "3", "memAlignment":
			memAlignment()
		case "4", "gc":
			GCNote()
		case "5", "goroutine":
			goroutine()
		default:
			if prompt != "0" && prompt != "exit" {
				fmt.Println("输入有误!")
			}
			flag = true
		}
		if flag {
			fmt.Println("------------Bye Bye-----------")
			break
		}
	}
}
