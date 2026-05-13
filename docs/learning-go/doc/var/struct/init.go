package structdoc

import (
	"fmt"
	"unsafe"
)

// 实例化结构体方式 5 种
func Instantiate() {
	fmt.Println("实例化 struct 的方式: 5 种")
	// 结构体中包含 _ 占位符时
	c := Cat{"miao", 18, struct{}{}}
	fmt.Println(c)
	fmt.Println("-----------------")

	fmt.Println("\033[1;32m声明方式1\033[0m: 声明一个结构体类型变量")
	var stu1 Student // 声明一个 Student 类型的变量
	stu1.name = "小刚"
	fmt.Println("var stu1 Student // 声明一个 Student 类型的变量")
	fmt.Println("stu1.name = \"小刚\"")
	// stu1 = /*&*/Student{name: "变量"} // cannot use &(Student literal) (value of *Student) as Student value in assignment
	// fmt.Println("// stu1 = /*&*/Student{name: \"变量\"} // cannot use &(Student literal) (value of *Student) as Student value in assignment")
	// {name:小刚 address: age:0}	40	8
	fmt.Printf("stu1 = %+v 内存占用大小为 %d 内存对齐系数为 %d\n", stu1, unsafe.Sizeof(stu1), unsafe.Alignof(stu1))
	fmt.Println("-----------------")

	var stu2 = Student{} // 实例化 Student 结构体未传值
	stu2.name = "实例化时未传任何值使用字段类型零(默认)值"
	fmt.Println("\033[1;32m声明方式2\033[0m: ")
	fmt.Println(`var stu2 = Student{} // 实例化 Student 结构体未传值`)
	// {name:实例化时未传任何值使用字段类型零(默认)值 address: age:0}	40	8
	fmt.Printf("stu2 = %+v 内存占用大小为 %d 内存对齐系数为 %d\n", stu2, unsafe.Sizeof(stu2), unsafe.Alignof(stu2))
	fmt.Println("-----------------")

	fmt.Println("\033[1;32m声明方式3\033[0m: 实例化时省略结构体字段名")
	var stu3 = Student{"小明", "beijing", 18}
	// var stu3 = Student{name: "小明", address: "beijing", age: 18} // 键值对形式
	fmt.Println(`// var stu3 = Student{"小明", "beijing", 18} // 值列表形式, 传入值的顺序必须和字段声明顺序保持一致`)
	// {name:小明 address:beijing age:18}	40	8
	fmt.Printf("stu3 = %+v 内存占用大小为 %d 内存对齐系数为 %d\n", stu3, unsafe.Sizeof(stu3), unsafe.Alignof(stu3))
	fmt.Println(tab, "省略字段名时需要初始化结构体的所有字段")
	fmt.Println(tab, "初始值的顺序必须和字段在结构体中的声明顺序一致")
	fmt.Println(tab, "键值对和值列表的初始化形式不能混用")
	fmt.Println("-----------------")

	fmt.Println("\033[1;32m声明方式 4 和 5\033[0m")
	fmt.Println("stu4 和 stu5 的类型为结构体指针")
	var stu4 *Student = new(Student) // 声明 Student 类型的结构体指针未初始化
	stu4.name = "小二"
	// stu4 = /*&*/Student{name: "指针类型"} // cannot use (Student literal) (value of type Student) as *Student Value in assignment
	fmt.Println(`var stu4 *Student = new(Student) // 声明一个 Student 类型的结构体指针未初始化`)
	fmt.Println("// stu4 = /*&*/Student{name: \"指针\"} // cannot use (Student literal) (value of type Student) as *Student Value in assignment")
	fmt.Println("stu4.name = \"小二\"")
	// &{name:小二 address: age:0}	8	8
	fmt.Printf("stu4 = %+v 内存占用大小为 %d 内存对齐系数为 %d\n", stu4, unsafe.Sizeof(stu4), unsafe.Alignof(stu4))
	fmt.Println("-----------------")
	var stu5 *Student = &Student{"小三", "beijing", 18} // 声明 Student 类型的结构体指针并初始化
	fmt.Println(`var stu5 *Student = &Student{"小三", "beijing", 18} `)
	// &{name:小三 address:beijing age:18}	8	8
	stu5.name = "小四" // 结构体指针变量操作结构体字段时可以不加变量前的 *, 编译器做了优化
	fmt.Printf("stu5 = %+v 内存占用大小为 %d 内存对齐系数为 %d\n", stu5, unsafe.Sizeof(stu5), unsafe.Alignof(stu5))
	fmt.Printf("stu5 的 Name 为 %v, 地址为 %v, 年龄为 %v\n", stu5.name, stu5.address, stu5.age) // 小四 beijing 18
	fmt.Println("----------------------------------------")
}
