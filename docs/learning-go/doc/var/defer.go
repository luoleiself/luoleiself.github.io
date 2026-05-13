package main

import (
	"fmt"
)

func deferNote() {
	fmt.Println("-------------deferNote()-------------")
	fmt.Println("defer 延迟函数, 用于延迟一个函数(包括匿名函数)或者方法在当前的调用栈即将返回时执行")
	fmt.Println(tab, "defer语句只能出现在函数或者方法的内部")
	fmt.Println(tab, "如果有多个 defer 语句, 当前调用栈执行完成时, 会按照\033[1;32m逆序\033[0m的方式执行 defer 函数")
	fmt.Println(tab, "defer 延迟函数的参数在执行 defer 延迟语句时确定, 不是在延迟函数实际调用时确定")
	fmt.Println(tab, "如果存在 return 语句, 则优先于 defer 语句执行")
	fmt.Println("----------------")
	fmt.Println("触发条件")
	fmt.Println("  1. 包含 defer 语句的函数返回时")
	fmt.Println("  2. 包含 defer 语句的函数执行到最后时")
	fmt.Println("  3. 当前 goroutine 发生 panic 时")
	fmt.Println("---------------------------------")

	testPrintAdd()
	fmt.Println("---------------------------------")

	testOrder()
	fmt.Println("---------------------------------")

	closureNote()
	fmt.Println("---------------------------------")

	fmt.Println(`
  func deferReturn() (t int) {
    defer func() {
      t = t * 10
    }()
	
    return 1
  }`)
	fmt.Printf("return 语句先执行返回 1, 然后执行 defer 语句修改 t 的值, 最后输出结果 %d \n", deferReturn())
	fmt.Println("---------------------------------")
}

// 测试
func testPrintAdd() {
	fmt.Println(`
    a, b := 5, 6
    defer printAdd(a, b, true) // 延迟执行函数
    a = 7
    b = 10
    printAdd(a, b, false) // 未延迟执行函数 `)
	fmt.Println("----------------")
	a, b := 5, 6
	defer printAdd(a, b, true) // 延迟执行函数 11
	a = 7
	b = 10
	printAdd(a, b, false) // 未延迟执行函数 17
}

func printAdd(a, b int, t bool) {
	if t {
		fmt.Printf("延迟执行函数 printAdd, 参数 a 为 %d, 参数 b 为 %d, a + b = %d \n", a, b, a+b) // 5 6 11
	} else {
		fmt.Printf("未延迟执行函数 printAdd, 参数 a 为 %d, 参数 b 为 %d, a + b = %d \n", a, b, a+b) // 7 10 17
	}
}

func testOrder() {
	fmt.Println(`
    defer testA() // defer 延迟函数最后逆序执行
    testB()
    defer fmt.Println("defer fmt.Println()")
    defer testC() // defer 延迟函数最后逆序执行
    fmt.Println("testOrder() over...")`)
	fmt.Println("----------------")
	defer testA()                            // defer 延迟函数最后逆序执行 // 5
	testB()                                  // 1
	defer fmt.Println("defer fmt.Println()") // 4
	defer testC()                            // defer 延迟函数最后逆序执行 // 3
	fmt.Println("testOrder() over...")       // 2
}

func testA() {
	fmt.Println("testA...")
}
func testB() {
	fmt.Println("testB...")
}
func testC() {
	fmt.Println("testC...")
}

func closureNote() {
	fmt.Println("-----------closureNote()----------------")
	fmt.Println("闭包: 引用了自由变量的函数, 被引用的自由变量和函数一同存在, 即使已经离开了自由变量的环境也不会被释放或者删除, 在闭包中可以继续使用这个自由变量")
	fmt.Println("---------------------------------")

	acc := Accumulate(0) // 创建一个累加器
	fmt.Println(acc())
	fmt.Println(acc())
	fmt.Println(acc())
	fmt.Println("---------------------------------")

	generator := PlayerGen("Tom") // 创建一个生成器
	name, hp := generator()       // 调用生成器返回闭包中引用的局部变量
	fmt.Printf("name = %s hp = %d\n", name, hp)
	fmt.Println("---------------------------------")
}
func Accumulate(v int) func() int {
	// v 局部变量
	return func() int { // 返回一个函数, 函数内引用局部变量 v
		v++
		return v
	}
}

func PlayerGen(name string) func() (string, int) {
	hp := 250                     // 声明局部变量 hp
	return func() (string, int) { // 返回一个函数, 函数内引用局部变量 hp
		return name, hp
	}
}

func deferReturn() (t int) {
	defer func() {
		t = t * 10
	}()

	return 1
}
