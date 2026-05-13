package interfacedoc

import (
	"fmt"
)

type AInterface interface {
	testA(string) // 重复定义的方法，会报错
	testB()
}

// type BInterface interface {
// 	testA() // 重复定义的方法，会报错
// 	testC()
// }
// type CInterface interface {
// 	AInterface // duplicate method testA
// 	BInterface
// }
// type CInterfaceImpl struct{}

//	func (c *CInterfaceImpl) testA(name string) {
//		fmt.Println("CInterfaceImpl testA func...", name)
//	}
//
//	func (c *CInterfaceImpl) testB() {
//		fmt.Println("CInterfaceImpl testB func...")
//	}
//
//	func (c *CInterfaceImpl) testC() {
//		fmt.Println("CInterfaceImpl testC func...")
//	}

type psay interface {
	SayHello() string
}

type Person struct {
	name string
	age  uint8
}

func (p *Person) SayHello() string {
	return fmt.Sprintf("姓名: %s, 年龄: %d", p.name, p.age)
}

func NewPerson1() psay {
	return &Person{"zhangsan", 19}
}

func NewPerson2() Person {
	return Person{"lisi", 20}
}

func Readme() {
	fmt.Println("------------readme()----------------")
	fmt.Print("接口: 是一组对其他类型行为的抽象和概括的方法的集合, 接口是引用类型, 接口有两个属性, Type, Value, 默认值都为 nil. ")
	fmt.Println(tab, "只有当 Type 和 Value 都为 nil 时, 接口才为 nil")
	var ai AInterface
	fmt.Printf("\tvar ai AInterface 的类型为 %T 值为 %v\n", ai, ai)
	fmt.Println(tab, "如果一个函数接受接口类型作为参数, 那么实际上它可以传入该接口的任意一个实现类的对象作为参数")
	fmt.Println(tab, "定义一个接口类型变量, 实际上可以赋值给任意一个实现了该接口的对象")
	fmt.Println(tab, "接口继承多个接口, 如果多个接口中存在重复定义的方法时, 会报方法重复定义的错")
	fmt.Println(tab, "只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口, 那样只会增加不必要的抽象, 导致不必要的运行时损耗")
	fmt.Println(tab, "保存有类型不同的值的空接口进行比较时, Go语言会优先比较值的类型。因此类型不同, 比较结果也是不相同的")
	var a interface{} = 100                                                                                                  // a保存整型
	var b interface{} = "hi"                                                                                                 // b保存字符串
	fmt.Println(tab, tab, "var a interface{} = 100\n", tab, tab, "var b interface{} = \"hi\"\n", tab, tab, "a == b", a == b) // 两个空接口不相等
	fmt.Println(tab, "不能比较空接口中的动态值, 如 slice, map 等内容可以变化的无法进行比较")
	// var c interface{} = []int{10}
	// var d interface{} = []int{10}
	// fmt.Println("var c interface{} = []int{10}\nvar d interface{} = []int{10}\nc == d", c == d) // comparing uncomparable type []int
	fmt.Println("-----------------")

	fmt.Println("指向具体类型的指针可以实现一个接口, 但是指向接口的指针永远不可能实现该接口")
	fmt.Println("-----------------")

	fmt.Println("声明接口: 结构体实现某个接口，必须实现了接口中所有定义的方法")
	fmt.Println(`
  type 接口标识符 interface {
    方法名(参数列表) 返回值列表
    ...
  }`)
	fmt.Println("------------------------------------")

	// var ci = CInterfaceImpl{}
	// ci.testA("hello world")

	fmt.Println("***1*** 空接口可以承载任意值,但不代表任意类型就可以承接空接口类型的值")
	fmt.Println("var i interface{} = 100; var j int = i; // 不能够赋值, 需要类型断言. 但是可以使用 j := i 方式赋值")
	fmt.Println("***2*** 当空接口承载数组和切片后, 该对象无法再进行切片")
	fmt.Println(`
    sli := []int{1, 2, 3, 4, 5}
    var sliI interface{}
    sliI = sli

    sliI[1:2] // 会报错 invalid operation, 不能对接口类型进行切片操作
    sliI.([]int)[1:2] // 使用类型断言
	`)
	fmt.Println("-----------------")
	sli := []int{1, 2, 3, 4, 5}
	var sliI interface{}
	sliI = sli
	fmt.Println("sliI.([]int)[1]=", sliI.([]int)[1])
	fmt.Println("sliI.([]int)[1:2] 使用类型断言 ", sliI.([]int)[1:2])
	fmt.Println("------------------------------------")

	fmt.Println("如果结构体实现的方法的接受者变量为指针类型，声明使用接口类型变量时需要传递实例的地址")
	fmt.Println(`func (c *Computer) Working(usb Usb) {
} 
var u Usb = &Computer`)
	fmt.Println("------------------------------------")

	p1 := NewPerson1()
	fmt.Printf("p1 的类型为 %T 地址为 %p 值为 %v\n", p1, &p1, p1)
	p2 := NewPerson2()
	fmt.Printf("p2 的类型为 %T 地址为 %p 值为 %v\n", p2, &p2, p2)
	fmt.Println("-----------------")
	fmt.Println("如果 NewPerson 方法内 return 值的类型和此方法声明的返回值类型不一致则报错")
	fmt.Println(tab, "cannot use &Person{…} (value of type *Person) as Person value in return statement")
	fmt.Println(tab, "cannot use Person{…} (value of type Person) as *Person value in return statement")
	fmt.Println("如果 NewPerson 方法声明的返回值类型为接口类型, 接收者类型为指针类型, 则 NewPerson 方法的返回值只能为指针类型, 否则报错")
	fmt.Println(tab, "cannot use Person{…} (value of type Person) as psay value in return statement: Person does not implement psay (method sayHello has pointer receiver)")
	fmt.Println("如果 NewPerson 方法声明的返回值类型为接口类型, 接收者类型为值类型, 则 NewPerson 方法的返回值可以为值类型或者指针类型")
	fmt.Println("------------------------------------")
}
