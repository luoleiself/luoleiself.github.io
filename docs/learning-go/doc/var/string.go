package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func stringNote() {
	fmt.Println("-----------------stringNote()------------------")
	fmt.Println("string 不是线程安全的")
	fmt.Println("string 是一个不可变的字节序列, 可以包含任意的数据, 底层是一个byte数组, 由于一个字符占用的字节数不确定, 所以无法通过下标的方式获取对应位置的字符")
	fmt.Println(" 只读的采用 UTF-8 编码的字节切片, len() 返回的是字符串字节占用数, 不是 rune 数(这个不太好理解)")
	fmt.Println(" 字符串中超过一个字节表示的字符需要使用 utf8.RuneCountInString() 或者 []rune 计算字符串长度")
	fmt.Println("StringHeader 是字符串的运行时表示，包含指向字节数组的指针和字节数组的长度两个字段共(系统平台差异最大)占用 16 个字节")
	fmt.Println(`
  type StringHeader struct {
    Data uintptr
    Len  int
  }`)
	fmt.Println("-----------------------------------------------")

	var str = "中"
	fmt.Println(len(str), []rune(str), []byte(str)) // 3 [20013] [228 184 173]
	fmt.Println("-----------------")

	s := "hello world"
	// string 0xc000058270 11 "hello world" 16
	fmt.Printf("s 的类型为 %T 地址为 %p 长度为 %d 值为 %q 占用内存大小 %d\n", s, &s, len(s), s, unsafe.Sizeof(s))
	fmt.Println("s[0]= ", s[0])                 // 104
	fmt.Println("s[len(s) - 1]= ", s[len(s)-1]) // 100
	fmt.Println("-----------------")
	s = s + " gg "
	// 0xc000058270 14 "hello world gg " 16
	fmt.Printf("s = s + \" gg \" 的地址为 %p 长度为 %d 值为 %q 占用内存大小为 %d\n", &s, len(s), s, unsafe.Sizeof(s))
	s1 := "he中国"
	// string 0xc0000582b0 8 "he中国" 16
	fmt.Printf("s1 的类型为 %T 地址为 %p 长度为 %d 值为 %q 占用内存大小 %d \n", s1, &s1, len(s1), s1, unsafe.Sizeof(s1))
	fmt.Println(`utf8.RuneCountInString(s1)`)
	// string 0xc0000582b0 4 "he中国" 16
	fmt.Printf("s1 的类型为 %T 地址为 %p 长度为 %d 值为 %q 占用内存大小 %d \n", s1, &s1, utf8.RuneCountInString(s1), s1, unsafe.Sizeof(s1))
	fmt.Println(`len([]rune(s1))`)
	// string 0xc0000582b0 4 "he中国" 16
	fmt.Printf("s1 的类型为 %T 地址为 %p 长度为 %d 值为 %q 占用内存大小 %d \n", s1, &s1, len([]rune(s1)), s1, unsafe.Sizeof(s1))
	fmt.Println("-----------------")
	s4 := "hello"
	// string 0xc0000882f0 5 hello 16
	fmt.Printf("s4 的类型为 %T 地址为 %p 长度为 %d 值为 %q 占用内存大小 %d \n", s4, &s4, len(s4), s4, unsafe.Sizeof(s4))

	fmt.Println("-----------------------------------------------")

	fmt.Println("使用 range 遍历字符串和普通 for 循环的区别")
	var rangeStr = "中国"
	for i, v := range rangeStr {
		fmt.Print(i, v, " ")                // 0 20013  3 22269
		fmt.Printf("i: %d, v: %q \n", i, v) // i: 0 v: '中'  i: 3 v: '国'
	}
	fmt.Println()
	for i := 0; i < len(rangeStr); i++ {
		fmt.Printf("rangeStr[%d] = %d ", i, rangeStr[i]) // rangeStr[0] = 228 rangeStr[1] = 184 rangeStr[2] = 173 rangeStr[3] = 229 rangeStr[4] = 155 rangeStr[5] = 189
	}
	fmt.Println("-----------------------------------------------")

	fmt.Println("修改字符串需要先将字符串转为特定类型的切片")
	s2 := "helloworld@163.com"
	fmt.Printf("s2 的长度为 %d 的值为 %q \n", len(s2), s2) // 18 "helloworld@163.com"
	s3 := []rune(s2)
	fmt.Printf("s3 := []rune(s2) 的类型为 %T 长度为 %d 值为 %q \n", s3, len(s3), s3)
	s3[0] = '北'
	fmt.Println("string(s3) 转换为字符串类型")
	fmt.Printf("s3[0] = '北' 的长度为 %d 值为 %q \n", len(s3), string(s3)) // 18 "北elloworld@163.com"
	fmt.Println("-----------------------------------------------")

	fmt.Println("string byte rune 之间的相互转换")
	fmt.Println(" string => byte")
	fmt.Println("   1. 使用循环遍历字符串, 将遍历的字符追 append 已 make 的 byte 切片中")
	fmt.Println("   2. 使用 []byte(str) 直接转换为 byte 切片")
	fmt.Println("   3. 借助 reflect.stringHeader 包转换")
	fmt.Println(" byte => string")
	fmt.Println("   1. 使用 string(b)")
	fmt.Println("   2. 借助 unsafe.Pointer 包转换")
	fmt.Println(" string => rune")
	fmt.Println("   1. 使用循环遍历字符串, 将遍历的字符串 append 已 make 的 rune 切片中")
	fmt.Println("   2. 使用 []rune(str) 直接转换为 rune 切片")
	fmt.Println(" rune => string")
	fmt.Println("   1. 使用 string(r)")
	fmt.Println(" byte <==> rune")
	fmt.Println("   借助 string 转换")
	fmt.Println("-----------------------------------------------")
}
