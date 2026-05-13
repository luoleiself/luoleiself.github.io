package main

import (
	"fmt"
)

func cmdEV() {
	fmt.Println("----------------cmdEV()----------------")
	fmt.Println("go env	输出和配置 go 环境变量信息, 见 environment.go 文件")
	fmt.Println(tab, "-u    重置 go 的环境变量")
	fmt.Println(tab, "-json 输出 go 环境变量的 JSON 格式")
	fmt.Println(tab, "-w NAME=VALUE 设置 go 环境变量")
	fmt.Println("---------------------------------------")

	fmt.Println("go version 打印 go 可执行文件的构建信息")
	fmt.Println(tab, "-m	打印每个可执行文件的嵌入模块版本信息(如果可用), 在输出中, 模块信息由版本行后面的多行组成, 每行由一个前导制表符缩进")
	fmt.Println(tab, "-v	打印无法识别的文件")
	fmt.Println(tab, "")
}
