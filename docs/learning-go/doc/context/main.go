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
		fmt.Println("input: 1 | readme ----打印 doc.go 内容----")
		fmt.Println("input: 2 | withValue ----打印 withValue 内容----")
		fmt.Println("input: 3 | withTimeout ----打印 withTimeout 内容----")
		fmt.Println("input: 4 | withTimeoutCause ----打印 withTimeoutCause 内容----")
		fmt.Println("input: 5 | withDeadline ----打印 withDeadline 内容----")
		fmt.Println("input: 6 | withDeadlineCause ----打印 withDeadlineCause 内容----")
		fmt.Println("input: 7 | withCancel ----打印 withCancel 内容----")
		fmt.Println("input: 8 | withCancelCause ----打印 withCancelCause 内容----")
		fmt.Println("-------------------------------")
		var flag bool
		fmt.Print("请输入显示内容的指令: ")
		_, err := fmt.Scanln(&prompt) // 返回成功扫描的条目个数和遇到的任何错误
		if err != nil {
			fmt.Println("fmt.Scanln(&prompt) err = ", err)
			break
		}
		switch prompt {
		case "readme", "1":
			readme()
		case "withValue", "2":
			WithValueNote()
		case "withTimeout", "3":
			WithTimeoutNote()
		case "withTimeoutCause", "4":
			WithTimeoutCauseNote()
		case "withDeadline", "5":
			WithDeadlineNote()
		case "withDeadlineCause", "6":
			WithDeadlineCauseNote()
		case "withCancel", "7":
			WithCancelNote()
		case "withCancelCause", "8":
			WithCancelCauseNote()
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
