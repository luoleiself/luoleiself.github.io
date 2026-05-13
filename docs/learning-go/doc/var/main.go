package main

import (
	fmt "fmt"
	"strings"

	interfacedoc "github.com/luoleiself/learning-go/var/interface"
	structdoc "github.com/luoleiself/learning-go/var/struct"
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
		fmt.Println("input: 1 | readme ----打印 README.go 内容----")
		fmt.Println("input: 2 | declare ----打印 declare.go 内容----")
		fmt.Println("input: str | string ----打印 string.go 内容----")
		fmt.Println("input: arr | array ----打印 array.go 内容----")
		fmt.Println("input: slice ----打印 slice.go 内容----")
		fmt.Println("input: map ----打印 map.go 内容----")
		fmt.Println("input: s | struct ----打印 struct 内容----")
		fmt.Println("input: inter | interface ----打印 interface 内容----")
		fmt.Println("input: 10 | switch ----打印 switch.go 内容----")
		fmt.Println("input: for | range ----打印 range.go 内容----")
		fmt.Println("input: func ----打印 func.go 内容----")
		fmt.Println("input: defer ----打印 defer.go 内容----")
		fmt.Println("input: panic ----打印 panic.go 内容----")
		fmt.Println("input: alias ----打印 type_alias.go 内容----")
		fmt.Println("input: c | convert | conversion ----打印 type_conversion.go 内容----")
		fmt.Println("input: g | generic ----打印 generic.go 内容----")
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
		case "2", "declare":
			declareNote()
		case "str", "string":
			stringNote()
		case "arr", "array":
			arrayNote()
		case "slice":
			sliceNote()
		case "map":
			mapNote()
		case "s", "struct":
			structdoc.Readme()
			structdoc.Instantiate()
			structdoc.EmptyStruct()
		case "inter", "interface":
			interfacedoc.Readme()
			interfacedoc.AssertionNote()
			interfacedoc.FuncImplInterfaceNote()
		case "10", "switch":
			switchNote()
		case "for", "range":
			rangeNote()
		case "func":
			funcNote()
		case "defer":
			deferNote()
		case "panic":
			panicNote()
		case "alias":
			typeAliasNote()
		case "c", "convert", "conversion":
			typeConversionNote()
		case "g", "generic":
			genericTypeNote()
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
