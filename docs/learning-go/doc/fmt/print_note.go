package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintNote() {
	fmt.Println("---------------PrintNote()---------------")
	fmt.Println("fmt 包的方法")
	fmt.Println(tab, "func Append(b []byte, a ...any) []byte // 1.19")
	fmt.Println(tab, "func Appendf(b []byte, format string, a ...any) []byte // 1.19")
	fmt.Println(tab, "func Appendln(b []byte, a ...any) []byte // 1.19")
	fmt.Println(tab, "func FormatString(state State, verb rune) string // 1.19")
	fmt.Println("--------------")
	fmt.Println(tab, "func Print(a ...interface{}) (n int, err error) // 相邻的两个参数都不是字符串, 会添加空格")
	fmt.Println(tab, "func Fprint(w io.Writer, a ...interface{}) (n int, err error) // 相邻的两个参数都不是字符串, 会添加空格")
	fmt.Println(tab, "func Sprint(a ...interface{}) string // 相邻的两个参数都不是字符串, 会添加空格")

	fmt.Println(tab, "func Printf(format string, a ...interface{}) (n int, err error)")
	fmt.Println(tab, "func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)")
	fmt.Println(tab, "func Sprintf(format string, a ...interface{}) string")

	fmt.Println(tab, "func Println(a ...interface{}) (n int, err error) // 总是会在两个相邻的参数之间添加空格")
	fmt.Println(tab, "func Fprintln(w io.Writer, a ...interface{}) (n int, err error) // 总是会在两个相邻的参数之间添加空格")
	fmt.Println(tab, "func Sprintln(a ...interface{}) string // 总是会在两个相邻的参数之间添加空格")

	fmt.Println(tab, "func Scan(a ...interface{}) (n int, err error) // 读取空白分隔的值")
	fmt.Println(tab, "func Fscan(r io.Reader, a ...interface{}) (n int, err error) // 读取空白分隔的值")
	fmt.Println(tab, "func Sscan(str string, a ...interface{}) (n int, err error) // 读取空白分隔的值")

	fmt.Println(tab, "func Scanf(format string, a ...interface{}) (n int, err error) // 读取空白分隔的值")
	fmt.Println(tab, "func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) // 读取空白分隔的值")
	fmt.Println(tab, "func Sscanf(str string, format string, a ...interface{}) (n int, err error) // 读取空白分隔的值")

	fmt.Println(tab, "func Scanln(a ...interface{}) (n int, err error)")
	fmt.Println(tab, "func Fscanln(r io.Reader, a ...interface{}) (n int, err error)")
	fmt.Println(tab, "func Sscanln(str string, a ...interface{}) (n int, err error)")

	fmt.Println(tab, "func Errorf(format string, a ...interface{}) error")
	fmt.Println(tab, "type Stringer interface { String() string } // 实现了 Stringer 接口的类型(即有 String 方法), 定义了该类型值的原始显示")
	fmt.Println(tab, tab, "当采用任何接受字符串的 verb(%\\v %\\s %\\q %\\x %\\X)动作格式化一个操作数时, 或者被不使用格式字符串如 Print 函数打印操作数时, 会调用 String 方法来生成输出的文本")
	fmt.Println(tab, "type GoStringer interface { GoString() string } // 实现了 GoStringer 接口的类型(即有 GoString 方法), 定义了该类型值的 go 语法表示")
	fmt.Println(tab, tab, "当采用 verb (%\\#\\v) 格式化一个操作数时, 会调用 GoString 方法来生成输出的文本")
	fmt.Println("--------------")

	fmt.Println("控制台输出字体格式:")
	fmt.Println(tab, "\\033[0;40;30m{value}\\033[0m")
	fmt.Println(tab, "字体格式")
	fmt.Println(tab, tab, "0 关闭所有属性, 1 设置高亮, 4 下划线, 5 闪烁, 7 反显, 8 消隐")
	fmt.Println(tab, "字体颜色: 30 - 37")
	fmt.Println(tab, tab, "30 黑色, 31 红色, 32 绿色,	33 黄色, 34 蓝色, 35 紫色, 36 天蓝, 37 白色")
	fmt.Println(tab, "字体背景色: 40 - 47")
	fmt.Println(tab, tab, "40 黑色, 41 红色, 42 绿色,	43 黄色, 44 蓝色, 45 紫色, 46 天蓝, 47 白色")
	fmt.Println(tab, "其他:")
	fmt.Println(tab, tab, "nA 光标上移 n 行")
	fmt.Println(tab, tab, "nB 光标下移 n 行")
	fmt.Println(tab, tab, "nC 光标右移 n 行")
	fmt.Println(tab, tab, "nD 光标左移 n 行")
	fmt.Println(tab, tab, "2J 清屏")
	fmt.Println(tab, tab, "K 清除从光标到行尾的内容")
	fmt.Println(tab, tab, "y;xH 设置光标位置")
	fmt.Println(tab, tab, "s 保存光标位置")
	fmt.Println(tab, tab, "u 恢复光标位置")
	fmt.Println(tab, tab, "?25l 隐藏光标")
	fmt.Println(tab, tab, "?25h 显示光标")
	fmt.Println("\033[1;32m", "高亮显示绿色文字", "\033[0m")
	fmt.Println("------------------------------")

	fmt.Fprintf(os.Stdout, "%s\n", "Hello world! - unbuffered")
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "Hello world! - buffered")
	buf.Flush()
	fmt.Println("------------------------------")

	fmt.Print("希望结果中包含小数部分则需要浮点数参与运算 10 / 4 = ", 10/4, "\n")
	fmt.Println("------------------------------")
	const name, age = "Kim", 22
	fmt.Print('\n', " m\n") // 不能使用 format 参数字符串，两个相邻的参数不为字符串之间添加空格
	fmt.Print(name, " is ", age, " years old.\n")
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%sis%dyearsold.\n", name, age)    // 第一个参数为 format 参数字符串
	fmt.Println("------------------------------") // 不能使用 format 参数字符串, 相邻参数之间添加空格输出并在输出结束后换行
	str := fmt.Sprint(name, "is", age, "years old.")
	fmt.Println(str)
	val := fmt.Sprintf("%s is %d years old.", name, age)
	fmt.Print(val, "\n")
	fmt.Println("-------------------")
	fmt.Println(name, "is", age, "years old.")
	fmt.Printf("%sis%dyearsold.\n", name, age)
	fmt.Println("hello", "world", "everybody", age)
	fmt.Println("-------------------")
	a := 10
	var ptr *int = &a
	fmt.Println("a 的值为: ", a)
	fmt.Println("a 的地址为: ", ptr)
	fmt.Println("*ptr 指针的值为: ", *ptr)
	fmt.Printf("ptr 的类型为: %T \n", ptr)
	fmt.Printf("hello world\n")
	fmt.Println("--------------------------------")

	b := 1234
	var ptr1 *int
	var pptr1 **int
	ptr1 = &b
	pptr1 = &ptr1 // fmtDemo/main.go:37:10: cannot use &ptr1 (value of type **int) as type *int in assignment
	fmt.Print("变量 b 的地址为: ", &b, "\n")
	fmt.Println("ptr1", ptr1)
	fmt.Println("pptr1", pptr1)
	fmt.Printf("变量 b = %d \n", b)
	fmt.Print("指针变量 ptr 地址为: ", &ptr1, "\n")
	fmt.Print("指向指针的指针变量 pptr1 地址为: ", &pptr1, "\n")
	fmt.Printf("指针变量 *ptr1 = %d\n", *ptr1)
	fmt.Printf("指向指针的指针变量 **pptr1 = %v\n", **pptr1)

	test()
}

func test() {
	fmt.Println("--------func test---------")
	fmt.Print("hello func test", "\n")
}
