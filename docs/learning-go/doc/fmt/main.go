package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	manual()
}

func manual() {
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: print ----打印 print_note.go 内容----")
		fmt.Println("input: scan ----打印 scan_note.go 内容----")
		fmt.Println("-------------------------------")
		var flag bool
		fmt.Print("请输入显示内容的指令: ")
		_, err := fmt.Scanln(&prompt) // 返回成功扫描的条目个数和遇到的任何错误
		if err != nil {
			fmt.Println("fmt.Scanln(&prompt) err = ", err)
			break
		}
		switch prompt {
		case "print":
			PrintNote()
		case "scan":
			ScanNote()
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
