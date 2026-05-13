package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	id     int64
	name   string
	age    int16
	height float32
}

// - 表示空白字节, _ 表示占用字节, 数字表示字段偏移量
type A struct {
	// 0 - - - 4 _ _ _ 8 _ _ _ _ _ _ _ // 结构体占用 16 字节, 是对齐系数 8 的整数倍, Sizeof 16
	a int8  // 1 1
	b int32 // 4 4
	c int64 // 8 8
}
type B struct {
	// 0 _ _ _ - - - - 8 _ _ _ _ _ _ _ 16 - - - - - - - // 结构体占用 17 字节, 不是对齐系数 8 的整数倍, 外部长度需要在末尾填充 7 个空白字节, Sizeof 24
	b int32 // 4 4
	c int64 // 8 8
	a int8  // 1 1
}
type C struct {
	// 0 _ _ _ _ _ _ _ 8 _ _ _ 12 - - - // 结构体占用 13 字节, 不是对齐系数 8 的整数倍, 外部长度需要在末尾填充 3 个空白字节, Sizeof 16
	c int64 // 8 8
	b int32 // 4 4
	a int8  // 1 1
}
type D struct {
	// 0 _ _ _ _ _ _ _ 8 - - - 12 _ _ _ // 结构体占用 16 字节, 是对齐系数 8 的整数倍, Sizeof 16
	c int64 // 8 8
	a int8  // 1 1
	b int32 // 4 4
}

