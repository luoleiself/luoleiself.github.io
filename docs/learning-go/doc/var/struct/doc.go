package structdoc

import (
	"fmt"
)

type Age uint8

func (a *Age) Add(i uint8) {
	(*a) += Age(i)
}
func (a *Age) String() {
	fmt.Println("age is ", *a)
}

/*
	特性			 具名结构体字段				匿名结构体字段（嵌入）

定义方式			显式命名字段			直接嵌入结构体或接口
访问方式		 通过字段名访问			可以直接访问嵌入结构体的字段和方法
继承性				  不具备				 自动继承嵌入结构体的所有字段和方法
代码简洁性		  更加冗长			  更加简洁
字段冲突处理	  无冲突	        如果有同名字段或方法，需要显式指明
多态性	         无	           支持，特别是嵌入接口时
可读性	      更容易理解	     对于简单的嵌入关系，可读性好；复杂时可能难懂
*/

func Readme() {
	fmt.Println("Go 没有构造函数, 虚函数(运行时确定, 多态), 析构函数, 抽象结构体, 方法重载的概念")
	fmt.Println("声明 结构体(struct) 类型, 结构体是值类型, 结构体的定义只是一种内存布局的描述, 只有当结构体实例化时, 才会真正的分配内存")
	fmt.Println("结构体是由零个或多个任意类型的值聚合成的实体, 每个值都称为结构体的成员")
	fmt.Println("结构体整体占用内存字节数需要时结构体类型成员对齐边界(成员中类型占用内存最大的倍数)")
	fmt.Println("实例化是根据结构体的定义的格式创建一份和格式一致的内存区域")
	fmt.Println(tab, "类型名: 标识自定义结构体的名称, 同一个包内不能重复")
	fmt.Println(tab, "struct{}: 标识结构体, 将 struct{} 结构体定义为类型名的类型")
	fmt.Println(tab, "字段名: 标识结构体中的字段必须唯一")
	fmt.Println(tab, "字段类型: 标识结构体各个字段的类型")
	fmt.Println("使用 new 或者 & 构造的实例的类型是类型的指针, 结构体指针变量操作结构体字段时可以不加变量前的 *, 编译器做了优化")
	fmt.Println("匿名结构体不需要使用 type 关键字声明就可以直接使用, 声明和初始化需要同时进行")
	fmt.Println("-----------------")
	fmt.Println("检查结构体是否为空")
	fmt.Println(tab, "1. 使用零值字面量检查 s1 == Student{}")
	fmt.Println(tab, "2. 对于具有不可比较字段, 使用 reflect.DeepEqual(s1, Student{}) 进行检查")
	fmt.Println(tab, "3. 使用反射包 Value 的 IsZero() 检查 reflect.ValueOf(s1).IsZero()")
	fmt.Println("-----------------")

	fmt.Println("匿名结构体，匿名字段，方法继承，方法重写")
	fmt.Println("----------------------------------------")

	fmt.Println("结构体是值类型")
	var s = Student{name: "小明", age: 18, address: "北京市"}
	fmt.Printf("学生 s 的值为 %+v \n", s) // {name:小明 address:北京市 age:18}
	var s1 = s
	s1.name = "小蓝"
	fmt.Print("var s1 = s\ns1.name = \"小蓝\"\n")
	fmt.Printf("s1 的值为 %+v\n", s1) // {name:小蓝 address:北京市 age:18}
	fmt.Printf("s 的值为 %+v \n", s)  // {name:小明 address:北京市 age:18}
	fmt.Println("----------------------------------------")

	fmt.Println("方法: 一种作用于特定类型变量的函数, 这种特定类型变量称为接受者, 类似其他语言的 self 或者 this, 接受者强调了方法具有作用对象")
	fmt.Println(tab, "接受者不管是使用值接收者, 还是指针接收者, 一定要搞清楚类型的本质：对类型进行操作的时候，是要改变当前值，还是要创建一个新值进行返回")
	fmt.Println(`
  func (接受者变量 接受者类型) 方法名(参数列表) (返回值列表) {
  }`)
	fmt.Println("---------------")

	fmt.Println("方法表达式: 方法赋值给变量")
	fmt.Println(`
  var age = Age(25)
  var add = age.Add
  var printAge = age.String
  printAge()
  add(uint8(5))
  printAge()`)
	var age = Age(25)
	var add = age.Add
	var printAge = age.String
	printAge()
	add(uint8(5))
	printAge()
	fmt.Println("----------------------------------------")

	fmt.Println("方法继承: 如果匿名字段结构体实现了一个方法，那么包含这个匿名字段的结构体也能调用该匿名字段中的方法")

	fmt.Println("----------------------------------------")

	fmt.Println("方法重写: 一个包含了匿名字段的结构体也实现了该匿名字段实现的方法, 存在继承关系时，方法调用按照就近原则")

	fmt.Println("----------------------------------------")

	fmt.Println("匿名结构体: 没有名字的结构体，无须通过 type 定义可以直接使用， 创建匿名结构体时同时创建对象")
	cat := struct {
		name, color string
		age         uint8
	}{"小花", "黄色", 18}
	fmt.Println(`
  cat := struct {
    name, color string
    age         uint8
  }{"小花", "黄色", 18}`)
	fmt.Printf("cat 的值为 %+v \n", cat)
	fmt.Println("----------------------------------------")

	fmt.Println("匿名字段: 结构体中包含一个没有字段名的类型，同一个类型只能有一个匿名字段，通过匿名结构体字段可以模拟继承关系")
	fmt.Println(`
  type Student struct {
    Person // 匿名结构体字段模拟继承关系
    name address string
    age unit8
    float64 // 默认类型匿名字段
  }`)
	fmt.Println("----------------------------------------")
}
