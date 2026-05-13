package main

import (
	"fmt"
)

func cmdWork() {
	fmt.Println("------------cmdWork()-------------")
	fmt.Println("工作空间：支持同时管理多个模块, 解决依赖关系")
	fmt.Println("在文件中声明工作空间, go.work 文件指定工作空间中每个模块的模块目录的相对路径, 当不存在 go.work 文件时, 工作区由包含当前目录的的单个模块组成")
	fmt.Println("go work 命令首先检查 GOWORK 环境变量来确定它是否在工作区上下文中,如果 GOWORK 设置为 off, 则该命令处于单模块上下文中")
	fmt.Println("如果为空或未提供, 该命令将搜索当前工作目录, 然后搜索后续父目录以查找 go.work")
	fmt.Println("如果找到文件, 该命令将在它定义的工作空间中运行, 否则, 工作区将仅包含工作目录的模块, 如果 GOWORK 命名以 .work 结尾的现有文件的路径, 将启用工作区模式。任何其他值都是错误的")
	fmt.Println("------------")
	fmt.Println("工作空间管理多个模块相当于本地私有模块, 可以任意修改, 而不需要使用 go get 命令下载的模块缓存中的模块版本")
	fmt.Println("------------")
	fmt.Println("go.work 文件指令")
	fmt.Println(tab, "1. go: 标识 go.work 文件要使用的 go 工具链版本")
	fmt.Println(tab, "2. use: 将磁盘中的模块添加到工作区, 参数是包含模块的 go.mod 文件的目录的相对路径, 不包含子模块(其参数目录的子目录中包含的模块)")
	fmt.Println(tab, "3. replace: 与 go.mod 文件中的 replace 指令类似, 将模块的特定版本或模块的所有版本的内容替换为其他位置的内容")
	fmt.Println("------------")
	fmt.Println("go work init [moddirs] 在当前目录中初始化一个新的工作空间并将传入的模块路径写入 go.work 文件")
	fmt.Println("go work use [-r] [moddirs] 将模块路径添加到 go.work 文件中")
	fmt.Println(tab, "-r 递归模式搜索指定模块路径中的模块")
	fmt.Println("go work sync 将工作空间的构建列表同步回工作区的模块")
	fmt.Println("go work edit [editing flags] [go.work] 在命令行窗口中编辑 go.work 文件")
	fmt.Println(tab, "editing flags")
	fmt.Println(tab, tab, "1. -fmt 重新格式化 go.work 文件而不进行其他更改, 此标志只能单独使用, go work edit -fmt")
	fmt.Println(tab, tab, "2. -replace, use, dropuse, dropreplace 指定模块的新模块与旧模块的替代, 重复命令, 根据传入的顺序执行")
	fmt.Println(tab, tab, "3. -go 设置 go.work 文件的 go 语言版本, go work edit -go=1.20")
	fmt.Println(tab, tab, "4. -print 文本格式输出 go.work 文件")
	fmt.Println(tab, tab, "5. -json JSON 格式化输出 go.work 文件")
	fmt.Println("------------")
	fmt.Println(`
	workspace
	|-- hello
		|-- hello.go
		|-- go.mod
		|-- go.sum
	|-- example
		|-- ..  // 其他包文件
		|-- stringutil
			|-- Reverse.go
			|-- ToUpper.go // 新增自定义方法
		|-- go.mod
		|-- go.sum
	|-- go.work`)

	fmt.Println("创建并初始化工作区: go work init")
	fmt.Println("在工作区目录下创建 example.com/hello 模块, hello.go 文件并写入内容")
	fmt.Println("在工作区任意目录下使用 go work use ./hello 命令将 hello 模块添加到工作区")
	fmt.Println("workspace 目录下克隆第三方模块 git clone https://go.googlesource.com/example")
	fmt.Println("工作区任意目录下使用 go work use ./example 命令将 example 模块添加到工作区")
	fmt.Println(`
	// 自动生成 go.work 文件
	go 1.19

	use (
		./example
		./hello
	)`)
	fmt.Println(`
	// 文件中添加导入 stringutil 包并重命名
	package main
	
	import (
		"fmt"
		stringUtil "golang.org/x/example/stringutil"
	)
	
	func main() {
		fmt.Println("传入参数 hello")
		fmt.Println("第三方模块的 stringutil 包默认提供的方法 Reverse")
		fmt.Printf("输出结果 \%\v \n", stringUtil.Reverse("hello"))
	}`)

	fmt.Println("在 hello 模块中使用 go mod tidy 命令更新模块依赖")

	fmt.Println("工作区任意目录下执行命令 go run example.com/hello 查看结果")
	fmt.Println("在 ./example/stringutil 目录下添加自定义方法 ToUpper.go 文件")
	fmt.Println(`
	package stringutil

	import "unicode"
	
	func ToUpper(s string) string {
		r := []rune(s)
		for i := range r {
			r[i] = unicode.ToUpper(r[i])
		}
		return string(r)
	}`)
	fmt.Println("在 hello.go 文件中添加调用 ToUpper 方法")
	fmt.Println(`
	fmt.Println("使用自己向第三方模块中添加的方法: ToUpper")
	fmt.Println("添加的方法在 ./example/stringutil/ToUpper.go 文件中")
	fmt.Printf("输出结果 \%\v \n", stringUtil.ToUpper("hello"))`)
	fmt.Println("工作区任意目录下执行命令 go run example.com/hello 查看结果")
}