func memAlignment() {
	fmt.Println("-----------内存对齐-----------")
	fmt.Println(tab, "是一种提高内存访问效率的技术, 指将结构体中的字段按照一定规则分配到内存中的对齐方式,")
	fmt.Println(tab, " 它的原理是操作系统访问内存时是按照字长为单位的, 字长是 CPU 一次能读取的内存数据的大小,")
	fmt.Println(tab, " 比如在 64 位机器上, 字长为 8 字节, 如果内存数据的地址是字长的整数倍, 那么 CPU 就可以一次读取到完整的数据,")
	fmt.Println(tab, " 否则就需要多次访问内存, 造成效率降低. 内存对齐还可以保证内存数据的原子性, 比如在 32 位平台上进行 64 位的原子操作,")
	fmt.Println(tab, " 就必须要求数据是 8 字节对齐的, 否则可能会出现 panic")
	// fmt.Println("为了方便内存对齐, go 位变量提供了对齐系数(alignof), 表示变量的地址必须是对齐系数的整数倍,")
	fmt.Println(" 通过内存对齐后, 就可以保证在访问一个变量地址时")
	fmt.Println("    如果该变量占用内存大小小于字长时, 保证一次访问就能得到数据")
	fmt.Println("    如果该变量占用内存大小大于字长时, 保证第一次内存访问的首地址是该变量的首地址")
	fmt.Println("---------")
	fmt.Println("  unsafe.Sizeof(v ArbitraryType) uintptr 返回类型 v 本身数据所占用的字节数, 不包含 v 所指向的内容的大小 (返回值是 \"顶层\" 的数据占有的字节数).")
	fmt.Println("  unsafe.Alignof(v ArbitraryType) uintptr 返回类型 v 的对齐方式(即类型 v 在内存中所占用的字节数), 如果是结构体类型的字段的形式, 则返回字段 f 在该结构体中的对齐方式")
	fmt.Println("    对于任意类型的变量 x, 返回值至少为 1")
	fmt.Println("    对于结构体类型的变量 x, 返回值为结构体 x 中每个字段 f 的 unsafe.Alignof(x.f) 中的最大值")
	fmt.Println("    对于数组类型的变量 x, 返回值为构成数组的元素的类型的对齐倍数")
	fmt.Println("  unsafe.Offsetof(v ArbitraryType) uintptr 只适用于结构体中的字段相对于结构体的内存偏移量, 结构体的第 1 个字段的偏移量都是 0, 通过这个函数可以直接根据结构体字段的偏移量对齐进行操作(无论字段是否导出)")
	fmt.Println("    返回类型 v 所代表的结构体字段在结构体中的偏移量, 它必须为结构体类型的字段的形式, 即它返回该结构体起始处与该字段起始处之间的字节数")
	fmt.Println("---------")
	fmt.Println("基本类型内存对齐方式: 对齐系数通常等于变量类型的大小")
	fmt.Println(tab, "例如 int8 内存对齐系数为 1, int16 内存对齐系数为 2, int64 内存对齐系数为 8, float32 内存对齐系数为 4, float64 内存对齐系数为 8")
	fmt.Println(tab, " bool 内存对齐系数为 1, byte 内存对齐系数为 2, rune 内存对齐系数为 4, string 内存对齐系数为 8")
	fmt.Println(tab, "如果当前地址不能被 对齐系数 整除, 则填充空白字节, 直至可以被整除")
	fmt.Println(tab, "数组类型, unsafe.Alignof() 等于构成数组的元素类型的对齐倍数")
	fmt.Println("---------")
	fmt.Println(`
  type Test struct{ 
    id int64
    name string
    age int16
    height float32 
  }`)
	fmt.Println("结构体内存对齐: 其对齐系数等于其所有字段对齐系数中的最大值, Alignof 最大返回 8, 结构体实际大小为对齐系数的整数倍")
	fmt.Println("  Test 结构体的对齐系数为 max(Alignof(id) 8, Alignof(name) 8, Alignof(age) 2, Alignof(height) 4) 中的最大值 8, 在分配一个 Test 对象时, 它的地址必须是 8 的整数倍")
	fmt.Println(" 结构体类型需要考虑两个方面的内存对齐: \033[1;32m内部字段对齐和外部长度填充\033[0m")
	fmt.Println("   内部字段对齐: 指结构体每个字段的偏移量(offset)必须是该字段自身类型大小和该字段对齐系数中较小值的整数倍, 如有需要编译器会在成员之间填充空白字节")
	fmt.Println(tab, "  name string 的对齐系数为 8, 大小为 16, name 字段的内存偏移值必须是 8 的倍数, 当前地址为 8, 不需要填充空白字节")
	fmt.Println(tab, "  height float32 的 AlignOf 为 4, SizeOf 为 4, 当前地址为 26 不是对齐系数 4 的整数倍, 需要填充空白字节至 28")
	fmt.Println("   外部长度填充: 指结构体占用的内存大小必须是结构体中最大成员长度和结构体默认对齐系数中较小值的整数倍")
	fmt.Println(tab, "  如果当前结构体占用内存大小不是整数倍, 则需要在结构体后面填充空白字节")
	fmt.Println(tab, "  Test 结构体最大字段 name 长度为 16, 默认对齐系数为 8, 结构体内存占用大小为 8(id) + 16(name) + 2(age) + 2(空白字节) + 4(height) = 32 是 8 的整数倍")
	fmt.Println("---------")
	fmt.Println("  Test 结构体的对齐系数为 max(Alignof(id) 8, Alignof(name) 8, Alignof(age) 2, Alignof(height) 4) 中的最大值 8, 在分配一个 Test 对象时, 它的地址必须是 8 的整数倍")
	fmt.Println("    Fields             Alignof")
	fmt.Println("  id int64         0  1  2  3  4  5  6  7") // Sizeof(id) = 8	Alignof(id) = 8 // 当前地址为 0, 可以被 8 整除
	// Sizeof(name) = 16	Alignof(name) = 8 // 当前地址为 8, 可以被 8 整除
	fmt.Println("  name string      8  9  10 11 12 13 14 15 16 17 18 19 20 21 22 23")
	fmt.Println("  age int16        24 25 ") // Sizeof(age) = 2	Alignof(age) = 2 // 当前地址为 24, 可以被 2 整除
	fmt.Println("                   —— —— ")
	fmt.Println("  height float32   28 29 30 31") // Sizeof(height) = 4	Alignof(height) = 4 // 当前地址为 26, 距离 26 最近的且可以被 4 整除的地址为 28, 因此需要在前面填充 2 个空白字节
	fmt.Println("  结构体 Test 内存占用大小为 32 字节, 为内存对齐系数 8 的整倍数, 已经符合规则")
	fmt.Println("------")
	var t = Test{1, "zhangsan", 18, 180}
	fmt.Printf("结构体 Test 的内存大小为 %d, 对齐系数为 %d 值为 %+v\n", unsafe.Sizeof(t), unsafe.Alignof(t), t) // 32 8
	fmt.Println("------------------")
	fmt.Println(`
// - 表示空白字节, _ 表示占用字节, 数字表示字段偏移量
type A struct {
  // 0 - - - 4 _ _ _ 8 _ _ _ _ _ _ _ // 结构体占用 16 字节, 是对齐系数 8 的整数倍, Sizeof 16
  a int8  // 1 1
  b int32 // 4 4
  c int64 // 8 8
}
type B struct {
  // 0 _ _ _ - - - - 8 _ _ _ _ _ _ _ 16 - - - - - - - // 结构体占用 17 字节, 不是对齐系数 8 的整数倍, 外部长度需要在末尾填充 7 个空白字节, Sizeof 24
  b int32 // 4 4
  c int64 // 8 8
  a int8  // 1 1
}
type C struct {
  // 0 _ _ _ _ _ _ _ 8 _ _ _ 12 - - - // 结构体占用 13 字节, 不是对齐系数 8 的整数倍, 外部长度需要在末尾填充 3 个空白字节, Sizeof 16
  c int64 // 8 8
  b int32 // 4 4
  a int8  // 1 1
}
type D struct {
  // 0 _ _ _ _ _ _ _ 8 - - - 12 _ _ _ // 结构体占用 16 字节, 是对齐系数 8 的整数倍, Sizeof 16
  c int64 // 8 8
  a int8  // 1 1
  b int32 // 4 4
}`)
	var a = A{1, 2, 3}
	var b = B{2, 3, 1}
	var c = C{3, 2, 1}
	var d = D{3, 1, 2}
	fmt.Printf("A 占用内存大小为 %d, 对齐系数为 %d 值为 %v\n", unsafe.Sizeof(a), unsafe.Alignof(a), a) // 16 8
	fmt.Printf("B 占用内存大小为 %d, 对齐系数为 %d 值为 %v\n", unsafe.Sizeof(b), unsafe.Alignof(b), b) // 24 8
	fmt.Printf("C 占用内存大小为 %d, 对齐系数为 %d 值为 %v\n", unsafe.Sizeof(c), unsafe.Alignof(c), c) // 16 8
	fmt.Printf("D 占用内存大小为 %d, 对齐系数为 %d 值为 %v\n", unsafe.Sizeof(d), unsafe.Alignof(d), d) // 16 8
	fmt.Println("------------------")
}
