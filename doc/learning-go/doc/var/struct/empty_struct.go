package structdoc

import (
	"fmt"
	"unsafe"
)

func EmptyStruct() {
	fmt.Println("空结构体: 不包含任何字段的结构体, 可以拥有方法")
	fmt.Println(tab, "零内存占用, 空结构体不占用任何内存空间")
	fmt.Println(tab, "地址相同, 无论创建多少个空结构体, 它们所指向的地址都是相同的")
	fmt.Println(tab, "无状态, 空结构体不包含任何字段, 因此不能有状态")
	var a int
	var b string
	var e struct{}
	var e1 struct{}
	fmt.Printf("unsafe.Sizeof(a) %v\n", unsafe.Sizeof(a)) // 8
	fmt.Printf("unsafe.Sizeof(b) %v\n", unsafe.Sizeof(b)) // 16
	fmt.Printf("unsafe.Sizeof(e) %v\n", unsafe.Sizeof(e)) // 0
	fmt.Printf("空结构体 e 的地址为 %p\n", &e)                    // 0x434f00
	fmt.Printf("空结构体 e1 的地址为 %p\n", &e1)                  // 0x434f00
	fmt.Printf("&e == &e1 %t\n", &e == &e1)               // true
	fmt.Println("-----------------")
	fmt.Println("runtime 包下的 malloc.go 源码中, 当要分配的对象大小 size 为 0 时, 会返回指向 zerobase 的指针,")
	fmt.Println(" zerobase 是一个用于分配零字节对象的基准地址, 它不占用任何实际的内存空间")
	fmt.Println(`
  var zerobase uintptr
	func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
    ······
    if size == 0 {
      return unsafe.Pointer(&zerobase)
    }
    ......`)
	fmt.Println("使用场景")
	fmt.Println(tab, "1. 实现 Set 集合类型, go 没有提供 Set 集合类型, 使用 map 模拟实现一个 Set 集合")
	fmt.Println(tab, tab, "由于 map 的 key 具有唯一性, 将元素存储为 key, value 没有实际作用, 使用空结构体作为 value 节省内存")
	fmt.Println(tab, "2. 用于通道信号, 空结构体通常用于 Goroutine 之间的信号传递, 尤其是不关心通道中传递的具体数据, 只需要一个触发信号时")
	fmt.Println(tab, "3. 作为方法接收器, 使用空结构体实现接口定义的方法")
	fmt.Println("----------------------------------------")
}
