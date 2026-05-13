package main

import (
	// "flag"
	"errors"
	"fmt"
)

func funcNote() {
	fmt.Println("-------------funcNote()--------------")
	fmt.Println("go 是编译型语言, 函数调用与编写的顺序是无关")
	fmt.Println("函数分类:")
	fmt.Println(tab, "普通的带有名字的函数")
	fmt.Println(tab, "匿名函数或者 lambda 函数")
	fmt.Println(tab, "方法")
	fmt.Println("形参列表:")
	fmt.Println(tab, "形参没有默认值, 函数调用时需要提供所有的实参")
	fmt.Println(tab, "如果一组形参或者返回值有相同的类型, 则可以在最后一个变量后面写类型 func f(i, j, k int, s, t string) (a, b int)")
	fmt.Println(tab, "如果形参列表不确定参数的个数时, 可以使用可变参数 ...Type, 有且只有一个可变参数保存为一个切片声明在参数列表的最后位置 func f(name string, args ...int){}")
	fmt.Println(tab, tab, "可变参数内部使用一个切片存储接收的参数")
	fmt.Println("返回值列表: ")
	fmt.Println(tab, "返回值只有一个类型声明或者没有返回值时可以省略返回值列表的小括号, 其他情况下不能省略返回值列表的小括号")
	fmt.Println(tab, "命名返回值时, 函数体内的 return 语句后面可以省略返回值, 尽量避免使用空返回值, 并且返回值列表不能省略小括号")
	fmt.Println(tab, "命名多返回值时, 所有返回值必须都声明变量, 否则编译报错 syntax error: mixed named and unnamed function parameters")
	fmt.Println(tab, "命名返回值的参数作用域在函数内, 不会影响函数外的任何变量")
	fmt.Println("递归函数:")
	fmt.Println(tab, "一个问题可以被拆分成多个小问题")
	fmt.Println(tab, "拆分前后的问题除了数据规模不同, 但处理问题的思路是一样的")
	fmt.Println(tab, "不能无限制的调用自身, 需要有退出递归状态的条件")
	fmt.Println("---------------------------")

	fmt.Println("函数类型变量: ")
	fmt.Println(`
  var f func() = fn
  f()
  func fn() {
    fmt.Println("fn was called...")
  }`)
	var f func() = fn
	f()
	fmt.Println("---------------------------")

	fmt.Println("函数调用实参传递")
	fmt.Println(`fnArgs(100, 33, 25, "hello", "world")
  func fnArgs(i, j, k int, s, t string) {
    fmt.Printf("i=%\d j=%\d k=%\d s=%\q t=%\q\n", i, j, k, s, t)
  }`)
	fnArgs(100, 33, 25, "hello", "world")
	fmt.Println("---------------------------")

	fmt.Println("函数调用返回值")
	fmt.Println(fnReturn())
	fmt.Println(`func fnReturn() (a, b int, name string, age int) {
    a = 1
    b = 2
    name = "hello world"
    age = 18
    return
}`)
	fmt.Println("--------------------------")

	fmt.Println("未命名多返回值")
	result1, remainder1, _ := divAndRemainder(5, 2)
	fmt.Printf("result1 = %d remainder1 = %d\n", result1, remainder1) // result1 = 2 remainder1 = 1
	x1, y1, _ := divAndRemainder(7, 3)
	fmt.Printf("x1 = %d y1 = %d\n", x1, y1) // x1 = 2 y1 = 1

	fmt.Println("命名多返回值")
	result2, remainder2, _ := divAndRemainder2(5, 2)
	fmt.Printf("result2 = %d remainder2 = %d\n", result2, remainder2) // result2 = 2 remainder2 = 1
	x2, y2, _ := divAndRemainder2(7, 3)
	fmt.Printf("x2 = %d y2 = %d\n", x2, y2) // x2 = 2 y2 = 1
	fmt.Println("--------------------------")

	// 匿名函数作为 map 的键值
	fmt.Println(`// 匿名函数作为 map 的键值
// var skillParam = flag.String("skill", "", "skill to preform")
var skillParam = new(string)
*skillParam = "fire"
// 声明并初始化 map, key 为 字符串类型, value 为 func() 匿名函数类型返回值为 字符串 类型
var skill = map[string]func() string{
    "fire": func() string { return "fire skill" },
    "run":  func() string { return "run skill" },
    "fly":  func() string { return "fly skill" },
}

if fn, ok := skill[*skillParam]; ok {
    res := fn()
    fmt.Println("res= ", res)
} else {
    fmt.Println("skill not found")
}`)

	// var skillParam = flag.String("skill", "", "skill to preform")
	var skillParam = new(string)
	*skillParam = "fire"
	// 声明并初始化 map, key 为 字符串类型, value 为 func() 匿名函数类型返回值为 字符串 类型
	var skill = map[string]func() string{
		"fire": func() string { return "fire skill" },
		"run":  func() string { return "run skill" },
		"fly":  func() string { return "fly skill" },
	}

	if fn, ok := skill[*skillParam]; ok {
		res := fn()
		fmt.Println("res= ", res)
	} else {
		fmt.Println("skill not found")
	}
	fmt.Println("--------------------------")

	fmt.Println("递归函数")
	for i := 1; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) is: %d\n", i, fibonacci(i))
	}
	fmt.Println("-------------")
	fnum := 7
	fmt.Printf("%d 的阶乘是 %d\n", fnum, factorial(uint64(fnum)))
	fmt.Println("--------------------------")
}

func fn() {
	fmt.Println("fn was called...")
}

func fnArgs(i, j, k int, s, t string) {
	fmt.Printf("i=%d j=%d k=%d s=%q t=%q\n", i, j, k, s, t)
}

func fnReturn() (a, b int, name string, age int) {
	a = 1
	b = 2
	name = "hello world"
	age = 18
	return
}

func fibonacci(n int) int {
	res := 0
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

func factorial(n uint64) (res uint64) {
	if n > 0 {
		res = n * factorial(n-1)
		return
	}
	return 1
}

// 函数返回值
func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// 命名返回值, 所有返回值必须都声明变量, 否则编译报错
// syntax error: mixed named and unnamed function parameters
func divAndRemainder2(numerator, denominator int) (result, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		// 命名返回值 return 后可不跟返回值，但是不能省略 return 关键字
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, nil
}
