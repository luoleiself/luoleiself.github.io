package main

import (
	"fmt"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("package_env_CMD")
	manual()
}

func manual() {
	var prompt string
	for {
		fmt.Println("-----------API手册------------")
		fmt.Println("input: 0 | exit ----退出----")
		fmt.Println("input: 1 | readme ----打印 README.go 内容----")
		fmt.Println("input: 2 | build | run ----打印 cmd_build_run.go 内容----")
		fmt.Println("input: 3 | clean | fix | generate | vet ----打印 cmd_clean_fix_generate_vet.go 内容----")
		fmt.Println("input: 4 | doc | fmt ----打印 cmd_doc_fmt.go 内容----")
		fmt.Println("input: 5 | env | version ----打印 cmd_env_version.go 内容----")
		fmt.Println("input: 6 | get | install | list	----打印 cmd_get_install_list.go 内容----")
		fmt.Println("input: 7 | mod	----打印 cmd_mod.go 内容----")
		fmt.Println("input: 8 | work	----打印 cmd_work.go 内容----")
		fmt.Println("input: 9 | test | tool	----打印 cmd_test_tool.go 内容----")
		fmt.Println("input: 10 | environ	----打印 environment.go 内容----")
		fmt.Println("input: 11 | package	----打印 package.go 内容----")
		fmt.Println("input: 12 | internal	----打印 internal.go 内容----")
		fmt.Println("input: 13 | module	----打印 module.go 内容----")
		fmt.Println("-------------------------------")
		var flag bool
		fmt.Print("请输入显示内容的指令: ")
		_, err := fmt.Scanln(&prompt) // 返回成功扫描的条目个数和遇到的任何错误
		if err != nil {
			fmt.Println("fmt.Scanln(&prompt) err = ", err)
			break
		}
		switch prompt {
		// case "1", "readme":
		// 	readme()
		case "2", "build", "run":
			cmdBR()
		case "3", "clean", "fix", "generate", "vet":
			cmdCFGV()
		case "4", "doc", "fmt":
			cmdDF()
		case "5", "env", "version":
			cmdEV()
		case "6", "get", "install", "list":
			cmdGIL()
		case "7", "mod":
			cmdMod()
		case "8", "work":
			cmdWork()
		case "9", "test", "tool":
			cmdTT()
		case "10", "environ":
			environmentVariable()
		case "11", "package":
			pkgDoc()
		case "12", "internal":
			cmdInternalPkg()
		case "13", "module":
			moduleDoc()
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
