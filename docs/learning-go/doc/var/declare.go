package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 全局变量 方式1
var n1 int = 100
var n2 int = 200
var name string = "Jerry"

// 全局变量 方式2
var (
	n3    int = 300
	n4        = 400
	n5    float32
	title = "Tom"
)

// 全局变量 方式3: 由编译器自动类型推断, 全局变量和局部变量都可用, 不可手动添加类型
var n9, n10, newName = 900, 1000, "Animal"

type Gender uint8

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func (g Gender) String() string {
	if g == 0 {
		return "FEMALE"
	} else if g == 1 {
		return "MALE"
	}
	return "hello world"
}

const (
	B int = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

type Month uint8
type Week uint8

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
const (
	Sunday Week = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func declareNote() {
	fmt.Println(FEMALE, MALE, THIRD)
	fmt.Println("-----------declareNote()----------")
	fmt.Println(1<<1, 1<<2, 1<<3, 1<<4) // 1x2**1=2 	1x2**2=4 	1x2**3=8 	1x2**4=16
	fmt.Println(2<<1, 2<<2, 2<<3, 2<<4) // 2x2**1=4 	2x2**2=8 	2x2**3=16 	2x2**4=32
	fmt.Println(3<<1, 3<<2, 3<<3, 3<<4) // 3x2**1=6 	3x2**2=12	3x2**3=24	3x2**4=48
	fmt.Println(1>>1, 1>>2, 1>>3, 1>>4) // 1/2**1=0 	1/2**2=0 	1/2**3=0 	1/2**4=0
	fmt.Println(2>>1, 2>>2, 2>>3, 2>>4) // 2/2**1=1 	2/2**2=0 	2/2**3=0 	2/2**4=0
	fmt.Println(3>>1, 3>>2, 3>>3, 3>>4) // 3/2**1=1 	3/2**2=0 	3/2**3=0 	3/2**4=0
	fmt.Println("-------------------------------")

	var n = 1.1
	var a = 10
	var name string = "Jerry"
	var name1 string = ""
	var b byte = 'b'
	var b1 = 'e'
	var c = '中'
	var d rune = '国'
	fmt.Println("声明变量时,由编译器类型推导后的变量类型(取所属类型的最大字节数存储):  ")
	fmt.Print(tab, "字符类型 byte[1] rune[4] 取 rune[4] 相当于 int32\n")
	fmt.Print(tab, "浮点数 float[32] float[64] 取 float[64] \n")
	fmt.Print(tab, "整类型取 int[4|8] 根据系统类型\n")
	fmt.Printf("var n = 1.1 声明的 n 的类型为 %T 值为 %f 占用字节数 %d\n", n, n, unsafe.Sizeof(n))                              // 占用字节 8
	fmt.Printf("var a = 10 声明的 a 的类型为 %T 值为 %d  占用字节数 %d\n", a, a, unsafe.Sizeof(a))                              // 占用字节 8
	fmt.Printf("var name string = \"Jerry\" 声明的 name 的类型为 %T 值为 %s 占用字节数  %d\n", name, name, unsafe.Sizeof(name)) // 占用字节 16
	fmt.Printf("var name1 string = \"\" 声明的 name1 的类型为 %T 值为 %s 占用字节数  %d\n", name1, name1, unsafe.Sizeof(name1)) // 占用字节 16
	fmt.Printf("var b byte = 'b' 声明的 b 的类型为 %T 值为 %c 占用字节数 %d\n", b, b, unsafe.Sizeof(b))                         // 占用字节 1
	fmt.Printf("var b1 = 'e' 声明的 b1 的类型为 %T 值为 %c 占用字节数 %d\n", b1, b1, unsafe.Sizeof(b1))                         // 占用字节 4
	fmt.Printf("var c = '中' 声明的 c 的类型为 %T 值为 %c 占用字节数 %d\n", c, c, unsafe.Sizeof(c))                              // 占用字节 4
	fmt.Printf("var d rune = '国' 声明的 d 的类型为 %T 值为 %c 占用字节数 %d\n", d, d, unsafe.Sizeof(d))                         // 占用字节 4
	fmt.Println("-------------------------------")

	var i32 int32 = 32
	var i int = 10
	var i8 int8 = 8
	var i16 int16 = 16
	var i64 int64 = 64
	var s1 string = "he"
	var s2 string = ""
	var s3 string = "你好世界"
	var ui uint = 10
	var ui8 uint8 = 8
	var ui16 uint16 = 16
	var ui32 uint32 = 32
	var ui64 uint64 = 64

	fmt.Printf("i 的类型为 %T 内存地址为 %p 值为 %v \n", i, &i, i) // int 0xc000094030 10
	// int8 类型变量的内存地址为上一个整数类型变量的内存地址加上当前类型的字节占用大小
	fmt.Printf("i8 的类型为 %T 内存地址为 %p 值为 %v\n", i8, &i8, i8) // int8 0xc000094038 8

	// 字符串类型变量的内存地址为上一个声明字符串变量的内存地址加上 16 个字节
	fmt.Printf("s1 的类型为 %T 内存地址为 %p 值为 %v\n", s1, &s1, s1) // string 0xc00009c050 "hello world"
	fmt.Printf("s2 的类型为 %T 内存地址为 %p 值为 %v\n", s2, &s2, s2) // string 0xc00009c060
	fmt.Printf("s3 的类型为 %T 内存地址为 %p 值为 %v\n", s3, &s3, s3) // string 0xc00009c070 "你好世界"

	// int16 类型变量的内存地址为上一个整数类型变量的内存地址加上当前类型的字节占用大小
	fmt.Printf("i16 的类型为 %T 内存地址为 %p 值为 %v\n", i16, &i16, i16) // int16 0xc00009403a 16

	// 无符号整数类型变量的内存地址为上一个整数型变量的内存地址加上当前类型的字节占用大小
	fmt.Printf("ui 的类型为 %T 内存地址为 %p 值为 %v\n", ui, &ui, ui)         // uint 0xc000094048 10
	fmt.Printf("ui8 的类型为 %T 内存地址为 %p 值为 %v\n", ui8, &ui8, ui8)     // uint8 0xc000094050 8
	fmt.Printf("ui16 的类型为 %T 内存地址为 %p 值为 %v\n", ui16, &ui16, ui16) // uint16 0xc000094052 16
	fmt.Printf("ui32 的类型为 %T 内存地址为 %p 值为 %v\n", ui32, &ui32, ui32) // uint32 0xc000094054 32
	fmt.Printf("ui64 的类型为 %T 内存地址为 %p 值为 %v\n", ui64, &ui64, ui64) // uint64 0xc000094058 64

	// int32 类型变量的内存地址为上一个整数类型变量的内存地址加上当前类型的字节占用大小
	fmt.Printf("i32 的类型为 %T 内存地址为 %p 值为 %v\n", i32, &i32, i32) // int32 0xc000094028 32

	// int64 类型变量的内存地址为上一个整数类型变量的内存地址加上当前类型的字节占用大小
	fmt.Printf("i64 的类型为 %T 内存地址为 %p 值为 %v\n", i64, &i64, i64) // int64 0xc000094040 64
	fmt.Println("-------------------------------")

	variable()
	fmt.Println("-------------------------------")

	globalVariable()
	fmt.Println("-------------------------------")

	localVariable()
	fmt.Println("-------------------------------")

	constNote()
	fmt.Println("-------------------------------")

	pointerNote()
	fmt.Println("-------------------------------")
}
func variable() {
	fmt.Println("声明: 只声明未初始化的变量会被赋值为类型零值, 数值型和字符型为 0, 字符串为空字符串, 布尔为false, 数组的零值为元素类型的零值, slice map pointer interface chan func的零值为nil")
	fmt.Println("var 变量名 数据类型")
	fmt.Println("---------------")
	fmt.Println("零值是变量没有做初始化时系统默认设置的值")
	var lv bool
	fmt.Printf("  布尔类型变量零值 %t\n", lv) // false
	var lv1 string
	fmt.Printf("  字符串类型变量零值 %q\n", lv1) // ""
	var lv2 byte
	fmt.Printf("  字符类型变量零值 %v\n", lv2) // 0
	var lv3 int
	fmt.Printf("  整数类型变量零值 %v\n", lv3) // 0
	var lv4 float32
	fmt.Printf("  浮点类型变量零值 %v\n", lv4) // 0
	var lv5 [1]int
	fmt.Printf("  数组类型变量零值 %v\n", lv5) // [0]
	var lv6 []int
	fmt.Printf("  slice类型变量零值 %#v\n", lv6) // []int(nil)
	var lv7 map[int]string
	fmt.Printf("  map类型变量零值 %#v\n", lv7) // map[int]string(nil)
	var lv8 chan int
	fmt.Printf("  chan类型变量的零值 %#v\n", lv8) // (chan int)(nil)
	var lv9 interface{}
	fmt.Printf("  interface类型变量的零值 %#v\n", lv9) // <nil>
	fmt.Println("---------------")
}

func globalVariable() {
	fmt.Println("变量声明方式1: 变量声明的标准格式, 全局变量和局部变量都可用")
	fmt.Print("var n1 int = 100\n")
	fmt.Print("var n2 int = 200\n")
	fmt.Print("var name string = Jerry\n")
	fmt.Printf("输出声明的变量: n1 = %d, n2 = %d, name = %s\n", n1, n2, name)
	fmt.Println("----------------")
	fmt.Println("变量声明方式2(批量声明): 批量声明变量, 可添加类型, 也可以由编译器自动类型推导, 全局变量和局部变量都可用")
	fmt.Print(`
var (
  n3    int = 300
  n4        = 400
  title     = "Tom"
)`)
	fmt.Printf("\n输出声明的变量: n3 = %d, n4 = %d, title = %s, n5 = %.2f\n", n3, n4, title, n5)
	fmt.Println("----------------")
	fmt.Println("变量声明方式3(行内批量声明): 由编译器自动类型推断, 全局变量和局部变量都可用, \033[1;32m不可手动添加不同类型\033[0m")
	// var n9, n10 int, newName string = 900, 100, "Animal" // syntax error: unexpected comma at end of statement
	fmt.Println("// var n9, n10 int, newName string = 900, 100, \"Animal\" // syntax error: unexpected comma at end of statement")
	fmt.Print("var n9, n10, newName = 900, 1000, \"Animal\"\n")
	fmt.Printf("输出声明的变量: var n9, n10, newName = %d, %d, %s\n", n9, n10, newName)
	var n11, n12 int = 12, 13 // OK
	fmt.Printf("输出声明的变量: var n11, n12 int = %d, %d\n", n11, n12)
	fmt.Println("----------------")
}

func localVariable() {
	fmt.Println("变量声明方式4: 简洁声明方式 := 声明方式只能在函数和方法内使用")
	fmt.Println(tab, "局部变量的简洁声明方式1: \033[1;32m不能添加类型和 var 关键字\033[0m")
	ln5 := 500
	ln6 := 600
	localName := "Dog"
	fmt.Println(tab, "ln5 := 500")
	fmt.Println(tab, "ln6 := 600")
	fmt.Println(tab, "localName := \"Dog\"")
	fmt.Print(tab)
	fmt.Printf("输出声明的局部变量: ln5 := %d, ln6 := %d, localName := %s\n", ln5, ln6, localName)
	fmt.Println("----------------")
	fmt.Println(tab, "局部变量的简洁声明方式2: 使用简洁声明方式批量声明, 不能添加类型, 否则编译报错")
	ln7, ln8, localTitle := 700, 800, "Cat"
	fmt.Println(tab, "ln7, ln8, localTitle := 700, 800, \"Cat\"")
	fmt.Print(tab)
	fmt.Printf("输出声明的局部变量: ln7, ln8, localTitle := %d, %d, %s\n", ln7, ln8, localTitle)
	fmt.Println("----------------")
}

func constNote() {
	fmt.Println("--------------constNote()--------------")
	fmt.Println("定义常量的表达式必须为能被编译器求值的常量表达式")
	fmt.Println(tab, "const 声明的常量的值只能为\033[1;32m基本数据类型\033[0m")
	fmt.Println(tab, "常量组中如果不指定类型和初始值, 则与上一行非空常量的值相同")
	fmt.Println(tab, "批量常量声明: 除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式，对应的常量类型也是一样的")
	fmt.Println(tab, "常量间的所有算术运算、逻辑运算和比较运算的结果也是常量, 对常量的类型转换操作或以下函数调用都是返回常量结果: len、cap、real、imag、complex 和 unsafe.Sizeof")
	fmt.Println("------------------------------")

	fmt.Println("常量无法寻址, 不能进行取指针操作")
	// const j = 100
	// fmt.Println(&j, j) // invalid operation: cannot take address of j (untyped int constant 100)
	fmt.Println(`
const j = 100
fmt.Println(&j, j) // invalid operation: cannot take address of j (untyped int constant 100)/`)
	fmt.Println("-------------------------------")

	const c1 = 3 / 2
	fmt.Println(`const c1 = 3 / 2`)
	fmt.Printf("c1 的类型为 %T 值为 %v\n", c1, c1)
	fmt.Println("-------------")
	const (
		a string = "1"
		b
		c = 3
		d
	)
	fmt.Printf("a 的类型为 %T 的值为 %v \nb 的类型为 %T 的值为 %v \nc 的类型为 %T 的值为 %v \nd 的类型为 %T 的值为 %v\n", a, a, b, b, c, c, d, d) // 1 1 3 3
	fmt.Println("------------------------------")

	fmt.Println("iota 常量累加器, 只能用在 const 批量声明语句中 否则报错: cannot use iota outside constant declaration")
	fmt.Println(tab, "每当 const 关键字出现时, 自动重置 iota 为 0")
	fmt.Println(tab, "const 声明语句中每新增一行常量声明 iota 计数一次(可以理解为 const 语句中的行索引)")
	const a0 = iota // const 声明第一行, iota 重置为 0
	const (
		a1         = iota             // const 声明第一行, iota 重置为 0
		a2                            // const 声明新增第 2 行， iota 为 1
		a3         = 6                // 自定义常量
		a4                            // 不赋值的操作和最近上一行非空常量的值相同
		a5         = iota             // const 声明新增第 5 行, iota 为 4
		a6, a7, a8 = iota, iota, iota // iota 计数以行为单位, 单行批量声明计数仍为单行计数
		a9         = 'h'              // 自定义常量
		a10                           // 不赋值的操作和最近上一行非空常量的值相同
		a11        = iota             // const 声明新增第 9 行, iota 为 8
		a12                           // const 声明新增第 10 行, iota 为 9
	)
	// a0 = 0, a1 = 0, a2 = 1, a3 = 6, a4 = 6, a5 = 4, a6 = 5, a7 = 5, a8 = 5, a9 = 104, a10 = 104, a11 = 8, a12 = 9
	fmt.Printf("a0 = %v, a1 = %v, a2 = %v, a3 = %v, a4 = %v, a5 = %v, a6 = %v, a7 = %v, a8 = %v, a9 = %q, a10 = %q, a11 = %v, a12 = %v\n", a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	fmt.Println("------------------------------")

	fmt.Printf("Friday 的类型为 %T 值为 %v\n", Friday, Friday)
	fmt.Println("------------------------------")

	fmt.Println("无类型常量: 编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有 256bit 的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串")
	fmt.Println("通过延迟明确常量的具体类型，不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换")
	fmt.Println("math.Pi 无类型的浮点数常量, 可以直接用于任意需要浮点数或复数的地方, math.Pi 被确定为特定类型，则需要有一个明确的强制类型转换")
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Printf("x 的类型为 %T 值为 %v \ny 的类型为 %T 值为 %v \nz 的类型为 %T 值为 %v\n", x, x, y, y, z, z)
}

func swap(a, b *int) {
	fmt.Printf("交换前 a 的地址为 %p 值为 %v 指针变量的值为 %v b 的地址为 %p 值为 %v 指针变量的值为 %v \n", &a, a, *a, &b, b, *b)
	b, a = a, b
	fmt.Printf("交换后 a 的地址为 %p 值为 %v 指针变量的值为 %v b 的地址为 %p 值为 %v 指针变量的值为 %v \n", &a, a, *a, &b, b, *b)
}

func pointerNote() {
	fmt.Println("------------pointerNote()-------------")
	x, y := 1, 2
	fmt.Printf("x 的地址为 %v 值为 %v y 的地址为 %p 值为 %v \n", &x, x, &y, y)
	swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println("----------------------------")

	var num *int
	fmt.Printf("num 的类型为 %T 地址为 %p 值为 %v\n", &num, num, num)
	fmt.Println("----------------------------")
}
