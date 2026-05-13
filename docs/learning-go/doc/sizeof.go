package main

import (
	"fmt"
	"unsafe"
)

type Eg struct {
	a byte
	b int16
	c int64
	d int32
}

func unsafeSizeof() {
	fmt.Println("-----------unsafe-----------")
	var b, str, arr, sl, m, ch = true, "he wor", [...]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3}, map[string]string{"h": "H"}, make(chan int, 1)
	var i8 int8 = 1
	var i16 int16 = 2
	var i64 int64 = 3
	var ui8 uint8 = 4
	var ui16 uint16 = 5
	var ui64 uint64 = 6
	var f32 float32 = 1.0
	var f64 float64 = 1.1
	var fn func() string = func() string {
		fmt.Println("hello fn")
		return "hello fn"
	}
	var inter any = "inter"
	fmt.Printf("b: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(b), unsafe.Alignof(b))             // 1 1
	fmt.Printf("i8: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(i8), unsafe.Alignof(i8))          // 1 1
	fmt.Printf("i16: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(i16), unsafe.Alignof(i16))       // 2 2
	fmt.Printf("i64: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(i64), unsafe.Alignof(i64))       // 8 8
	fmt.Printf("ui8: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(ui8), unsafe.Alignof(ui8))       // 1 1
	fmt.Printf("ui16: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(ui16), unsafe.Alignof(ui16))    // 2 2
	fmt.Printf("ui64: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(ui64), unsafe.Alignof(ui64))    // 8 8
	fmt.Printf("f32: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(f32), unsafe.Alignof(f32))       // 4 4
	fmt.Printf("f64: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(f64), unsafe.Alignof(f64))       // 8 8
	fmt.Printf("str: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(str), unsafe.Alignof(str))       // 16 8
	fmt.Printf("arr: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(arr), unsafe.Alignof(arr))       // {48} 8
	fmt.Printf("sl: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(sl), unsafe.Alignof(sl))          // 24 8
	fmt.Printf("m: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(m), unsafe.Alignof(m))             // 8 8
	fmt.Printf("ch: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(ch), unsafe.Alignof(ch))          // 8 8
	fmt.Printf("fn: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(fn), unsafe.Alignof(fn))          // 8 8
	fmt.Printf("inter: Sizeof() %d, Alignof() %d\n", unsafe.Sizeof(inter), unsafe.Alignof(inter)) // 16 8

	var s1, s2, s3 = "", "呵呵", "hello world"
	fmt.Printf("s1 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(s1), unsafe.Alignof(s1), s1) // 16 8 ""
	fmt.Printf("s2 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(s2), unsafe.Alignof(s2), s2) // 16 8 "呵呵"
	fmt.Printf("s3 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(s3), unsafe.Alignof(s3), s3) // 16 8 "hello world"
	fmt.Println("----------------------------")

	var a1 [4]int
	var a2 [6]int
	var a3 [4]float32
	var a4 [5]rune
	fmt.Printf("a1 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(a1), unsafe.Alignof(a1), a1) // 32 8 [4]int{0, 0, 0, 0}
	fmt.Printf("a2 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(a2), unsafe.Alignof(a2), a2) // 48 8 [6]int{0, 0, 0, 0, 0, 0}
	fmt.Printf("a3 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(a3), unsafe.Alignof(a3), a3) // 16 4 [4]float32{0, 0, 0, 0}
	fmt.Printf("a4 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(a4), unsafe.Alignof(a4), a4) // 20 4 [5]int32{0, 0, 0, 0, 0}
	fmt.Println("----------------------------")

	var sl1 []int
	var sl2 []rune
	var sl3 []string
	var sl4 []float32
	fmt.Printf("sl1 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(sl1), unsafe.Alignof(sl1), sl1) // 24 8 []int(nil)
	fmt.Printf("sl2 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(sl2), unsafe.Alignof(sl2), sl2) // 24 8 []int32(nil)
	fmt.Printf("sl3 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(sl3), unsafe.Alignof(sl3), sl3) // 24 8 []string(nil)
	fmt.Printf("sl4 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(sl4), unsafe.Alignof(sl4), sl4) // 24 8 []float32(nil)
	fmt.Println("----------------------------")

	var m1 = map[int]string{1: "1", 2: "2", 3: "3"}
	var m2 = map[string]int{"1": 1, "2": 2}
	var m3 = map[string]float32{"1": 1.0, "2": 2.0, "3": 3.0, "4": 4.0}
	fmt.Printf("m1 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(m1), unsafe.Alignof(m1), m1) // 8 8 map[int]string{1:"1", 2:"2", 3:"3"}
	fmt.Printf("m2 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(m2), unsafe.Alignof(m2), m2) // 8 8 map[string]int{"1":1, "2":2}
	fmt.Printf("m3 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(m3), unsafe.Alignof(m3), m3) // 8 8 map[string]float32{"1":1, "2":2, "3":3, "4":4}
	fmt.Println("----------------------------")

	var c1, c2, c3 = make(chan int), make(chan uint8, 5), make(chan struct{})
	fmt.Printf("c1 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(c1), unsafe.Alignof(c1), c1) // 8 8 (chan int)(0xc00001c0c0)
	fmt.Printf("c2 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(c2), unsafe.Alignof(c2), c2) // 8 8 (chan uint8)(0xc0000240e0)
	fmt.Printf("c3 的内存占用大小为 %d 内存对齐系数为 %d 值为 %#v\n", unsafe.Sizeof(c3), unsafe.Alignof(c3), c3) // 8 8 (chan struct {})(0xc00001c120)
	fmt.Println("----------------------------")

	var f1 = func() {
		fmt.Println("f1")
	}
	var f2 = func() {
		fmt.Println("f2")
		fmt.Println("f2")
	}
	fmt.Printf("f1 的内存占用大小为 %d 内存对齐系数为 %d\n", unsafe.Sizeof(f1), unsafe.Alignof(f1)) // 8 8 (func())(0x4d0ce0)
	fmt.Printf("f2 的内存占用大小为 %d 内存对齐系数为 %d\n", unsafe.Sizeof(f2), unsafe.Alignof(f2)) // 8 8 (func())(0x4d0d60)
	fmt.Println("----------------------------")

	eg := Eg{}
	sa := unsafe.Sizeof(eg.a)
	sb := unsafe.Sizeof(eg.b)
	sc := unsafe.Sizeof(eg.c)
	sd := unsafe.Sizeof(eg.d)
	fmt.Printf("sa %d sb %d sc %d sd %d\n", sa, sb, sc, sd)    // 1 2 8 4
	fmt.Printf("sa + sb + sc + sd = %d\n", sa+sb+sc+sd)        // 15
	fmt.Printf("unsafe.Sizeof(eg) 值为 %d\n", unsafe.Sizeof(eg)) // 24
	fmt.Println(eg)
	// 各字段的内存大小总和是小于结构体所占用的内存, 因为编译器在字段中间加上填充字节
	// 如果将 c 和 d 的位置调换之后, unsafe.Sizeof(eg) 返回 16
	fmt.Println("----------------------------")
}
