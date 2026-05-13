package main

import (
	"fmt"
	"os"
	"strings"
)

func ScanNote() {
	fmt.Println("---------------ScanNote()---------------")
	var i int
	n, err := fmt.Fscan(strings.NewReader("100"), &i)
	if err != nil {
		fmt.Println("fmt.Fscan err =", err)
	}
	fmt.Printf("i = %v 成功扫描的条目个数为 %d\n", i, n) // 100 1
	fmt.Println("-------------------------------------")
	var (
		i2 int
		b2 bool
		s2 string
	)
	r2 := strings.NewReader("5 true gophers")
	n2, err := fmt.Fscanf(r2, "%d %t %s", &i2, &b2, &s2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Printf("i2 = %v, b2 = %v, s2 = %v 成功扫描的条目个数为 %d\n", i2, b2, s2, n2) // 5 true gophers 3
	fmt.Println("-------------------------------------")

	var i3 string
	n3, err := fmt.Sscan("hello world", &i3)
	if err != nil {
		fmt.Println("fmt.Sscan err =", err)
	}
	fmt.Printf("i3 = %v 成功扫描的条目个数为 %d\n", i3, n3) // hello 1
	fmt.Println("-------------------------------------")
	var (
		i4 int
		s4 string
	)
	n4, err := fmt.Sscanf("1234567 ", "%5s%d", &s4, &i4)
	if err != nil {
		fmt.Println("fmt.Sscanf err =", err)
	}
	fmt.Printf("i4 = %v, s4 = %v 成功扫描的条目个数为 %d\n", i4, s4, n4) // 67 12345 2
	fmt.Println("-------------------------------------")

	var i5, i5_1 int
	fmt.Print("请输入两个用空格分隔的数字: ")
	n5, err := fmt.Scan(&i5, &i5_1)
	if err != nil {
		fmt.Println("fmt.Scan err =", err)
	}
	fmt.Printf("i5 = %v, i5_1 = %v, i5 + i5_1 = %d 成功扫描的条目个数为 %d\n", i5, i5_1, i5+i5_1, n5) // 1 2 3 2
	fmt.Println("-------------------------------------")

	// var i6 int
	// var s6 string
	// fmt.Println("请按照模板输入: go hello 10")
	// n6, err := fmt.Scanf("go %s %d", &s6, &i6)
	// if err != nil {
	// 	fmt.Println("fmt.Scanf err =", err)
	// }
	// fmt.Printf("模板扫描不好玩, 扫出来的结果: %d %v 成功扫描的条目个数为 %d\n", i6, s6, n6) // 0 0
	fmt.Println("-------------------------------------")
}
