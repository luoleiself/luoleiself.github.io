package main

import (
	"fmt"
	"net/http"
)

type MyT interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | string | bool | float32 | float64 | *http.Client | [10]int | []int | map[int]int | struct{}
}

// type Numbers[T int32 | int64 | float32 | float64] []T

type Number interface {
	int32 | int64 | float32 | float64
}

// 等价于上面定义的 Numbers 类型
type Numbers[T Number] []T

type Float interface {
	float32 | float64
}
type Float32 interface {
	Float
	float64
}

type KV[K int32 | float32, V int8 | bool] map[K]V

type User[T int32 | string, TS []T | []string] struct {
	Id     T
	Emails TS
}

type MyStruct[S int | string, P map[S]string] struct {
	Name    string
	Content S
	Job     P
}

func Min[T int | int32 | int64 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// type Foo[T any] []T   // true
// type Bar[T any] T     // false 不能单独使用泛型形参作为泛型类型
// type Baz[T any] struct { T } // false 不能使用泛型形参作为匿名字段

// Equal 检查两个可比较类型的值是否相等
func Equal[T comparable](a, b T) bool {
	return a == b
}

func genericTypeNote() {
	fmt.Println("--------------genericTypeNote()---------------")
	fmt.Println("泛型: 允许在强类型语言中编写代码时使用一些以后才指定的类型, 在实例化时作为参数指明这些类型")
	fmt.Println("  1. 函数和类型(包含接口)支持类型参数(type parameter), 方法暂不支持")
	fmt.Println("  2. 支持类型推导, 大多数情况下, 调用泛型函数时可省略类型实参(type argument)")
	fmt.Println("  3. 通常以单个大写字母命名类型参数")
	fmt.Println("  4. 类型参数必须要有约束(constraints)")
	fmt.Println("  5. type parameter as RHS")
	fmt.Println("   .\\generic.go:7:50: cannot use a type parameter as RHS in type declaration")
	fmt.Println("   报错: 还不支持单独使用类型形参作为泛型类型, 需要结合struct、slice和map等类型来使用")
	fmt.Println("  6. 泛型不支持匿名结构体和匿名字段")
	fmt.Println("----------------")

	fmt.Println("泛型不支持直接使用类型断言, 需要通过转换到 interface{} 或者使用 reflect 来实现")
	fmt.Println(tab, "func (s *Stack[T]) Push(elem T){ var a = interface{} a = elem switch a.(type)}")
	fmt.Println("==泛型代码只能使用其类型参数已知可实现的操作==")
	fmt.Println("任何定义的形参, 在使用时都需要有按顺序一一对应的实参")
	fmt.Println("为什么泛型使用 [] 而不是 <>")
	fmt.Println(tab, "当解析一个函数的代码块如 v := F<T> 时, 编译器无法理解 < 是一个操作符还是一个泛型的实例化, 为使 Go 的语法解析保持足够的简单")
	fmt.Println("----------------")

	fmt.Println("将接口类型定义为类型集合(type set), 包括没有方法的接口类型, 泛型之前称为方法集合(method set), 多个类型都实现了接口, 因此可以把此接口看成多个类型的集合")
	fmt.Println("    类型集合为了简化泛型约束的使用, 提升阅读性, 同时增加了复用功能, 通过接口定义的方式使用")
	fmt.Println("       type Constraint1 interface { Type1 | ~Type2 | ... }")
	fmt.Println("接口类型的变量可以存储接口类型集中任何类型的值, 可以分成\033[1;33m基本接口\033[0m和\033[1;33m一般接口\033[0m")
	fmt.Println("    1. 基本接口: 接口定义里面只有方法没有类型, go 1.18 之前的写法")
	fmt.Println("       1. 可以定义变量, 例如 var err error")
	fmt.Println("       2. 可以作为类型约束, 例如 type ReadOrWriters[T io.Reader | io.Writer] []T")
	fmt.Println("    2. 一般接口: 接口定义中只要包含类型约束(无论是否包含方法)")
	fmt.Println("       1. 不能用来定义变量(限制一般接口只能用在泛型内), 只能用来定义类型约束")
	fmt.Println("函数参数类型推导: 从函数参数的类型推导出类型参数的情形")
	fmt.Println("    函数参数类型推导仅适用于函数参数中使用的类型参数, 不适用于仅用于函数结果或仅在函数体中使用的类型参数")
	fmt.Println("    例如 func testT[T any]() T 不适用函数参数类型推导")
	fmt.Println("----------------")

	fmt.Println(`type Numbers[T int32 | int64 | float32 | float64] []T`)
	fmt.Println(tab, "Numbers[T] 表示泛型类型")
	fmt.Println(tab, "T 表示泛型形参")
	fmt.Println(tab, "T int32 | int64 | float32 | float64 表示类型形参列表")
	fmt.Println(tab, "int32 | int64 | float32 | float64 表示类型约束")
	fmt.Println(tab, "[]T 表示定义类型")
	fmt.Println(tab, "Numbers[int32] 表示类型实例化, int32 表示类型实参, 必须是类型约束中定义的类型")
	fmt.Println("----------------")
	fmt.Println("多个泛型形参")
	fmt.Println(`type KV[K int32 | float32, V int8 | bool] map[K]V`)
	fmt.Println(tab, "KV[K,V] 表示泛型类型")
	fmt.Println(tab, "K V 表示泛型形参")
	fmt.Println(tab, "K int32 | float32, V int8 | bool 表示类型形参列表")
	fmt.Println(tab, "int32 | float32 表示 K 的类型约束, int | bool 表示 V 的类型约束")
	fmt.Println(tab, "KV[int32, bool] 表示类型实例化, int32 表示 K 的类型实参, bool 表示 V 的类型实参")
	fmt.Println("----------------")
	fmt.Println("嵌套类型形参")
	fmt.Println(`
  type User[T int32 | string, TS []T | []string] struct {
    Id     T
    Emails TS
  }`)
	fmt.Println(tab, "T 和 TS 表示类型形参")
	fmt.Println(tab, "int32 | string 表示 T 的类型约束, []T | []string 表示 TS 的类型约束")
	fmt.Println("----------------")

	fmt.Println("使用接口预定义类型约束: 可以避免类型约束过长")
	fmt.Println(`type Number interface {
    int32 | int64 | float32 | float64
  }`)
	fmt.Println("内置接口: any, comparable, ordered")
	fmt.Println(tab, "any 其实就是 interface{} 的别名, 语义上会比 interface{} 更加清晰")
	fmt.Println(tab, "comparable 约束的类型集是所有可比较类型的集合, 允许使用 == 和 != 操作, 但不支持比较大小 >, <")
	fmt.Println(tab, " 它仅能作为类型参数的约束使用, 而不能作为变量的类型")
	fmt.Println(tab, "ordered 约束的类型集是所有可比较类型的集合, 支持比较大小 >, <")
	fmt.Println("近似约束: ~ ")
	fmt.Println(" 在实例化泛型时, 不仅可以直接使用对应的实参类型, 如果实参的底层类型在类型约束中, 也可以使用")
	fmt.Println(`
  type MyInt int
  type Ints[T int | int32] []T

  // 修改: 如果底层类型 match, 就可以正常进行泛型的实例化
  // type Ints[T ~int | ~int32] []T
  
  a := Ints[int]{10, 20} // ok
  b := Ints[MyInt]{10, 20} // fail
  // MyInt does not implement int|int32 (possibly missing ~ for int in constraint int|int32)
  `)
	fmt.Println("interface 集合操作")
	fmt.Println(tab, "空集: 一个空的类型约束, 没有实际意义, 编译器不会阻止这样使用")
	fmt.Println(tab, " type Null interface {\n\tfloat32\n\tint32\n }")
	fmt.Println(tab, "并集: 使用竖线分割的多个类型")
	fmt.Println(tab, "交集: 将类型分别写到多行, Float32 是 Float 和 float64 的交集, Float 是 float32 | float64 的并集")
	fmt.Println(tab, " Float32 只指定了 float32 一个泛型约束")
	fmt.Println(`
  // 并集
  type Float interface{
    float32 | float64
  }
  // 交集
  type Float32 interface {
    Float
    float64
  }`)
	fmt.Println("----------------")

	fmt.Println("泛型函数")
	fmt.Println(`
 func Min[T int | int32 | int64 | float32 | float64](a, b T) T {
  if a < b {
    return a
  }
  return b
 }`)
	fmt.Println("支持自动类型推导")
	fmt.Printf("Min(10, 20) 的值为 %v // 不再需要写成 Min[int32](10, 20)\n", Min(10, 20))               // 10
	fmt.Printf("Min(10.1, 20.2) 的值为 %v // 不再需要写成 Min[float64](10.1, 20.2)\n", Min(10.1, 20.2)) // 10.1
	fmt.Println("----------------")

	// var a Numbers[int32] = []int32{1, 2, 3, 4}
	var a = Numbers[int32]{1, 2, 3, 4}
	// var b Numbers[float64] = []float64{23.23, 33.45}
	var b = Numbers[float64]{23.23, 33.45}
	fmt.Printf("Numbers[int32]{1, 2, 3, 4} 的类型为 %T 值为 %v\n", a, a)     //  main.Numbers[int32]  [1 2 3 4]
	fmt.Printf("Numbers[float64]{23.23, 33.45} 的类型为 %T 值为 %v\n", b, b) // main.Numbers[float64]  [23.23 33.45]
	fmt.Println("----------------")

	// var c KV[int32, bool] = map[int32]bool{10: true}
	var c = KV[int32, bool]{10: true}
	fmt.Printf("KV[int32, bool]{10: true} 的类型为 %T 值为 %v\n", c, c) //  main.KV[int32,bool]  map[10:true]
	fmt.Println("----------------")

	// var d User[int32, []string]
	// d.Id = 10
	// d.Emails = []string{"123@qq.com", "456@sina.com"}
	var d = User[int32, []string]{Id: 10, Emails: []string{"123@qq.com", "456@sina.com"}}
	//  main.User[int32,[]string]  {10 [123@qq.com 456@sina.com]}
	fmt.Printf("User[int32, []string]{Id:10,Emails: []string{\"123@qq.com\", \"456@sina.com\"}} \n 的类型为 %T 值为 %v\n", d, d)
	fmt.Println("----------------")

	// var myStruct1 MyStruct[int, map[int]string]
	// myStruct1.Name = "1"
	// myStruct1.Content = 1
	// myStruct1.Job = map[int]string{1: "job"}
	var myStruct1 = MyStruct[int, map[int]string]{Name: "1", Content: 1, Job: map[int]string{1: "job"}}
	// myStruct1 {Name:1 Content:1 Job:map[1:job]}
	fmt.Printf("myStruct1 %+v\n", myStruct1)
	fmt.Println("--------")
	// var myStruct2 MyStruct[string, map[string]string]
	// myStruct2.Name = "2"
	// myStruct2.Content = "2"
	// myStruct2.Job = map[string]string{"2": "job"}
	var myStruct2 = MyStruct[string, map[string]string]{Name: "2", Content: "2", Job: map[string]string{"2": "job"}}
	// myStruct2 main.MyStruct[string,map[string]string]{Name:"2", Content:"2", Job:map[string]string{"2":"job"}}
	fmt.Printf("myStruct2 %#v\n", myStruct2)
	fmt.Println("----------------")

	/*
		// 泛型不支持匿名结构体
		generic.go:188:22: expected '{', found '['
		generic.go:194:3: expected '}', found 'EOF'
		exit status 2
	*/
	// var struct1 = struct[T int|float64]{
	// 	Name string
	// 	Age int
	// 	Weight T
	// }[int]{"xiaoming", 18, 50}
	// fmt.Println("struct1 ", struct1)
	fmt.Println("----------------")

	basicInterface()

	generalInterface()
}

type myInterface[T int | string] interface {
	WriteOne(data T) T
	ReadOne() T
}

type noteString struct{}

func (n noteString) String() string {
	return "noteString"
}
func (n noteString) WriteOne(one string) string {
	fmt.Print(n.String(), " ---- ")
	return one
}
func (n noteString) ReadOne() string {
	fmt.Print(n.String(), " ---- ")
	return " "
}

type noteInt struct{}

func (n noteInt) String() string {
	return "noteInt"
}
func (n noteInt) WriteOne(one int) int {
	fmt.Print(n.String(), " ---- ")
	return one
}
func (n noteInt) ReadOne() int {
	fmt.Print(n.String(), " ---- ")
	return 0
}

func basicInterface() {
	fmt.Println("基本接口: 泛型形参列表位置在接口定义而非方法定义, 定义空结构体实现接口的所有方法")

	var one myInterface[string] = noteString{}
	fmt.Println(one.WriteOne("hello")) // noteString ---- hello
	fmt.Println(one.ReadOne())         // noteString ----
	fmt.Println("--------")

	var two myInterface[int] = noteInt{}
	fmt.Println(two.WriteOne(10)) // noteInt ---- 10
	fmt.Println(two.ReadOne())    // noteInt ---- 0
	// fmt.Println(two.WriteOne("hello")) //  cannot use "hello" (untyped string constant) as int value in argument to two.WriteOne
	fmt.Println("two.WriteOne(\"hello\")")
	fmt.Println("cannot use \"hello\" (untyped string constant) as int value in argument to two.WriteOne")
	fmt.Println("不能在 two.WriteOne 的参数中使用 \"hello\"(非类型化字符串常量)作为 int 值")
	fmt.Println("----------------")
}

// type myInterface2[T int | string] interface {
// 	int | string

// 	WriteOne2(data T) T
// 	ReadOne2() T
// }

// type note2 int

// func (n note2) WriteOne2(one int) int {
// 	return one
// }
// func (n note2) ReadOne2() int {
// 	return 0
// }

func generalInterface() {
	fmt.Println("一般接口: 只能用来定义类型约束")

	// var one myInterface2[int] // cannot use type myInterface2[int] outside a type constraint: interface contains type constraints
	// fmt.Println(one.WriteOne2(10))
	// fmt.Println(one.ReadOne2())
	fmt.Println("var one myInterface2[int]")
	fmt.Println("cannot use type myInterface2[int] outside a type constraint: interface contains type constraints")
	fmt.Println("不能在类型约束之外使用类型 myInterface2[int]: 接口包含类型约束")
	fmt.Println("----------------")
}

// 泛型集合
type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: map[T]struct{}{}}
}
func (s *Set[T]) Add(value T) {
	s.items[value] = struct{}{}
}
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.items[value]
	return exists
}
func (s *Set[T]) Print() {
	for v := range s.items {
		fmt.Println(v)
	}
}
