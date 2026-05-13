package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("reflect")
	manual()
}

func manual() {
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: 1 | readme ----打印 doc.go 内容----")
		fmt.Println("input: 2 | map ----打印 reflect_map.go 内容----")
		fmt.Println("input: 3 | new ----打印 reflect_new.go 内容----")
		fmt.Println("input: 4 | slice ----打印 reflect_slice.go 内容----")
		fmt.Println("input: 5 | func ----打印 reflect_func.go 内容----")
		fmt.Println("input: 6 | struct ----打印 reflect_struct.go 内容----")
		fmt.Println("input: 7 | chan ----打印 reflect_chan.go 内容----")
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
		case "2", "map":
			reflectMap()
		case "3", "new":
			reflectNew()
		case "4", "slice":
			reflectSlice()
		case "5", "func":
			reflectFunc()
		case "6", "struct":
			reflectStruct()
		case "7", "chan":
			reflectChan()
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
