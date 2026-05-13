package main

import (
	"fmt"
	// "testing"
)

func sliceNote() {
	fmt.Println("--------------sliceNote()---------------")
	fmt.Println("slice 是引用类型, 底层是对数组的一个连续片段的引用, 切片一般用于快速地操作一块数据集合")
	fmt.Printf("只声明未初始化的 slice 长度为 0 无法进行操作, 提示: panic: runtime error: index out of range [0] with length 0\n")
	fmt.Println("对切片下标越界的元素的操作会编译失败")
	fmt.Println(tab, "切片的容量为切片起始位置到对象的最后位置的长度")
	fmt.Println(tab, "子切片的容量为原切片的容量减去子切片的起始偏移量")
	fmt.Println("\033[1;32m声明方式1\033[0m: 初始化 slc := []string{\"hello\", \"world\", \"haHa\"}")
	fmt.Println("\033[1;32m声明方式2\033[0m: make方法 slc1 := make([]int, length, [capacity])")
	fmt.Println("\033[1;32m声明方式3\033[0m: 通过数组截取 slc2 := arr[startIndex:endIndex:maxIndex] 会产生多个切片共享剩余容量")
	fmt.Println(tab, "新切片: len = endIndex - startIndex, cap = maxIndex - startIndex")
	fmt.Println(tab, "startIndex 截取的起始下标, 默认从 0 开始")
	fmt.Println(tab, "endIndex 截取的结束下标(不包含), 默认到末尾")
	fmt.Println(tab, "maxIndex 为新切片保留的原切片的最大下标(不包含), 默认为截取对象的容量")
	fmt.Println(tab, "切片声明方式同样支持数组声明的指定下标声明")
	fmt.Println("派生表达式: startIndex:endIndex 此方式可能会造成所有原切片中未使用的容量都和子切片共享, 导致多个切片添加并互相覆盖数据")
	fmt.Println("完全派生表达式: startIndex:endIndex:maxIndex")
	fmt.Println(tab, "maxIndex 定义了原切片容量中的最后位置, 用来确定新切片可以使用的容量, maxIndex - startIndex 为新切片的实际容量")
	fmt.Println("-----------------------------")

	fmt.Println("SliceHeader 是运行时的切片, 它不保证使用的可移植性, 安全性;")
	fmt.Println(`
  type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
  }`)
	fmt.Println("-----------------------------")

	a0 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Printf("a0 的类型为 %T 长度为 %d 容量为 %d 值为 %v\n", a0, len(a0), cap(a0), a0) // 类型为 [11]string 长度 11 容量 11 值为 [a b c d e f g h i j k]
	s0 := a0[2:8]
	fmt.Printf("s0 := a0[2:8] 的类型为 %T 长度为 %d 容量为 %d 值为 %v\n", s0, len(s0), cap(s0), s0) // 类型为 []string 长度 6 容量 9 值为 [c d e f g h]
	s1 := a0[4:7]
	fmt.Printf("s1 := a0[4:7] 的类型为 %T 长度为 %d 容量为 %d 值为 %v\n", s1, len(s1), cap(s1), s1) // 类型为 []string 长度 3 容量 7 值为 [e f g]
	s2 := s0[3:9]
	fmt.Printf("s2 := s0[3:9] 的类型为 %T 长度为 %d 容量为 %d 值为 %v\n", s2, len(s2), cap(s2), s2) // 类型为 []string 长度 6 容量 6 值为 [f g h i j k]
	s3 := s1[4:7]
	fmt.Printf("s3 := s1[4:7] 的类型为 %T 长度为 %d 容量为 %d 值为 %v\n", s3, len(s3), cap(s3), s3) // 类型为 []string 长度 3 容量 3 值为 [i j k]
	fmt.Println("-----------------------------")

	fmt.Println("\033[1;31m新切片容量未超出则共享剩余容量并互相覆盖数据\033[0m")
	var x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	var y = x[:2]
	var z = x[2:]
	fmt.Print("var x = make([]int, 0, 5)\nx = append(x, 1, 2, 3, 4)\nvar y = x[:2]\nvar z = x[2:]\n")
	fmt.Printf("x 长度为 %d 容量为 %d 值为 %v\n", len(x), cap(x), x) // x 长度为 4 容量为 5 值为 [1 2 3 4]
	fmt.Printf("y 长度为 %d 容量为 %d 值为 %v\n", len(y), cap(y), y) // y 长度为 2 容量为 5 值为 [1 2]
	fmt.Printf("z 长度为 %d 容量为 %d 值为 %v\n", len(z), cap(z), z) // z 长度为 2 容量为 3 值为 [3 4]
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Print("y = append(y, 30, 40, 50)\nx = append(x, 60)\nz = append(z, 70)\n")
	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z) // x: [1 2 30 40 70], y: [1 2 30 40 70], z: [30 40 70]
	fmt.Println("------------")
	fmt.Println("\033[1;32m解决办法 1.\033[0m 新切片容量超出原切片容量创建新切片则不会影响原切片")
	var x1 = make([]int, 0, 5)
	x1 = append(x1, 1, 2, 3, 4)
	var y1 = x1[:2]
	var z1 = x1[2:]
	fmt.Print("var x1 = make([]int, 0, 5)\nx1 = append(x1, 1, 2, 3, 4)\nvar y1 = x1[:2]\nvar z1 = x1[2:]\n")
	fmt.Printf("x1 长度为 %d 容量为 %d 值为 %v\n", len(x1), cap(x1), x1) // x1 长度为 4 容量为 5 值为 [1 2 3 4]
	fmt.Printf("y1 长度为 %d 容量为 %d 值为 %v\n", len(y1), cap(y1), y1) // y1 长度为 2 容量为 5 值为 [1 2]
	fmt.Printf("z1 长度为 %d 容量为 %d 值为 %v\n", len(z1), cap(z1), z1) // z1 长度为 2 容量为 3 值为 [3 4]
	y1 = append(y1, 30, 40, 50, 60)
	x1 = append(x1, 70, 700)
	z1 = append(z1, 80, 90, 100)
	fmt.Print("y1 = append(y1, 30, 40, 50, 60)\nx1 = append(x1, 70, 700)\nz1 = append(z1, 80, 90, 100)\n")
	fmt.Printf("x1: %v, y1: %v, z1: %v\n", x1, y1, z1) // x1: [1 2 3 4 70 700], y1: [1 2 30 40 50 60], z1: [3 4 80 90 100]
	fmt.Println("------------")
	fmt.Println("\033[1;32m解决办法 2.\033[0m 使用完全派生表达式")
	var x2 = make([]int, 0, 5)
	x2 = append(x2, 1, 2, 3, 4)
	var y2 = x2[:2:2]
	var z2 = x2[2:4:4]
	fmt.Print("var x2 = make([]int, 0, 5)\nx2 = append(x2, 1,2,3,4)\nvar y2 = x2[:2:2]\nvar z2 = x2[2:4:4]\n")
	fmt.Printf("x2 长度为 %d 容量为 %d 值为 %v\n", len(x2), cap(x2), x2) // x2 长度为 4 容量为 5 值为 [1 2 3 4]
	fmt.Printf("y2 长度为 %d 容量为 %d 值为 %v\n", len(y2), cap(y2), y2) // y2 长度为 2 容量为 2 值为 [1 2]
	fmt.Printf("z2 长度为 %d 容量为 %d 值为 %v\n", len(z2), cap(z2), z2) // z2 长度为 2 容量为 2 值为 [3 4]
	y2 = append(y2, 30, 40, 50, 60)
	x2 = append(x2, 70, 700)
	z2 = append(z2, 80, 90, 100)
	fmt.Print("y2 = append(y2, 30, 40, 50, 60)\nx2 = append(x2, 70, 700)\nz2 = append(z2, 80, 90, 100)\n")
	fmt.Printf("x2: %v, y2: %v, z2: %v\n", x2, y2, z2) // x2: [1 2 3 4 70 700], y2: [1 2 30 40 50 60], z2: [3 4 80 90 100]
	fmt.Println("-----------------------------")

	fmt.Println("slice 复制切片, 如果两个切片的长度不一样, 将按照其中较小的切片的元素个数进行复制")
	fmt.Println("内置函数复制切片 copy(dst, src []Type) int")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	fmt.Printf("slice1 = %v slice2 = %v\n", slice1, slice2) // slice1 = [1 2 3 4 5] slice2 = [6 7 8]
	copy(slice2, slice1)
	fmt.Printf("copy(slice2, slice1);\nslice1 = %v slice2 = %v\n\n", slice1, slice2) // slice1 = [1 2 3 4 5] slice2 = [1 2 3]
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{6, 7, 8}
	fmt.Printf("slice3 = %v slice4 = %v\n", slice3, slice4) // slice3 = [1 2 3 4 5] slice4 = [6 7 8]
	copy(slice3, slice4)
	fmt.Printf("copy(slice3, slice4);\nslice3 = %v slice4 = %v\n", slice3, slice4) // slice3 = [6 7 8 4 5] slice4 = [6 7 8]
	fmt.Println("-----------------------------")

	fmt.Println("--------切片作为函数参数传递--------")
	var ts6 = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("ts6 的值为 %v \n", ts6) // [1 2 3 4 5 6]
	fmt.Println("调用函数 testSlice(ts6)")
	testSlice(ts6)                            // [100 2 3 4 5 6]
	fmt.Printf("main() 内的 ts6 的值为 %v\n", ts6) // [100 2 3 4 5 6]
	fmt.Println("-----------------------------")

	ts7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("ts7 = %v 使用移位法删除指定元素 %v\n", ts7, DeleteSlice3(ts7, 6)) // [1 2 3 4 5 7 8 9 10]
	fmt.Println("-----------------------------")
	// BenchmarkDeleteSlice4(&testing.B{})
}

