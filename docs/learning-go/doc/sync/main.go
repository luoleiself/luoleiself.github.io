package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("sync")
	manual()
}

func manual() {
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: 1 | readme ----打印 README.go 内容----")
		fmt.Println("input: 2 | mutex  ----打印 Mutex_demo.go 内容----")
		fmt.Println("input: 3 | rwm ----打印 RWMutex_demo.go 内容----")
		fmt.Println("input: 4 | wg ----打印 WaitGroup_demo.go 内容----")
		fmt.Println("input: 5 | once ----打印 Once_demo.go 内容----")
		fmt.Println("input: 6 | cond	----打印 Cond_demo.go 内容----")
		fmt.Println("input: 7 | pool	----打印 pool_demo.go 内容----")
		fmt.Println("input: 8 | p2	----打印 pool_demo_2.go 内容----")
		fmt.Println("input: 9 | sm	----打印 sync_Map.go 内容----")
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
		case "2", "mutex":
			MutexNote()
		case "3", "rwm":
			RWMutexNote()
		case "4", "wg":
			WaitGroupNote()
		case "5", "once":
			OnceNote()
		case "6", "cond":
			CondNote()
		case "7", "pool":
			PoolNote()
		case "8", "p2":
			PoolNote2()
		case "9", "sm":
			syncMapNote()
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
