package main

import (
	"fmt"
)

func cmdGIL() {
	fmt.Println("----------------cmdGIL()--------------")
	fmt.Println("go get 添加依赖到当前模块中(go.mod文件)")
	fmt.Println(tab, "1.16 以上版本推荐使用 go install 编译安装包, 如果在不是主模块没有 go.mod 文件下运行时, replace 和 exclude 命令不会应用")
	fmt.Println(tab, "-d 该参数标识只执行下载包, 而不执行构建和安装, 1.18 开始总是开启")
	fmt.Println(tab, "-fix 该参数标识在解析依赖项或生成代码之前对下载的包运行修复工具")
	fmt.Println(tab, "-u	更新现有的依赖, 会强制更新它所依赖的其它全部模块到最新的次要版本或者修订版本, 不包括自身")
	fmt.Println(tab, tab, "-u=patch 更新现有的依赖到最新的修订版本")
	fmt.Println(tab, "-t 	指示 get 考虑构建命令行中指定的包测试所需的模块")
	fmt.Println(tab, "-u -t	更新所有直接依赖和间接依赖的模块版本, 包括单元测试中用到的")
	fmt.Println(tab, "go get golang.org/x/text@latest	拉取最新的版本, 若存在 tag, 则优先使用")
	fmt.Println(tab, "go get golang.org/x/text@master	拉取 master 分支的最新 commit")
	fmt.Println(tab, "go get golang.org/x/text@v0.3.2	拉取 tag 为 v0.3.2 的 commit")
	fmt.Println(tab, "go get golang.org/x/text@342b2e	拉取 hash 为 342b231 的 commit, 最终会被转换为 v0.3.2")
	fmt.Println(tab, "go get golang.org/x/text@none		使用 @none 移除指定依赖模块及其依赖模块")
	fmt.Println(tab, "")
	fmt.Println("---------------------------------------")

	fmt.Println("go install 编译并安装由导入路径命名的包")
	fmt.Println(tab, "可执行文件安装在由 GOBIN 环境变量命名的目录中, 如果未设置 GOBIN 环境变量, 则默认为 $GOPATH/bin 或 $HOME/go/bin")
	fmt.Println(tab, "$GOROOT 中的可执行文件安装在 $GOROOT/bin 或 $GOTOOLDIR 中, 而不是 GOBIN 中")
	fmt.Println(tab, "go install golang.org/x/tools/gopls@latest")
	fmt.Println(tab, "go install golang.org/x/tools/gopls@v0.6.4")
	fmt.Println("---------------------------------------")

	fmt.Println("go list 每行一个列出命名的包")
	fmt.Println(tab, "-json	使用 JSON 格式非 模板 格式输出列表信息")
	fmt.Println(tab, "-f	使用包模板的语法指定列表的替代格式")
	fmt.Println(tab, "-m	标志导致 list 列出模块⽽不是包")
	fmt.Println(tab, "-e	以容错模式加载和分析指定的代码包")
	fmt.Println(tab, "go list -m -json all: 以 json 的方式打印依赖详情")
}
