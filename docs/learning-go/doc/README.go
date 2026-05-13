package main

import (
	"fmt"
)

// type fake struct{ io.Writer }

// func fred(logger io.Writer) {
// 	if logger != nil {
// 		logger.Write([]byte("..."))
// 	}
// }

func readme() {
	fmt.Println("Go 语言也称为 Golang, 是由 Google 公司开发的一种静态强类型、编译型、并发型、并具有垃圾回收功能的编程语言, 不是一种纯粹的面向对象编程语言")
	fmt.Println("2006-01-02 15:04:05")
	fmt.Println("------------------")

	// 在Go语言中, 布尔类型的零值（初始值）为 false, 数值类型的零值为 0, 字符串类型的零值为空字符串\"\", 而指针、切片、映射、通道、函数和接口的零值则是 nil
	fmt.Println("Go 语言面向对象去掉了传统 OOP 语言的继承、方法重载、构造函数、析构函数、隐藏的 this 指针等等, 通过 匿名字段 实现继承, 通过 接口 实现多态")
	fmt.Println("---------")
	fmt.Println("继承的价值: 解决代码的复用性和可维护性, 满足 is - a 的关系")
	fmt.Println("接口的价值: 设计, 设计好各种规范(方法), 灵活, 解耦, 满足 like - a 的关系")
	fmt.Println("------------------")

	fmt.Println("命名规范")
	fmt.Println("\033[1;32mmodule\033[0m 模块名一般以斜线(slash)分隔的代码仓库托管服务器域名、用户名、模块根目录名组成, 模块根目录名可以使用 -(Kebab Case) 连字符")
	fmt.Println("\033[1;32mpackage\033[0m 包名一般使用简短有意义的小写形式, 尽量和包所在目录名称保持一致有助于理解, 包名不包含目录路径和 -(Kebab Case) 等特殊字符, (声明包不需要使用路径)")
	fmt.Println("\033[1;32m内部包\033[0m 以 \033[1;36minternal\033[0m 命名的目录或以 internal 命名的目录的子目录中")
	fmt.Println("\033[1;32mimport\033[0m 导入包时, 参数必须以模块名开始, 后面跟随导入包所在的目录至模块根目录的\033[1;32m绝对路径\033[0m(导入包需要使用路径), 如有多层子级目录不能省略中间子级路径")
	fmt.Println("\033[1;32m源文件名\033[0m一般使用简短有意义的小写蛇形(Snake Case)命名方式")
	fmt.Println("\033[1;32m测试函数\033[0m以 \033[1;36mTestXxx\033[0m 或 \033[1;36mBenchmarkXxx\033[0m 或 \033[1;36mFuzzXxx\033[0m 开头")
	fmt.Println("\033[1;32m测试源文件名\033[0m必须以 \033[1;36m_test.go\033[0m 结尾")
	fmt.Println("包外测试辅助源文件名以 \033[1;36mexport_test.go\033[0m 命名")
	fmt.Println("包外测试源文件名以 \033[1;36m被测试包名_test.go\033[0m 命名")
	fmt.Println("静态测试固件存放目录名以 \033[1;36mtestdata\033[0m 命名")
	fmt.Println("---------")
	fmt.Println("驼峰方式命名, 函数、类型、变量等可以被其他包访问的, 使用首字母大写, 只能在包内访问的使用首字母小写")
	fmt.Println("接口名应以 -er 结尾, 表示行为, 例如 Reader, Writer")
	fmt.Println("接口实现者命名可以使用接口名作为后缀, 例如 FileReader")
	fmt.Println("多方法接口用名词描述其功能, 例如 ReadWriteCloser 接口只包含 Read、Write、Close 方法")
	fmt.Println("接收者名称应简短, 通用类型首字母, 也可以使用两个字母简介")
	fmt.Println("-------------------------------")

	fmt.Println("Golang 内联: ")
	fmt.Println(tab, "在Go中, 一个 goroutine 会有一个单独的栈, 栈又会包含多个栈帧, 栈帧是函数调用时在栈上为函数所分配的区域.")
	fmt.Println(tab, "但其实, 函数调用是存在一些固定开销的, 例如维护帧指针寄存器BP、栈溢出检测等.")
	fmt.Println(tab, "因此, 对于一些代码行比较少的函数, 编译器倾向于将它们在编译期展开从而消除函数调用, 这种行为就是内联")
	fmt.Println("所谓内联指编译期间, 直接将调用函数的地方替换为函数的实现, 它可以减少函数调用的开销以提高程序的性能, 内联函数是直接复制镶嵌到主函数中去的, 就是将内联函数的代码直接放在内联函数的位置上")
	fmt.Println("Golang 内联优化范围")
	fmt.Println(tab, "超过 80 个节点的代码量就不再内联. 如 a = a + 1 就是 5 个节点")
	fmt.Println(tab, "每个存在内联优化的函数都会维护一个内联树(inlining tree), 用来找到原始代码在哪")
	fmt.Println("Golang 内联控制")
	fmt.Println(tab, "Go程序编译时, 默认将进行内联优化. 我们可通过 -gcflags=\"-l\" 选项全局禁用内联, 与一个-l禁用内联相反, 如果传递两个或两个以上的 -l 则会打开内联, 并启用更激进的内联策略.")
	fmt.Println(tab, "如果不想全局范围内禁止优化, 则可以在函数定义时添加 //go:noinline 编译指令来阻止编译器内联函数")
	fmt.Println("------------------")

	fmt.Println("线程安全: 在多个线程同时访问共享数据时满足以下条件")
	fmt.Println("    不会出现数据竞争(data race), 多个线程同时对同一数据进行读写操作, 导致数据不一致或未定义的行为")
	fmt.Println("    不会出现死锁(deadlock), 多个线程互相等待对方释放资源而无法继续执行的情况")
	fmt.Println("    不会出现饥饿(starvation), 某个线程因为资源分配不公而无法得到执行的情况")
	fmt.Println("------------------")

	fmt.Println(" 内存分配原则")
	fmt.Println(tab, "当函数外部对指针没有引用时, 优先分配在 栈 上")
	fmt.Println(tab, "当函数外部对指针存在引用时, 优先分配在 堆 上")
	fmt.Println(tab, "当函数内分配一个较大对象时, 优先分配在 堆 上")
	fmt.Println("内存逃逸: 当一个对象的指针被多个方法或线程引用时, 这个指针就发生了逃逸, go build -gcflags='-m -l' ")
	fmt.Println("  go 内存管理分为栈内存管理和堆内存管理, 栈内存管理由编译器管理, 堆上内存由程序管理, 在运行期间申请和释放(垃圾回收), 栈上分配内存比堆上分配内存效率更高, 不需要 GC 处理,")
	fmt.Println(" 内存逃逸的情况:")
	fmt.Println("  在方法内部把局部变量指针返回, 局部变量原本应该在栈中分配, 在栈中回收, 但由于返回时被外部引用, 因此其生命周期大于栈, 则溢出")
	fmt.Println("  发送指针或带有指针的值到 channel 中, 在编译时, 编译器不知道哪个 goroutine 会在 channel 上接收数据")
	fmt.Println("  在一个切片上存储指针或带指针的值, []*string 尽管其后面的数组可能是在栈上分配的, 但其引用的值一定是在堆上")
	fmt.Println("  slice 的背后数组被重新分配了, 因为 append 时可能会超出其容量(cap), 如果切片背后的存储要基于运行时的数据进行扩充, 就会在堆上分配")
	fmt.Println("  在 interface 类型上调用方法, 在 interface 类型上调用方法都是动态调用的--方法的真正实现只能在运行时知道,")
	fmt.Println("------------------")

	fmt.Println("交叉编译: 在一个平台上编译生成另一个平台上运行的可执行文件的过程")
	// CGO_ENABLED=0, 如果开启 go 将使用 C 编译器进行部分编译, 可能会影响跨平台编译的效果, 通常设置为 0
	// GOOS=[] 目标平台的系统名称, windows, linux, darwin
	// GOARCH=[] 目标平台的系统架构, 386, amd64, arm64
	fmt.Println("windows 中交叉编译时需要先修改环境变量再进行交叉编译, PS 中使用 $env:GOOS='linux' 临时修改环境变量")
	fmt.Println("linux 或者 mac 中交叉编译时修改环境变量")
	fmt.Println("------------------")
}
