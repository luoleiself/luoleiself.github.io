package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func typeConversionNote() {
	fmt.Println("-------------typeConversionNote()--------------")
	var a int8 = 10
	var b float32 = float32(a)
	var c = float32(a) + b
	var d = a + int8(b)
	fmt.Printf("var a int8 = 10 声明的 a 的类型 %T 值为 %v 占用字节数 %d\n", a, a, unsafe.Sizeof(a))                           // int8 10 1
	fmt.Printf("var b float32 = float32(a) 转换后的 b 的类型 %T 值为 %v 占用字节数 %d a 的类型为 %T \n", b, b, unsafe.Sizeof(b), a) // float32 10 4 int8
	fmt.Printf("var c = float32(a) + b 的 c 的类型为 %T 的值为 %v 占用字节数 %d\n", c, c, unsafe.Sizeof(c))                    // float32 20 4
	fmt.Printf("var d = a + int8(b) 的 d 的类型为 %T 的值为 %v 占用字节数 %d\n", d, d, unsafe.Sizeof(d))                       // int8 20 1
	fmt.Println("-------------")
	fmt.Println("或者可以使用 fmt.Sprintf(fmt string, a ...interface{}) string 方法返回字符串")
	fmt.Println("---------------------------")

	fmt.Println("字符串解析失败则取变量的默认值")
	failure1, _ := strconv.ParseInt("123hello", 0, 16)
	failure2, _ := strconv.ParseUint("123hello", 0, 16)
	failure3, _ := strconv.ParseFloat("123hello", 32)
	failure4, _ := strconv.ParseBool("123hello")
	fmt.Printf("failure1, _ := strconv.ParseInt(\"123hello\", 0, 16) failure1 的类型为 %T 值为 %v\n", failure1, failure1)  // int64 0
	fmt.Printf("failure2, _ := strconv.ParseUint(\"123hello\", 0, 16) failure2 的类型为 %T 值为 %v\n", failure2, failure2) // unit64 0
	fmt.Printf("failure3, _ := strconv.ParseFloat(\"123hello\", 32) failure3 的类型为 %T 值为 %v\n", failure3, failure3)   // float64 0
	fmt.Printf("failure4, _ := strconv.ParseBool(\"123hello\") failure4 的类型为 %T 值为 %v\n", failure4, failure4)        // bool false
	fmt.Println("---------------------------")

	fmt.Print("整形和字符串之间转换\n")
	fmt.Println("func FormatInt(i int64, base int) string")
	fmt.Println("  第一个参数 int64 类型, 第二个参数指定 2-36 之间的进制")
	fmt.Println("func ParseInt(s string, base int, bitSize int) (i int64, err error)")
	fmt.Println("  第二个参数为 2-36 之间的进制,如果为 0 则自动推导参数类型,\n  第三个参数为转换后的结果类型, 默认返回 int64")
	fmt.Println("------------")
	num2string := strconv.FormatInt(int64(127), 16) // 或者可以使用 fmt.Sprintf() 方法返回字符串
	num2string2num, _ := strconv.ParseInt(num2string, 16, 32)
	fmt.Printf("num2string := strconv.FormatInt(int64(127), 16) num2string 的类型为 %T 值为 %q\n", num2string, num2string)                      // string "7f"
	fmt.Printf("num2string2num,_ := strconv.ParseInt(num2string, 16, 32) num2string2num 的类型为 %T 值为 %d\n", num2string2num, num2string2num) // int64 127
	fmt.Println("------------")
	n2s := strconv.FormatInt(int64(110), 8)
	s2n, _ := strconv.ParseInt(n2s, 10, 32)
	s2n8, _ := strconv.ParseInt(n2s, 8, 32)
	fmt.Printf("n2s := strconv.FormatInt(int64(110), 8) n2s 的类型为 %T 值为 %v\n", n2s, n2s)    // string 156
	fmt.Printf("s2n, _ := strconv.ParseInt(n2s, 10, 32) s2n 的类型为 %T 值为 %v\n", s2n, s2n)    // int64 156
	fmt.Printf("s2n8, _ := strconv.ParseInt(n2s, 8, 32) s2n8 的类型为 %T 值为 %v\n", s2n8, s2n8) // int64 110
	fmt.Println("------------")
	fmt.Println("无符号整形")
	fmt.Println("func FormatUint(i uint64, base int) string")
	fmt.Println("  第一个参数 uint64 类型, 第二个参数指定 2-36 之间的进制")
	fmt.Println("func ParseUint(s string, base int, bitSize int) (n uint64, err error)")
	fmt.Println("  第二个参数为 2-36 之间的进制,如果为 0 则自动推导参数类型,\n  第三个参数为转换后的结果类型, 默认返回 uint64")
	fmt.Println("------------")
	nnum2string := strconv.FormatUint(uint64(256), 8) // 第二个参数指定 2-36 之间的进制
	nnum2string2nnum, _ := strconv.ParseUint(nnum2string, 16, 32)
	nnum2string2nnum8, _ := strconv.ParseUint(nnum2string, 8, 32)
	fmt.Printf("nnum2string := strconv.FormatUint(uint64(256), 8) nnum2string 的类型为 %T 值为 %q\n", nnum2string, nnum2string)                               // string "400"
	fmt.Printf("nnum2string2nnum,_ := strconv.ParseUint(nnum2string, 16, 32) nnum2string2nnum 的类型为 %T 值为 %v\n", nnum2string2nnum, nnum2string2nnum)     // uint64 1024
	fmt.Printf("nnum2string2nnum8, _ := strconv.ParseUint(nnum2string, 8, 32) nnum2string2nnum8 的类型为 %T 值为 %v\n", nnum2string2nnum8, nnum2string2nnum8) // uint64 256
	fmt.Println("strings.ToUpper(strconv.FormatUint(uint64(255), 16))", strings.ToUpper(strconv.FormatUint(uint64(255), 16)))                           // FF
	fmt.Println("---------------------------")

	fmt.Print("浮点型和字符串之间转换\n")
	fmt.Println("func FormatFloat(f float64, fmt byte, prec, bitSize int) string")
	fmt.Println("  第二个参数为 fmt 格式: 取值 f(-ddd.dddd), b(-ddddp+-ddd, 指数为二进制), e(-d.dddde+-dd, 十进制指数), E(-d.ddddE+-ddd, 十进制指数),")
	fmt.Println("    g(指数很大时用 e 格式, 否则用 f 格式), G(指数很大时用 E 格式, 否则用 f 格式)")
	fmt.Println("  第三个参数为精度prec, 对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G', 它控制总的数字个数。如果prec 为-1, 则代表使用最少数量的、但又必需的数字来表示f")
	fmt.Println("  第四个参数表示来源类型(32: float32、64: float64), 会据此进行舍入")
	fmt.Println("func ParseFloat(s string, bitSize int) (f float64, err error)")
	fmt.Println("  第二个参数 bitSize 指定了期望的接收类型, 默认返回 float64")
	fmt.Println("------------")
	f2string := strconv.FormatFloat(float64(123.456), 'f', -1, 32)
	f2string2f, _ := strconv.ParseFloat(f2string, 32)
	fmt.Printf("f2string := strconv.FormatFloat(float64(123.456), 'f', -1, 32) f2string 的类型为 %T 值为 %q\n", f2string, f2string) // string "123.456"
	fmt.Printf("f2string2f,_ := strconv.ParseFloat(f2string, 32) f2string2f 的类型为 %T 值为 %f\n", f2string2f, f2string2f)         // float64 123.456001
	fmt.Println("------------")
	float2string := strconv.FormatFloat(float64(789.110), 'e', 6, 32)
	float2string2float, _ := strconv.ParseFloat(float2string, 32)
	fmt.Printf("float2string := strconv.FormatFloat(float64(789.110), 'E', 6, 32) float2string 的类型为 %T 值为 %v\n", float2string, float2string)               // string  7.891100e+02
	fmt.Printf("float2string2float, _ := strconv.ParseFloat(float2string, 32) float2string2float 的类型为 %T 值为 %v\n", float2string2float, float2string2float) // float64 789.1099853515625
	fmt.Println("---------------------------")

	fmt.Println("布尔和字符串之间的转换")
	fmt.Println("func FormatBool(b bool) string")
	fmt.Println("func ParseBool(str string) (value bool, err error)")
	fmt.Println("  返回字符串表示的 bool 值, 接受 1, 0, t, f, T, F, true, false, True, False, TRUE, FALSE; 否则返回错误")
	fmt.Println("------------")
	bl2string := strconv.FormatBool(true)
	bl2string2bl, _ := strconv.ParseBool(bl2string)
	fmt.Printf("bl2string := strconv.FormatBool(true) bl2string 的类型为 %T 值为 %q\n", bl2string, bl2string)                   // string "true"
	fmt.Printf("bl2string2bl,_ := strconv.ParseBool(bl2string) bl2string2bl 的类型为 %T 值为 %t\n", bl2string2bl, bl2string2bl) // bool true
	fmt.Println("------------")
	b1, _ := strconv.ParseBool("1")
	fmt.Printf("strconv.ParseBool(\"1\") 的类型为 %T 值为 %t\n", b1, b1) // bool true
	b2, _ := strconv.ParseBool("0")
	fmt.Printf("strconv.ParseBool(\"0\") 的类型为 %T 值为 %t\n", b2, b2) // bool false
	b3, _ := strconv.ParseBool("t")
	fmt.Printf("strconv.ParseBool(\"t\") 的类型为 %T 值为 %t\n", b3, b3) // bool true
	b4, _ := strconv.ParseBool("f")
	fmt.Printf("strconv.ParseBool(\"f\") 的类型为 %T 值为 %t\n", b4, b4) // bool false
	b5, _ := strconv.ParseBool("T")
	fmt.Printf("strconv.ParseBool(\"T\") 的类型为 %T 值为 %t\n", b5, b5) // bool true
	b6, _ := strconv.ParseBool("F")
	fmt.Printf("strconv.ParseBool(\"F\") 的类型为 %T 值为 %t\n", b6, b6) // bool false
	b7, _ := strconv.ParseBool("true")
	fmt.Printf("strconv.ParseBool(\"true\") 的类型为 %T 值为 %t\n", b7, b7) // bool true
	b8, _ := strconv.ParseBool("false")
	fmt.Printf("strconv.ParseBool(\"false\") 的类型为 %T 值为 %t\n", b8, b8) // bool false
	b9, _ := strconv.ParseBool("True")
	fmt.Printf("strconv.ParseBool(\"True\") 的类型为 %T 值为 %t\n", b9, b9) // bool true
	b10, _ := strconv.ParseBool("False")
	fmt.Printf("strconv.ParseBool(\"False\") 的类型为 %T 值为 %t\n", b10, b10) // bool false
	b11, _ := strconv.ParseBool("TRUE")
	fmt.Printf("strconv.ParseBool(\"TRUE\") 的类型为 %T 值为 %t\n", b11, b11) // bool true
	b12, _ := strconv.ParseBool("FALSE")
	fmt.Printf("strconv.ParseBool(\"FALSE\") 的类型为 %T 值为 %t\n", b12, b12) // bool false
	fmt.Println("---------------------------")
}
