package main

import (
	"fmt"
)

func moduleDoc() {
	fmt.Println("目录结构")
	fmt.Println(tab, "bin: 存放编译后生成的二进制可执行文件")
	fmt.Println(tab, "pkg: 存放编译后生成的 .a 文件")
	fmt.Println(tab, "src: 存放项目的源代码, 可以是自己编写的代码, 也可以是使用 go get 下载的包, 这种管理方式称为 GOPATH 模式")
	fmt.Println("-------------------------------")

	fmt.Println("GOPATH mode")
	fmt.Println(tab, "1. go install 命令安装二进制库到 $GOBIN, 其默认路径为 $GOPATH/bin")
	fmt.Println(tab, "2. go install 命令安装编译好的包到 $GTH/pkg/ 下对应的平台目录中(由 GOOS 和 GOARCH 组合而成)")
	fmt.Println(tab, tab, "例如将 http://example.com/y/z 安装到 $GOPATH/pkg/http://example.com/y/z.a")
	fmt.Println(tab, "3. go get 命令下载源码包到 $GOPATH/src 中, 例如将 http://example.com/y/z 下载到 $GOPATH/src/example")
	fmt.Println(tab, "------------------")
	fmt.Println(tab, "go get 命令无法获取指定的版本")
	fmt.Println(tab, "所有源码和依赖包都存储在 GOPATH/src 目录下, 如果指定了多个 $GOPATH, 则在第一个 $GOPATH 的 src 下")
	fmt.Println(tab, "无法处理项目依赖库的版本问题")
	fmt.Println(tab, "无法处理依赖包的不同版本的的引用问题")
	fmt.Println(tab, "")
	fmt.Println("-------------------------------")

	fmt.Println("Go modules mode")
	fmt.Println(tab, "1. go 1.11(2018年8月)引入了 GO111MODULE 环境变量, 默认值为 auto")
	fmt.Println(tab, tab, " 当值为 off 时, go 命令将始终使用 GOPATH 模式")
	fmt.Println(tab, tab, " 当值为 on 时, go 命令将始终使用 Go modules 模式")
	fmt.Println(tab, tab, " 当值为 auto 或 不设置 时, go 命令行会根据当前工作目录来决定使用哪种模式, 如果当前目录在 $GOPATH/src 以外, 并且在根目录下存在 go.mod 文件, 那么 go 命令会启用 Go modules 模式,")
	fmt.Println(tab, tab, "     否则使用 GOPATH 模式, 这个规则保证了所有在 $GOPATH/src 中使用 auto 值时原有编译不受影响, 并且可以在其他目录中体验最新的 Go modules 模式")
	fmt.Println(tab, "2. go 1.13(2019年8月)调整了 GO111MODULE=auto 模式中对 $GOPATH/src 的限制, 如果一个代码库在 $GOPATH/src 中, 并且有 go.mod 文件存在, go 命令会启用 Go modules 模式, 这允许用户继续基于导入的层次结构中组织他们的检出代码, 但使用模块进行个别仓库的导入")
	fmt.Println(tab, "3. go 1.14 当主模块包含一个 vendor 目录, 并且它的 go.mod 文件指定 1.14 或更高版本时, 对于支持 -mod=vendor 的 go 命令, 将默认添加 -mod=vendor")
	fmt.Println(tab, "4. go 1.16(2021年2月)将 GO111MODULE=on 作为默认值, 默认启用 Go modules 模式, 即默认情况下 GOPATH 模式将被彻底关闭, 如果需要使用 GOPATH 模式则需要指定环境变量 GO111MODULE=auto 或 GO111MODULE=off")
	fmt.Println(tab, "5. go 1.NN(???) 将会废弃 GO111MODULE 环境变量 和 GOPATH 模式, 默认完全使用 Go modules 模式")
	fmt.Println("----------------")
	fmt.Println("module diagram illustrates -- hierarchy")
	fmt.Println("单个仓库管理一个模块")
	fmt.Println(`example.com/mymodule
    |-- LICENSE
    |-- go.mod
    |-- go.sum
    |-- package1
      |-- func1.go
      |-- func2.go
    |-- package2
      |-- func1.go
      |-- func2.go
	
import example.com/mymodule/package1 // 导入模块下的包
	`)

	fmt.Println("单个仓库管理多个模块")
	fmt.Println(tab, "模块路径: example.com/myrepo/module1")
	fmt.Println(tab, "版本号: module1/v1.2.3")
	fmt.Println(tab, "导入包: import example.com/myrepo/module1/package1")
	fmt.Println(tab, "用户 require 指令: example.com/myrepo/module1@module/v1.2.3 ")
	fmt.Println(`example.com/myrepo
    |-- module1   // 模块路径: example.com/myrepo/module1
      |-- LICENSE
      |-- go.mod
      |-- go.sum
      |-- package1
        |-- func1.go
        |-- func2.go
      |-- package2
        |-- func1.go
        |-- func2.go
    |-- module2 // 模块路径: example.com/myrepo/module2
      |-- LICENSE
      |-- go.mod
      |-- go.sum
      |-- package1
        |-- func1.go
        |-- func2.go
      |-- package2
        |-- func1.go
        |-- func2.go

import example.com/myrepo/module1/package1 // 导入模块下的包
import example.com/myrepo/module2/package1 // 导入模块下的包
	`)
}
