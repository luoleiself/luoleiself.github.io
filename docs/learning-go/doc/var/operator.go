package main

import (
	"fmt"
	"math"
)

func Operator() {
	fmt.Println("-------------Operator---------------")
	fmt.Println("_ 匿名变量(特殊标识符), 任何赋值给这个标识符的值都将被抛弃")
	fmt.Println("-------------")
	fmt.Println("// golang 不支持 ** 幂运算操作符, 以下只是注释示例, 需要使用 math.Pow math.Exp2 等函数计算幂运算")
	fmt.Printf("int8 %.f~%.f, uint8 %d~%d\n", -math.Exp2(7), math.Exp2(7)-1, 0, uint8(math.Exp2(8)-1))       // int8 -128~127, uint8 0~255
	fmt.Printf("int16 %.f~%.f, uint16 %d~%d\n", -math.Exp2(15), math.Exp2(15)-1, 0, uint16(math.Exp2(16)-1)) // int16 -32768~32767, uint16 0~65535
	fmt.Printf("int32 %.f~%.f, uint32 %d~%d\n", -math.Exp2(31), math.Exp2(31)-1, 0, uint32(math.Exp2(32)-1)) // int32 -2147483648~2147483647, uint32 0~4294967295
	// fmt.Printf("int64 %d~%d, uint64 %d~%d\n", -int64(math.Exp2(63)), int64(math.Exp2(63))-1, 0, uint64(math.Exp2(64)-1))
	fmt.Println("int64 -9223372036854775808~9223372036854775807, uint64 0~18446744073709551615")
	fmt.Println("-------------")

	fmt.Println("算术运算符: +, -, *, /, %, ++, --")
	fmt.Println("关系运算符: ==, !=, >, >=, <, <=")
	fmt.Println("赋值运算符: =, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=")
	fmt.Println("逻辑运算符: &&, ||, !")
	fmt.Println("位运算符: &, |, ^, <<, >>")
	fmt.Println("其他操作符: &(取地址), *(指针变量)")
}
