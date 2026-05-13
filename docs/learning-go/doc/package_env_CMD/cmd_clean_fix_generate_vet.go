package main

import (
	"fmt"
)

func cmdCFGV() {
	fmt.Println("------------------cmdCFGV()----------------")
	fmt.Println("go clean 清除对象文件和缓存文件")
	fmt.Println(tab, "-i	清除相应的已安装存档或二进制文件")
	fmt.Println(tab, "-n	打印它将执行的清除命令, 但不会运行这些命令")
	fmt.Println(tab, "-r	递归删除所有通导入路径命令的包的依赖项")
	fmt.Println(tab, "-x	执行删除命令时打印这些命令")
	fmt.Println(tab, "-cache	清除整个构建的缓存")
	fmt.Println(tab, "-testcache	清除构建的缓存中所有过期的结果")
	fmt.Println(tab, "-modcache		清除整个模块下载缓存, 包含那些版本依赖项的未打包的源码")
	fmt.Println(tab, "-fuzzcache	清除所有的模糊测试缓存")
	fmt.Println("-----------------------------------------")

	fmt.Println("go fix	把指定代码包的所有 Go 语言源码文件中的旧版本代码修正为新版本的代码")
	fmt.Println(tab, "go fix [package]")
	fmt.Println(tab, "-diff 不将修正后的内容写入文件, 而只打印修正前后的内容的对比信息到标准输出")
	fmt.Println(tab, "-r 只对目标源文件做有限的修正操作, 该标记的值即为允许的修正操作的名称, 多个名称之间用英文半角逗号分隔")
	fmt.Println(tab, "-force 即使源文件中的代码已经与 Go 语言的最新版本相匹配了, 也会强行执行指定的修正操作")
	fmt.Println("-----------------------------------------")

	fmt.Println("go generate 生成由现有文件中的指令描述的运行命令. 这些命令可以运行任何进程, 但目的是创建或更新Go源文件, go generate 永远不会通过 build, get, test等命令自动运行, 需要手动运行")
	fmt.Println(tab, "")
	fmt.Println("-----------------------------------------")

	fmt.Println("go vet 用于检查 Go 语言源码中静态错误的简单工具")
	fmt.Println(tab, "-all 进行全部检查")
	fmt.Println(tab, "-assign 检查赋值语句")
	fmt.Println(tab, "能够捕获的错误:")
	fmt.Println(tab, tab, "Printf 类函数调用时, 类型匹配错误的参数")
	fmt.Println(tab, tab, "定义常用的方法时, 方法签名的错误")
	fmt.Println(tab, tab, "错误的结构标签")
	fmt.Println(tab, tab, "没有指定字段名的结构字面量")
}
