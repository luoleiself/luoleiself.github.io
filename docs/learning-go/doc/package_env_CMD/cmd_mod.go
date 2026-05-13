package main

import (
	"fmt"
)

func cmdMod() {
	fmt.Println("------------cmdMod()-------------")
	fmt.Println("go mod 不依赖 $GOPATH, 可以脱离 $GOPATH 创建项目, go.mod 和 go.sum 文件是 go modules 版本管理的指导性文件")
	fmt.Println("go.mod 文件指令")
	fmt.Println(tab, "1. module: 定义主模块的路径, go.mod 文件只能包含一个模块指令")
	fmt.Println(tab, "2. go: 标识模块中的代码使用的 go 的版本, 版本必须是有效的 go 发行版本")
	fmt.Println(tab, "3. require: 声明给定模块依赖项的最低必需版本, 对于每个所需的模块版本, go 命令都会自动加载该版本的 go.mod 文件并合并该文件中的依赖,")
	fmt.Println(tab, tab, "go 命令会自动为某些依赖添加 `// indirect` 注释, 表示主模块中的任何包都不直接导入所需模块中的包")
	fmt.Println(tab, "4. exclude: 阻止 go 命令加载模块版本, 1.16 开始, 如果任何 go.mod 文件中的 require 命令引用的模块版本被主模块 go.mod 文件中的 exclude 指令排除, 则忽略该依赖项")
	fmt.Println(tab, "5. replace: 将模块的特定版本或模块的所有版本的内容替换为其他位置的内容, 可以使用其他模块路径和版本或特定于平台的文件路径指定替换")
	fmt.Println(tab, tab, "如果箭头左侧有版本, 则仅更换模块的特定版本, 其他版本将正常访问, 如果省略左侧版本, 则更换模块的所有版本")
	fmt.Println(tab, tab, "如果箭头右侧是绝对或相对路径, 则将其解释为替换模块根目录的本地文件路径, 该路径必须包含 go.mod 文件, 这种情况必须省略替换版本")
	fmt.Println(tab, tab, "如果箭头右侧路径不是本地路径, 则它必须是有效的模块路径, 这种情况需要指定一个版本, 同一模块版本也不能出现在生成列表中")
	fmt.Println(tab, tab, "指令仅适用于主模块的 go.mod 文件, 在其他模块中被忽略")
	fmt.Println(tab, "6. retract: 标识不应依赖 go.mod 定义的模块版本或版本范围, 收回的版本应在版本控制库和模块代理中保持可用, 以确保依赖它们的构建不会被破坏")
	fmt.Println(tab, tab, "当模块被标记回收时, 用户不会使用 go get, go mod tidy 或其他命令自动升级模块版本")
	fmt.Println("------------")
	fmt.Println(`
go.mod 注释只能使用 //, 不能使用 /* */
module github.com/my/repo // 声明当前模块的 module path, 一个模块内的所有包的引用路径都以它为前缀, 一般以 仓库+module name 的格式, 方便 go get 查找
// module github.com/my/repo/v2 // 如果模块的主要版本大于 2 以上时, 需要手动更新模块名称
// module helloWorld // 初始化 helloWorld 目录为 go modules 模式
go 1.16
require (
  github.com/some/dependency v1.2.3 // indirect // indirect 标识间接使用了这个库并没有列到某个 go.mod 中, go 1.17 以上放到单独 require 中
  github.com/another/dependency/v4 v4.0.0 // incompatible // incompatible 标识引入的某些模块不符合规范
)
exclude example.com/banana v1.2.4 // 排除某些引入的模块
replace(
  golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d // 依赖包的查找路径
  golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
retract(
  "1.1.1" // 声明某些版本已废弃需要升级新版本或者降级使用
)`)
	fmt.Println(`
repo
  |-- bar
    |-- bar.go
  |-- foo
    |-- foo.go
  |-- go.mod

import github.com/my/repo/foo // 引入 foo 包

import github.com/my/repo/v2/foo // 引入 v2 以上次要版本号下的包`)

	fmt.Println("go.sum 文件")
	fmt.Println(tab, "由 模块路径、模块版本、哈希校验值 组成, 其中 哈希校验值 用来保证当前缓存的模块不会被篡改, hash 是以 h1: 开头的字符串, 表示生成 checksum 的算法是第一版的 hash 算法(sha256)")
	fmt.Println(tab, "包含 h1:hash 和 go.mod h1:hash, 当 Go 认为用不到某个模块版本的时候就会省略它的 h1:hash, 只存在 go.mod h1:hash 的情况")
	fmt.Println("-------------------------------")

	fmt.Println("go mod 运行 go mod 相关命令必须依赖 go.mod 文件, 否则将报错")
	fmt.Println(tab, "go mod init: 初始化 go modules 项目, 生成 go.mod 文件")
	fmt.Println(tab, "go mod download: 手动触发下载依赖包到本地 cache (默认为 $GOPATH/pkg/mod 目录)")
	fmt.Println(tab, tab, "-x	打印执行命令")
	fmt.Println(tab, tab, "-json	将一系列描述 json 对象打印到标准输出")
	fmt.Println(tab, "go mod tidy: 添加缺少的包, 且删除无用的包, 维护 go.mod 和 go.sum 文件")
	fmt.Println(tab, "go mod graph: 打印项目的模块依赖结构")
	fmt.Println(tab, "go mod verify: 校验模块是否被篡改过")
	fmt.Println(tab, "go mod why: 查看为什么需要依赖")
	fmt.Println(tab, "go mod vendor: 导出项目所有依赖到 vendor 下")
	fmt.Println(tab, "go mod edit: 编辑 go.mod 文件, 接 -fmt 参数格式化 go.mod 文件, 接 -require=golang.org/x/text 添加依赖, 接 -droprequire=golang.org/x/text 删除依赖")
	fmt.Println(tab, tab, "-require=path@version 或 -droprequire=path 添加或移除指定版本模块依赖, 该参数主要为了 工具 去理解模块依赖, 普通用户应该更喜欢 go get path@version 和 go get path@none 添加和移除模块")
	fmt.Println(tab, tab, "-exclude=path@version 或 -dropexclude=path@version 添加或移除指定版本模块的排除")
	fmt.Println(tab, tab, "-replace=old[@v]=new[@v] 添加指定版本模块的替换")
	fmt.Println(tab, tab, "-dropreplace=old[@v] 移除指定版本模块的替换")
	fmt.Println(tab, tab, "-retract=version 或 -dropretract=version 添加或删除指定版本的撤回")
	fmt.Println(tab, tab, "-fmt 格式化输出 go.mod 文件")
	fmt.Println(tab, tab, "-print 文本格式化输出 go.mod 文件")
	fmt.Println(tab, tab, "-json JSON 格式化输出 go.mod 文件")
	fmt.Println(tab, tab, "-module 改变模块路径")
	fmt.Println(tab, tab, "-go=version 设置 go 版本")
}
