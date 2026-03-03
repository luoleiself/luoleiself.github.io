package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("基础知识")
	manual()
}
func manual() {
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: readme ----打印 regexp.go 内容----")
		fmt.Println("input: compile ----打印 compile.go 内容----")
		fmt.Println("input: find ----打印 find.go 内容----")
		fmt.Println("input: match ----打印 match.go 内容----")
		fmt.Println("input: replace ----打印 replace.go 内容----")
		fmt.Println("-------------------------------")
		var flag bool
		fmt.Print("请输入显示内容的指令: ")
		_, err := fmt.Scanln(&prompt) // 返回成功扫描的条目个数和遇到的任何错误
		if err != nil {
			fmt.Println("fmt.Scanln(&prompt) err = ", err)
			break
		}
		switch prompt {
		case "readme":
			Readme()
		case "compile":
			CompileNote()
		case "find":
			FindNote()
		case "match":
			MatchNote()
		case "replace":
			ReplaceNote()
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
