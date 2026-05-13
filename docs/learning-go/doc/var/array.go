package main

import (
	"fmt"
)

func arrayNote() {
	fmt.Println("-------------arrayNote()----------------")
	fmt.Println("声明: 数组是值类型")
	fmt.Println("var 变量名 [数组长度] 数组长度也可以是一个表达式 ---- 数组是由一个固定长度的特定类型元素组成的序列")
	fmt.Println("var arr [4]int // 数组中包含4个元素int类型默认值")
	fmt.Print("初始化声明 arr := [length]array_type{value,...}\n")
	fmt.Println("使用 ... 声明数组只能用在数组声明初始化时, 如果作为函数参数中的数组长度时则编译失败, syntax error: unexpected ..., expected expression")
	fmt.Println("-----------------------------")

	var t [4]bool
	fmt.Println("\033[1;32m声明方式1\033[0m: var t [4]bool")
	fmt.Printf("var t [4]bool 类型为 %T 长度为 %d 值为 %v\n", t, len(t), t) // [4]bool  4  [false false false false]
	arr := [...]float32{3.14, 6.6, 7.89, 9.19}
	fmt.Println("\033[1;32m声明方式2\033[0m: arr := [...]float32{3.14, 6.6, 7.89, 9.19}")
	fmt.Printf("类型为 %T 内存地址为 %p 长度为 %d 值为 %v\n", arr, &arr, len(arr), arr) // [4]float32  0xc0000109f0  4  [3.14 6.6 7.89 9.19]
	fmt.Println("----")
	arr3 := [...]int{1: 6, 0: 5, 2: 7, 3: 10, 5: 9}
	fmt.Println("\033[1;32m声明方式3\033[0m: 使用数字索引: arr3 := [...]int{1: 6, 0: 5, 2: 7, 3: 10, 5: 9}")
	fmt.Printf("类型为 %T 内存地址为 %p 长度为 %d 值为 %v\n", arr3, &arr3, len(arr3), arr3) // [6]int  0xc000014660  6  [5 6 7 10 0 9]
	fmt.Println("----")
	arr6 := [...]int{'a': 6, 'b': 5, 'c': 7}
	fmt.Println("\033[1;32m声明方式4\033[0m: 使用字符索引: arr4 := [...]int{'a': 6, 'b': 5, 'c': 7}")
	fmt.Println("字符索引本质上是 rune 类型")
	fmt.Println("len(arr6)", len(arr6)) // len(arr6) 100
	fmt.Println("-----------------------------")

	fmt.Println("值类型")
	arr2 := arr
	arr2[0] = 400
	fmt.Print("arr2 := arr\narr2[0] = 400\n")
	fmt.Printf("arr 的值为 %v\n", arr)  // [0 1 2 3]
	fmt.Printf("arr2 值为 %v\n", arr2) // [400 1 2 3]
	fmt.Println("------------")

	fmt.Println("指针类型, 可以使用 * 指针取值或者由编译器自动类型推断")
	arr4 := &arr
	arr4[2] = 250
	(*arr4)[0] = 500
	fmt.Print("arr4 := &arr\narr4[2] = 250\n(*arr4)[0] = 500\n")
	fmt.Printf("arr 的值为 %v\n", arr)  // [500 1 250 3]
	fmt.Printf("arr3 值为 %v\n", arr4) // &[500 1 250 3]
	fmt.Println("------------")

	var test3 = [3]int{1, 2, 3}
	testFunc(&test3)                    // *[3]int  [1 250 333]
	fmt.Printf("test3 的值为 %v\n", test3) //  [1 250 333]
	fmt.Println("-----------------------------")

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Print("a := [2]int{1, 2}\nb := [...]int{1, 2}\nc := [2]int{1, 3}\n")
	fmt.Printf("a == b 的值为 %t a == c 的值为 %t b == c 的值为 %t\n", a == b, a == c, b == c) // true false false
	fmt.Println("-----------------------------")

	fmt.Println("数组的地址为第一个元素的内存地址，第二个元素开始依次的地址为上一个元素的内存地址加上元素类型的内存占用字节数")
	var test1 [4]int32
	fmt.Println(test1) // [0 0 0 0]
	// [4]int32  0xc0000109b0  0xc0000109b0  0xc0000109b4  0xc0000109b8
	fmt.Printf("数组 test1 类型为 %T 地址为 %p 第1元素地址为 %p 第2元素地址为 %p 第3元素地址为 %p\n", test1, &test1, &test1[0], &test1[1], &test1[2])
	var test2 [4]int64
	// [4]int64  0xc00001a280  0xc00001a280  0xc00001a288  0xc00001a290
	fmt.Printf("数组 test2 类型为 %T 地址为 %p 第1元素地址为 %p 第2元素地址为 %p 第3元素地址为 %p\n", test2, &test2, &test2[0], &test2[1], &test2[2])

	fmt.Println("-------------")
	var arr5 [3][4]uint32
	fmt.Println("二维数组 arr5 ", arr5)                         // [[0 0 0 0] [0 0 0 0] [0 0 0 0]]
	fmt.Printf("二维数组 arr5 的类型为 %T 的地址为 %p \n", arr5, &arr5) // [3][4]uint32	0xc000010900
	for i, v := range arr5 {
		fmt.Printf("arr5[%d] 内存地址为 %p\n", i, &arr5[i])
		for j := range v {
			fmt.Printf("arr5[%d][%d] 内存地址为 %p\n", i, j, &arr5[i][j])
			/*
				[3][4]uint16 内存占用
				arr5[0] 内存地址为 0xc00008c000
				arr5[0][0] 内存地址为 0xc00008c000
				arr5[0][1] 内存地址为 0xc00008c002
				arr5[0][2] 内存地址为 0xc00008c004
				arr5[0][3] 内存地址为 0xc00008c006
				arr5[1] 内存地址为 0xc00008c008
				arr5[1][0] 内存地址为 0xc00008c008
				arr5[1][1] 内存地址为 0xc00008c00a
				arr5[1][2] 内存地址为 0xc00008c00c
				arr5[1][3] 内存地址为 0xc00008c00e
				arr5[2] 内存地址为 0xc00008c010
				arr5[2][0] 内存地址为 0xc00008c010
				arr5[2][1] 内存地址为 0xc00008c012
				arr5[2][2] 内存地址为 0xc00008c014
				arr5[2][3] 内存地址为 0xc00008c016
			*/

			/*
				[3][4]uint32 内存占用
				arr5[0] 内存地址为 0xc000010900
				arr5[0][0] 内存地址为 0xc000010900
				arr5[0][1] 内存地址为 0xc000010904
				arr5[0][2] 内存地址为 0xc000010908
				arr5[0][3] 内存地址为 0xc00001090c
				arr5[1] 内存地址为 0xc000010910
				arr5[1][0] 内存地址为 0xc000010910
				arr5[1][1] 内存地址为 0xc000010914
				arr5[1][2] 内存地址为 0xc000010918
				arr5[1][3] 内存地址为 0xc00001091c
				arr5[2] 内存地址为 0xc000010920
				arr5[2][0] 内存地址为 0xc000010920
				arr5[2][1] 内存地址为 0xc000010924
				arr5[2][2] 内存地址为 0xc000010928
				arr5[2][3] 内存地址为 0xc00001092c
			*/
		}
	}
	fmt.Println("-----------------------------")

	// testFunc1([...]int{1,2,3,4})
}
func testFunc(arr *[3]int) {
	fmt.Println("函数内修改指针类型参数的值的方式 func testFunc(arr *[3]int)")
	arr[1] = 250    // 由编译器自动推断类型
	(*arr)[2] = 333 // 使用 * 指针取值方式获取
	fmt.Println("arr[1] = 250 // 由编译器自动推断类型")
	fmt.Println("(*arr)[2] = 333 // 使用 * 指针取值方式获取")
	fmt.Printf("arr 的类型为 %T 值为 %v\n", arr, *arr)
}

// func testFunc1(arr [...]int){ // 编译报错 syntax error: unexpected ..., expected expression
// 	fmt.Println(arr)
// }
