package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args 是一个字符串切片, 保存了命令行参数, 第一个是程序名", os.Args)
	fmt.Println("flag 包解析命令行参数")
	fmt.Println("IntVar 和 Int, Int64 和 Int64Var, Uint 和 UintVar, Uint64 和 Uint64Var, Float64 和 Float64Var, Bool 和 BoolVar, StringVar 和 String, 方法使用方式一致")
	fmt.Println("flag.Parse() 从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行")
	fmt.Println("flag.Parsed() 返回是否Parse已经被调用过")
	fmt.Println("-------------------")

	var port int
	flag.IntVar(&port, "p", 8080, "端口号")
	userName := flag.String("u", "root", "用户名")
	fmt.Println("flag.IntVar(&port, \"p\", 8080, \"端口号\") 用指定的名称、默认值、使用信息注册一个int类型flag, 并将flag的值保存到名称指向的变量")
	fmt.Println("userName := flag.String(\"u\", \"root\", \"用户名\") 用指定的名称、默认值、使用信息注册一个string类型flag, 并返回保存了该flag的值的指针")
	fmt.Println("flag.Parsed() // 返回是否Parse已经被调用过", flag.Parsed()) // 返回是否Parse已经被调用过
	flag.Parse()
	fmt.Printf("username: %s, port: %d\n", *userName, port)
	fmt.Println("------------------------------------")

	fmt.Println("flag.Args() []string", "返回解析之后剩下的非 flag 参数(不包括命令名), []string")
	fmt.Println("flag.Arg(i int) string", "返回解析之后的第 i 个参数, 索引从 0 开始")
	fmt.Println("flag.NArg() int", "返回解析 flag 之后剩余参数的个数")
	fmt.Println("flag.NFlag() int", "返回解析时进行了设置的 flag 的数量")
	fmt.Println("------------------------------------")
}
