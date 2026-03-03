package main

import (
	"fmt"
	"time"
)

func pkgDoc() {
	fmt.Println("package: 包是多个 Go 源码的集合, 是一种高级的代码复用方案, 任何源文件都必须属于某个包")
	fmt.Println(" 源文件的第一行有效代码必须是除空行和注释行之外 package packageName 语句, 声明自己所在的包")
	fmt.Println(tab, "包名的定义不包括目录路径, 包名一般使用小写的简短且有意义的名称, 包名一般要和所在的目录同名, 也可以不同, 包名中不能包含 - 等特殊字符")
	fmt.Println(tab, "包一般使用域名作为目录名称, 这样能保证包名的唯一性, 例如 Github 项目的包一般会放到 GOPATH/src(pkg/mod 1.11)/github.com/userName/projectName 目录下")
	fmt.Println(tab, "包名为 main 的包为应用程序的入口, 编译后生成二进制可执行文件, 编译不包含 main 包的源文件时不会得到任何可执行文件")
	fmt.Println(tab, "同一个目录下的所有源文件只能属于同一个包, 同样属于同一个包的源文件不能放在多个目录下")
	fmt.Println(tab, "")
	fmt.Println(" 内部包(internal package): 一个名为 internal 的目录或一个名为 internal 的目录的子目录中")
	fmt.Println("-------------------------------")

	fmt.Println("同一个包中只能有一个 main 函数, 但是可以有多个 init 函数, 见 /net 目录下 main 包 ")
	fmt.Println("  入口唯一性, main 包是程序的唯一入口, 如果将 main 包作为其他普通包导入, 会导致 循环依赖风险, 多 main 函数冲突.")
	fmt.Println("  编译隔离性, main 包生成可执行文件, 其他普通包会编译成静态库文件供其他包复用")
	fmt.Println("  代码组织规范, main 包应仅包含程序启动和协调逻辑, 而不应该承载业务逻辑可复用代码")
	fmt.Println("  避免隐士副作用, 如果允许导入 main 包, main 包中的 init 函数将会自执行, 导致不可控的副作用")
	fmt.Println(tab, "init 函数先于 main 函数执行, 作为程序的初始化")
	fmt.Println(tab, "init 函数不能被调用或引用, 没有返回值, 只能在程序运行时自动执行")
	fmt.Println(tab, "init 函数只会执行一次, 与函数所在包被其他包引用的次数无关")
	fmt.Println("init 函数通常用于")
	fmt.Println("  变量初始化")
	fmt.Println("  检查 / 修复状态")
	fmt.Println("  注册器")
	fmt.Println("  运行计算")
	fmt.Println("-------------------------------")

	fmt.Println("import 语句通常放在源码文件开头 package 声明语句的下面")
	fmt.Println(tab, "导入包时需要使用包所在目录树的全路径, 导入的包名需要使用双引号包起来")
	fmt.Println(tab, "标准导入格式: import \"fmt\" // 导入 fmt 包, 使用时 fmt.Println()")
	fmt.Println(tab, "自定义别名格式: import F \"fmt\" // 使用时 F.Println()")
	fmt.Println(tab, "省略引用方式: import . \"fmt\" // 相当于把 fmt 包直接合并到当前程序中, 在使用 fmt 包内的方法时可以不用加包名调用 Println()")
	fmt.Println(tab, "匿名引用方式: import _ \"fmt\" // 匿名导入的包和其他方式导入的包一样都会被编译到可执行文件中, 此种方式如果导入包中包含 init 函数, 则只会执行 init 函数")
	fmt.Println(tab, "------------------")
	fmt.Println(tab, "Each package within a module is a collection of source files in the same directory that are compiled together")
	fmt.Println(tab, "canonical pseudo-version semantic hyphen smooth explicit")
	fmt.Println(tab, "1. 在使用第三方包的时候，当源码和.a均已安装的情况下, 编译器链接的是源码")
	fmt.Println(tab, "2. 使用第三方包源码，实际上是链接了以该最新源码编译的临时目录下的 .a 文件而已")
	fmt.Println(tab, "3. import后面的最后一个元素应该是路径, 就是目录, 并非包名")
	fmt.Println("-------------------------------")
	time.Sleep(8 * time.Second)

	fmt.Println("版本: 标识模块的不可变快照, 可以是发布版或预发布版, 每个版本都以字母 v 开头, 后跟由三个点分隔的非负整数组成的语义版本(主版本.次要版本.补丁版本), 例如 v1.2.3")
	fmt.Println("伪版本: 一种特殊格式的预发布版本, 指对没有语义版本标签可用的修订, 可用于在创建版本标签之前测试提交, 例如: v0.0.0-20191109021931-daa7c04131f5")
	fmt.Println(tab, "基本版本前缀: vX.0.0 或 vX.Y.Z-0, 通常从修订版之前的语义版本标签派生, 比基本版本高, 比下一个标记版本低")
	fmt.Println(tab, "时间戳: 创建修订的 UTC 时间, 在 Git 中指提交时间, 非创作时间")
	fmt.Println(tab, "修订标识符: 提交哈希的 12 个字符前缀")
	fmt.Println("使用 go get 模块名@commitID 的方式如果未找到标签 go 命令自动生成伪版本号并更新到 go.mod 文件")
	fmt.Println("最小版本选择(MVS): MVS 生成用于构建的模块版本列表, 从主模块开始遍历该图, 跟踪每个模块所需要的最高版本.")
	fmt.Println("在遍历结束时, 所需的最高版本构成构建列表: 它们是满足所有要求的最低版本")
	fmt.Println(tab, "")
	fmt.Println("-------------------------------")
}
