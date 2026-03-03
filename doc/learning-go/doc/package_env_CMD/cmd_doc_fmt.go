package main

import (
	"fmt"
)

func cmdDF() {
	fmt.Println("------------------cmdDF()----------------")
	fmt.Println("go doc	打印由其参数(package, const, func, type, var, method, struct field)标识的项关联的文档注释")
	fmt.Println(tab, "go doc strings # 输出 strings 包的相关文档信息")
	fmt.Println(tab, "-all	显示指定包的所有文档信息")
	fmt.Println(tab, "-c 	匹配符号时注意大小写")
	fmt.Println(tab, "-short	每个符号用一条线表示")
	fmt.Println(tab, "-u	显示未导出的符号, 方法, 字段的文档信息")
	fmt.Println("---------------------------------------")

	fmt.Println("go fmt	格式化包中的源码")
	fmt.Println(tab, "-n	打印命令将要执行")
	fmt.Println(tab, "-x	在执行命令时打印命令")
}
