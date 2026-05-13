package main

import (
	"fmt"
	"unsafe"
)

func readme() {
	fmt.Println("标识符: 是一种字符序列, Go语言对各种类型、变量、常量、方法、函数等命名时使用的字符序列, 标识符由若干个字母、下划线_、和数字组成, 且第一个字符必须是字母")
	fmt.Println("25个关键字: \n\tbreak\t default\t func\t interface\t select\t case\t defer\t go\t map\t struct\t chan\t else\t goto\t package\t switch\t const\t fallthrough\t if\t range\t type\t continue\t for\t import\t return\t var\t")
	fmt.Println("36个预定义标识符: ")
	fmt.Println("内建常量: true\t false\t iota\t nil\t")
	fmt.Println("内建类型: int\t int8\t int16\t int32\t int64\t uint\t uint8\t uint16\t uint32\t uint64\t uintptr\t byte\t rune\t float32\t float64\t complex64\t complex128\t bool\t string\t error\t")
	fmt.Println(tab, "有符号整数最高 bit 位表示符号位, n-bit 的取值范围从 -2^(n-1) 到 2^(n-1)-1. int8 类型的取值范围从 -128 到 127")
	fmt.Println(tab, "无符号整数的所有 bit 位用于表示非负数, 取值范围从 0 到 2^n-1. uint8 类型的取值范围从 0 到 255")
	fmt.Println("内置函数: make\t len\t cap\t new\t append\t copy\t close\t delete\t complex\t real\t imag\t panic\t recover\t print\t")
	fmt.Println("---------------")
	fmt.Println("数据类型: 基本数据类型 - 数值型(- 整数 - 浮点数 - 复数 -) - 字符型 - 字符串 - 布尔型")
	fmt.Println("数据类型: 复合数据类型 - 聚合类型(- 数组 - 结构体 -) - 引用类型(- 指针 - slice - map - func - channel -) - interface 类型 -")
	fmt.Println("---------------")

	fmt.Println("map slice func 不可比较")
	fmt.Println("array 可比较, 编译器知道两个数组是否一致")
	fmt.Println("struct 可比较, 可以逐个比较结构体的值")
	fmt.Println("channel 可比较, 必须是由同一个 make 生成, 同一个 channel 结果为 true, 否则为 false")
	fmt.Println("---------------")

	fmt.Println("普通指针类型: 只能用于传递地址(持有对象), 不能进行指针运算, 使用 &(取地址) *(根据地址取值)")
	fmt.Println("uintptr 可以保存任意指针的位模式的整数类型")
	fmt.Println(tab, "是一个能足够容纳指针位数大小的无符号整数类型, 可以进行指针运算, 无法持有对象, GC 不把 uintptr 当指针, uintptr 类型的目标会被 GC 回收")
	fmt.Println("unsafe.Pointer 可以指向任意类型的指针, 不能进行指针运算, 不能读取内存存储的值(想读取的话需要转成相应类型的指针)")
	fmt.Println(tab, "unsafe.Pointer 是一个桥梁, 让任意类型的指针实现相互转换, 也可以转换成 uintptr 进行指针运算")
	fmt.Println("\033[1;32m指针 <=> unsafe.Pointer <=> uintptr\033[0m")
	fmt.Println("  1. 任意类型的指针值都可以转换为 unsafe.Pointer")
	fmt.Println("  2. unsafe.Pointer 可以转换为任意类型的指针值")
	fmt.Println("  3. uintptr 可以转换为 unsafe.Pointer")
	fmt.Println("  4. unsafe.Pointer 可以转换为 uintptr")
	u1 := int64(100)
	var ptr *int = (*int)(unsafe.Pointer(&u1))
	fmt.Printf("ptr %d\n", *ptr) // 100 // int64 转换为 int
	fmt.Println("---------------")

	fmt.Println("make 和 new 都可以用于初始化分配内存的内建函数, 且在堆上分配内存, make 分配内存也初始化内存, new 只是将内存清零,并没有初始化内存")
	fmt.Println(" make 返回和参数类型一致的类型, 只能用来初始化 slice, map, channel")
	fmt.Println(" new 返回参数类型的指针类型, 指针指向的内存区域值默认为 nil, 可以初始化为任意类型")
	var a = new(int)
	*a = 10
	fmt.Println("var a = new(int)\n*a = 10")
	fmt.Println(*a)
	fmt.Println("------------------")

	fmt.Println("var 声明语句不能用在 if, switch, for 结构的关键字初始化处, 只能使用简洁声明方式声明")
	fmt.Println(" if var i = 1 { } // var declaration not allowed in if initializer")
	fmt.Println(" switch var i = 1 { } // var declaration not allowed in switch initializer")
	fmt.Println(" for var i = 1;i < 10;i++ { } // var declaration not allowed in for initializer")
	fmt.Println("---------------")

	fmt.Println("nil 预定义标识符不能比较, 不是关键字或保留字")
	fmt.Println(tab, "没有默认类型, nil 是唯一的一个 go 语言中没有默认类型的非类型值. 对于编译器来说, 必须从上下文中获取充足的信息才能推断出 nil 的类型")
	fmt.Println(tab, "不同类型的 nil 值占用的内存大小可能是不一样的")
	fmt.Println(tab, "不同类型 nil 的指针是一样的")
	fmt.Println(tab, "不同类型的 nil 是不能比较的")
	fmt.Println(tab, "两个相同类型的 nil 值也无法比较, 但是可以将不可以比较类型的空值直接与 nil 进行比较")
	fmt.Println("---------------")
}
