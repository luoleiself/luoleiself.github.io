package interfacedoc

import (
	"fmt"
	"strings"
)

func AssertionNote() {
	tab := strings.Repeat(" ", 4)
	fmt.Println("-----------assertionNote()---------------")
	fmt.Println("类型断言(Type Assertion) 是一个使用在接口值上的操作, 用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型")
	fmt.Println(tab, "仅能对静态类型为空接口(interface{})的对象进行断言, 否则会抛出错误")
	fmt.Println(tab, "如果 T 是一个具体类型名, 则类型断言用于判断接口变量 i 绑定的实例类型是否就是具体类型 T")
	fmt.Println(tab, "如果 T 是一个接口类型名, 则类型断言用于判断接口变量 i 绑定的实例类型是否同时实现了接口 T")
	fmt.Println("------------------------------------")
	var i interface{}
	i = 100
	fmt.Printf("i 的类型为 %T 值为 %v\n", i, i)
	b := i
	b = 2
	fmt.Printf("b := i; b = 2; 打印 b=%v \n", b)
	i = "hello world"
	fmt.Printf("i 的类型为 %T 值为 %v\n", i, i)
	i = 3.1415926
	fmt.Printf("i 的类型为 %T 值为 %v\n", i, i)
	i = true
	fmt.Printf("i 的类型为 %T 值为 %v\n", i, i)
	// var s string = i.(string)
	// fmt.Println("s= ", s)
	fmt.Println("------------------------------------")

	computer := Computer{}
	computer.Working(Phone{})
	computer.Working(Camera{})
	fmt.Println("----------------")
	fmt.Println("类型断言示例: 两种使用方式")
	fmt.Println(`
func (c *Computer) Working(usb Usb) {
    res := usb.Start("", 1)
    fmt.Println("res= ", res)
    // usb.Call() // Call 方法只有 Phone 结构体有, 此处使用接口变量接收传入的 Camera 结构体变量时会编译报错, 需要使用类型断言解决
    // 方式一: 直接赋值
    // 如果使用 接口对象.(实际类型) 方式, 可以结合 if else 使用
    if phone, ok := usb.(Phone); ok {
        phone.Call()
    }
    // 方式二: 类型查询
    // 如果使用 接口对象.(type) 方式则必须结合 switch 使用
    switch instance := usb.(type) {
    case Phone:
        instance.Call()
    }
    usb.Stop()
}`)
	fmt.Println("------------------------------------")

}

// 声明接口
type Usb interface {
	Start(string, int) bool
	Stop()
}

// 声明结构体
type Phone struct{}

// 声明结构体的方法
func (p *Phone) Call() {
	fmt.Println("手机打电话的功能....独一份")
}

// 实现 Usb 接口定义的方法
func (p Phone) Start(s string, i int) bool {
	fmt.Println("手机开始工作了...", s, i)
	return true
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作了...")
}

// 声明结构体
type Camera struct{}

// 实现 Usb 接口定义的方法
func (c Camera) Start(s string, i int) bool {
	fmt.Println("照相机开始工作了...", s, i)
	return false
}
func (c Camera) Stop() {
	fmt.Println("照相机停止工作了....")
}

// 声明结构体
type Computer struct{}

// 声明结构体的方法
func (c *Computer) Working(usb Usb) {
	res := usb.Start("", 1)
	fmt.Println("res= ", res)
	// usb.Call() // Call 方法只有 Phone 结构体有, 此处使用接口变量接收传入的 Camera 结构体变量时会编译报错, 需要使用类型断言解决
	// 方式一: 直接赋值
	// 如果使用 接口对象.(实际类型) 方式, 可以结合 if else 使用
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	// 方式二: 类型查询
	// 如果使用 接口对象.(type) 方式则必须结合 switch 使用
	switch instance := usb.(type) {
	case Phone:
		instance.Call()
	}
	usb.Stop()
}
