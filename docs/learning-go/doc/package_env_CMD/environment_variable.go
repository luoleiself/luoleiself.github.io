package main

import (
	"fmt"
)

func environmentVariable() {
	fmt.Println("---------environmentVariable()-----------")
	fmt.Println("环境变量--部分")
	fmt.Println(tab, "GOROOT	# go 命令安装目录")
	fmt.Println(tab, "GO111MODULE	# 控制 go 命令是以模块感知还是 GOPATH 模式运行, 值取 off, on, auto")
	fmt.Println(tab, "GOPATH    # 使用 GOPATH 模式时, go 项目目录以及 import 语句的解析路径, 默认 $HOME/go")
	fmt.Println(tab, "GOBIN   # go install 安装命令的目录, 默认 $GOPATH/bin")
	fmt.Println(tab, "GOMODCACHE    # go 命令下载存储的模块的目录, 默认 $GOPATH/pkg/mod")
	fmt.Println(tab, "GOCACHE   # go 命令存储构建信息以备将来构建使用, 默认 $HOME/.cache/go-build")
	fmt.Println(tab, "GOENV   # go 环境变量配置文件存储目录, 默认 $HOME/.config/go/env")
	fmt.Println(tab, "GOPROXY   # go 拉取模块版本时的镜像源, 默认 https://proxy.golang.org,direct  ", "direct 指示 Go 回源到模块版本的源地址去抓取")
	fmt.Println(tab, tab, "可以有多个 comma-separated 分隔的 url, 当查找包返回 404 或者 410 时会自动切换下一个 url 重复查找")
	fmt.Println(tab, "GOSUMDB   # go 拉取模块版本时校验模块数据是否被篡改的源, 默认 sum.golang.org")
	fmt.Println(tab, "GOPRIVATE/GONOPROXY/GONOSUMDB   # 设置当前项目依赖的私有配置, 一般设置 GOPRIVATE")
	fmt.Println(tab, "GOMOD=\"/dev/null\"	# Go Modules Mode 下 go.mod 的绝对路径, 如果没有 go.mod 文件, 则指向 Unix => \"/dev/null\", window => \"NUL\", 如果禁用模块感知, 则为空字符串")
	fmt.Println(tab, "GOTMPDIR	# go 命令将写入临时源文件、包和二进制文件的目录")
	fmt.Println(tab, "GOTOOLDIR	# go 工具(编译, 封面, 文档等)的安装目录, 默认为 $GOROOT/pkg/tool/")
	fmt.Println(tab, "GOVCS	 # 列出可用于匹配服务器的版本控制命令")
	fmt.Println(tab, "GOVERSION	 	# 安装的 go 树的版本, 有 runtime.Version 报告")
	fmt.Println(tab, "GOOS   # 表示目标操作系统, 有 darwin(Mac), linux, windows, android, netbsd, openbsd, solaris, plan9等")
	fmt.Println(tab, "GOARCH   # 表示目标架构, 常见的有 386, arm, amd64等")
	fmt.Println(tab, "GOHOSTARCH	# go 工具链二进制文件的架构(GOARCH)")
	fmt.Println(tab, "GOHOSTOS	# go 工具链二进制文件的操作系统(GOOS)")
	fmt.Println(tab, "AR 	# 使用 gccgo 编译器构建时用于操作库档案的命令, 默认值为 ar")
	fmt.Println(tab, "GOEXE   # 可执行文件名后缀, windows 上为 \".exe\", 其他系统上为 \"\"")
	fmt.Println(tab, "GOFLAGS   # 设置 go 命令 -flag=value 的参数, 可以被覆盖")
	fmt.Println(tab, "GCCGO	# 为 go build -compiler=gccgo 运行的 gccgo 命令")
	fmt.Println(tab, "")
}