// 函数内修改参数切片
func testSlice(ts6 []int) {
	ts6[0] = 100
	fmt.Printf("testSlice func 内修改 ts6[0] = 100 后的 ts6 的值为 %v\n", ts6)
}

// 截取法删除指定元素, 修改原切片, 删除元素时下标左移一位
func DeleteSlice1(a []int, elem int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == elem {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}

// 拷贝法删除指定元素, 不修改原切片, 需要重新初始化一个新切片
func DeleteSlice2(a []int, elem int) []int {
	tmp := make([]int, 0, len(a))
	for _, v := range a {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

// 移位法删除指定元素, 修改原切片
func DeleteSlice3(a []int, elem int) []int {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// 移位法删除指定元素,底层共享数组
func DeleteSlice4(a []int, elem int) []int {
	tgt := a[:0]
	for _, v := range a {
		if v != elem {
			tgt = append(tgt, v)
		}
	}
	return tgt
}

// func getSlice(n int) []int {
// 	a := make([]int, 0, n)
// 	for i := 0; i < n; i++ {
// 		if i%2 == 0 {
// 			a = append(a, 0)
// 			continue
// 		}
// 		a = append(a, 1)
// 	}
// 	return a
// }

// func BenchmarkDeleteSlice4(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = DeleteSlice3(getSlice(10), 0)
// 	}
// }
