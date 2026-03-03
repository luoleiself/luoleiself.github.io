package main

import (
	"fmt"
)

func cmdBR() {
	fmt.Println("------------------cmdBR()----------------")
	fmt.Println("go build	编译由导入路径命名的包及其依赖项, 但不安装结果, 编译时会忽略 _test.go 结尾的文件")
	fmt.Println(tab, "1. 当参数不为空时")
	fmt.Println(tab, tab, "1. 如果 fileName 为同一 main 包下的所有源文件名(可能有一个或者多个), 编译器将生成一个与第一个 fileName 同名的可执行文件")
	fmt.Println(tab, tab, tab, "eg: go build abc.go def.go ... 将生成一个 abc.exe(扩展名和系统有关) 可执行文件")
	fmt.Println(tab, tab, "2. 如果 fileName 为非 main 包下的源文件名, 编译器将只对该包进行语法检查, 不生成可执行文件")
	fmt.Println(tab, "2. 当参数为空时")
	fmt.Println(tab, tab, "1. 如果当前目录存在 main 包, 则会生成一个与当前目录名同名的 \"目录名.exe(扩展名和系统有关)\" 可执行文件")
	fmt.Println(tab, tab, tab, "eg: /hello/$ go build  将生成 hello.exe(扩展名和系统有关) 可执行文件")
	fmt.Println(tab, tab, "2. 如果当前目录不存在 main 包, 则只对当前目录下的源文件进行语法检查, 不生成可执行文件")
	fmt.Println(tab, "--------参数--------")
	fmt.Println(tab, "-a	强制重新生成已经是最新的包")
	fmt.Println(tab, "-o  强制 build 生成的可执行文件或对象写入到指定的输出文件或目录")
	fmt.Println(tab, "-n	打印命令但不运行它们")
	fmt.Println(tab, "-p n	并行执行构建或者测试命令时的 cpu 数量, 默认 GOMAXPROCS")
	fmt.Println(tab, "-v	打印编译的包名")
	fmt.Println(tab, "-x	打印编译期间所用到的其它命令")
	fmt.Println(tab, "-race 启用资源竞争检测")
	fmt.Println(tab, "-work	打印临时工作目录的名称, 退出时并不删除它们")
	fmt.Println(tab, "-mod mode	 模块下载包的模式使用: readonly, vendor, mod, 默认: 如果存在 vendor 目录, go.mod 中 go 版本在 1.14 或者更高, 则使用 -mod=vendor, 其他情况使用 -mod-readonly")
	fmt.Println(tab, "-pkg dir	从指定的位置安装和载入所有的包")
	fmt.Println(tab, "-toolexec 'cmd args'	通过链式调用触发指定的命令")
	fmt.Println(tab, "")
	fmt.Println("---------------------------------------")

	fmt.Println("go run 编译并运行命令的主 go 包, 并把编译后的可执行文件存放到临时工作目录")
	fmt.Println(tab, "参数只能为源码文件, 不接受代码包和测试文件")
	fmt.Println(tab, "-exec 如果未给出 -exec 标志, 则 GOOS 或 GOARCH 与系统默认值不同")
}
