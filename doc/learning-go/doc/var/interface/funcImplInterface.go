package interfacedoc

import (
	"fmt"
	_ "net/http"
	_ "net/url"
	_ "strings"
)

func FuncImplInterfaceNote() {
	fmt.Println("----------funcImplInterfaceNote()-------------")
	fmt.Println("接口型函数: 指的是用函数实现接口, 这样在调用的时候就会非常简便, 这种方式适用于只有一个函数的接口")
	fmt.Println("函数适配器")
	var invoker Invoker = &Struct{}
	invoker.Call("hello world")

	// 将匿名函数转为 FuncCaller 类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from FuncCaller ", v)
	})

	fmt.Println("----------------")
	fmt.Println(`
// 定义一个接口
type Invoker interface {
    Call(interface{})
}

// 定义函数类型别名
type FuncCaller func(interface{})

// 定义函数类型方法, 实现 Invoker 接口
func (f FuncCaller) Call(p interface{}) {
    f(p) // 调用f函数本体
}

// 将匿名函数转为 FuncCaller 类型, 再赋值给接口
var invoker Invoker = FuncCaller(func(v interface{}) {
    fmt.Println("from FuncCaller ", v)
})

invoker.Call("hello FuncCaller")`)
	invoker.Call("hello FuncCaller")
	fmt.Println("----------------------------------")
}

// 定义一个结构体
type Struct struct {
}

// 定义结构体方法
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct ", p)
}

// 定义一个接口
type Invoker interface {
	Call(interface{})
}

// 定义函数类型别名
type FuncCaller func(interface{})

// 定义函数类型方法, 实现 Invoker 接口
func (f FuncCaller) Call(p interface{}) {
	f(p) // 调用f函数本体
}
