package main

import (
	"fmt"
)

func calc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 函数类型别名
type myFuncType func(int, int) int

func calcFunc(n1, n2 int) int {
	return n1 + n2
}
func getSum(myFunc myFuncType, num1, num2 int) int {
	return myFunc(num1, num2)
}

func typeAliasNote() {
	fmt.Println("--------------typeAliasNote()---------------")
	fmt.Println("类型别名:  type TypeAlias = Type")
	fmt.Println(tab, "go 1.9 版本添加的新功能, 主要用于代码升级、工程重构、迁移中类型的兼容性问题")
	fmt.Println(tab, "类型别名只是 Type 的别名, 本质上与 Type 是同一个类型")
	fmt.Println(tab, "类型别名只会在代码中存在, 编译完成时, 不会有类型别名")
	fmt.Println(tab, "非本地类型不能定义新方法, 定义的类型别名不属于当前声明的包时, 无法添加新方法")
	fmt.Println(tab, tab, "package main type MyDuration = time.Duration func (m MyDuration) ES(s string) { } // cannot define new methods on non-local type time.Duration")
	fmt.Println(tab, "1.9 之前定义别名")
	fmt.Println(tab, tab, "type byte uint8")
	fmt.Println(tab, tab, "type rune int32")
	fmt.Println(tab, "1.9 之后定义别名")
	fmt.Println(tab, tab, "type byte = uint8")
	fmt.Println(tab, tab, "type rune = int32")
	fmt.Println("类型定义:")
	fmt.Println(tab, "依据基本数据类型声明一个新的数据类型")
	fmt.Println(tab, "使用类型定义的变量无法和原类型的变量进行运算(比较、赋值、算术), 需使用显示(强制)类型转换")
	fmt.Println(tab, "type NewType Type")
	fmt.Println("-----------------------------")

	type NewInt int  // 定义一个新的类型 NewInt, 新类型具备 int 类型的特性,
	type MyInt = int // 定义类型别名, 使用时同 int 一致
	var num1 NewInt = 100
	var num2 MyInt = 120
	fmt.Println(`type NewInt int  // 定义一个新的类型 NewInt, 新类型具备 int 类型的特性,
type MyInt = int // 定义类型别名, 使用时同 int 一致
var num1 NewInt = 100
var num2 MyInt = 120`)
	fmt.Printf("num1 的类型为 %T 值为 %v\n", num1, num1) // main.NewInt 100
	fmt.Printf("num2 的类型为 %T 值为 %v\n", num2, num2) // int 120
	fmt.Println("-----------------------------")

	fmt.Println("函数的返回值列表如果声明返回值变量, 函数内部 return 后面则可省略返回值")
	fmt.Println(`
func calc(n1, n2 int)(sum int, sub int){
	sum = n1 + n2
	sub = n1 - n2
	return 
}`)
	res1, res2 := calc(1, 2)
	fmt.Printf("res1, res2 := calc(1, 2); res1= %d res2= %d\n", res1, res2) // 3 -1
	fmt.Println("-----------------------------")

	fmt.Println("声明函数类型别名 type myFuncType func(int, int) (int, int)")
	myFunc := calc
	fmt.Printf("myFunc := calc 的类型为 %T \n", myFunc) // func(int, int) (int, int)
	fmt.Println("-----------------------------")

	result := getSum(calcFunc, 30, 20)
	fmt.Println(`// 函数类型别名
type myFuncType func(int, int) int // type 定义一个函数类型别名，包含参数类型，返回值类型

func calcFunc(n1, n2 int) int {
  return n1 + n2
}
func getSum(myFunc myFuncType, num1, num2 int) int { // 使用定义的函数类型别名
  return myFunc(num1, num2)
}
result := getSum(calcFunc, 30, 20)`)
	fmt.Println("")
	fmt.Printf("result= %d\n", result) // 50
	fmt.Println("-----------------------------")

	fmt.Println("示例")
	fmt.Println(`// 参考 bufio 包中的 Scanner 结构体的 split 方法, 接收 SplitFunc 类型的参数
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func ScanBytes(data []bytes, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	return 1, data[0:1], nil
}

func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return 1, []byte(string('h')), nil
}

// 此处的 splitFunc 类型标识参数为函数包含有三个参数,并且有三个返回值,可以理解为此处作为函数别名, 避免过长的参数类型定义
func Test(split SplitFunc, name string, age uint8) { 
	res, token, err := split([]byte(string('h')), false) // 此处接收符合 SplitFunc 类型的函数的三个返回值
	fmt.Println(res, token, err)
	fmt.Println(name, age)
}	

func main(){
	Test(ScanRunes, "hello world", 18) // 调用 Test 函数 将 ScanRunes 传入
}`)
	fmt.Println("-----------------------------")
}
