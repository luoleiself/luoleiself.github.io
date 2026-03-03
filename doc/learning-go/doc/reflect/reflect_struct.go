package main

import (
	"fmt"
	"reflect"
	"strings"
)

func reflectStruct() {
	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，Name() 会返回 ""
	fmt.Println("--------------reflectStruct()---------------")
	fmt.Println("func VisibleFields(t Type) []StructField // 获取 t 中的所有可见字段, 返回的字段包括匿名结构成员内的字段和未导出的字段")
	fmt.Println(tab, "必须是 struct 类型, 如果字段可通过 FieldByName 调用直接访问, 则该字段被定义为可见, go 1.17 支持")
	fmt.Println("func StructOf(fields []StructField) Type // 动态创建结构体")
	fmt.Println("  // 使用 reflect 构建结构体字段: 传入 reflect.StructField 类型切片 到 reflect.StructOf() 创建 reflect.Type 类型的结构体")
	fmt.Println(tab, "func (t Type) FieldByIndex(index []int) StructField")
	fmt.Println(tab, "   // 返回索引序列指定的嵌套字段的类型, 等价于用索引中的每个值链式调用本方法")
	fmt.Println(tab, "func (t Type) FieldByName(name string) (StructField, bool)")
	fmt.Println(tab, "   // 返回改类型名为 name 的字段(会查找匿名字段及其子字段), 布尔值表示是否找到, 如非结构体则会 panic")
	fmt.Println(tab, "func (t Type) FieldByNameFunc(match func(string) bool) (StructField, bool)")
	fmt.Println(tab, "   // 返回改类型第一个字段名满足函数 match 的字段, 布尔值表示是否找到, 如非结构体则会 panic")
	fmt.Println(tab, "func (t Type) Field(i int) StructField // 返回 struct 类型的第 i 个字段的类型, 如非结构体或 i 不在[0, NumField()] 内将会 panic")
	fmt.Println(tab, tab, "func (f StructField) IsExported() bool // 判断字段是否导出")
	fmt.Println(tab, "func (t Type) Method(int) Method	// 获取指定的方法")
	fmt.Println(tab, tab, "func (m Method) IsExported() bool // 判断方法是否导出")
	fmt.Println(tab, "func (t Type) MethodByName(string) (Method, bool) // 根据方法名返回改类型方法集中的方法, 布尔值表示是否找到")
	fmt.Println("---------------")
	fmt.Println("type StructTag string")
	fmt.Println("func (tag StructTag) Get(key string) string")
	fmt.Println(tab, "// 返回与标签字符串中的键关联的值, 如果未找到则返回空字符串, 如果标签不具有常规格式, 则 Get 返回的值是未指定的, 要确定标记是否显式设置为空字符串, 使用 Lookup")
	fmt.Println("func (tag StructTag) Lookup(key string) (value string, ok bool)")
	fmt.Println(tab, "// 作用同 Get 方法, ok 表示该返回值是否已在标记字符串中显式设置, 如果标签不具有常规格式, 则 Lookup 返回的值是未指定的")
	fmt.Println("---------------")
	fmt.Println("如果 func TypeOf(i interface{}) 或者 func ValueOf(i interface{}) 接收的指针类型, 统计结构体字段数量时需要先使用 Elem 方法获取指针类型的值的封装")
	fmt.Println("如果需要修改值类型原变量的值, 则需要传入指针类型变量, 使用 func ValueOf(i interface{}).Elem().set*() 方式修改原变量的值")
	fmt.Println(tab, "func (v Value) Len() int // 返回 v 的长度, 如果 Kind 不是 Chan, Array, Map, Slice, String 则会 panic")
	fmt.Println(tab, "func (v Value) Cap() int // 返回 v 的容量, 如果 Kind 不是 Chan, Array, Slice 则会 panic")
	fmt.Println(tab, "func (v Value) NumField() int // 返回 v 持有的结构体类型值的字段数, 如果 v 的 Kind 不是 Struct 会 panic")
	fmt.Println(tab, "func (v Value) NumMethod() int // 返回 v 持有值的方法集的方法数目")
	fmt.Println(tab, "func (v Value) MethodByName(name string) Value")
	fmt.Println(tab, tab, "// 返回 v 的名为 name 的方法的已绑定(到 v 的持有值的)的状态的函数形式的 Value 封装, 如果未找到返回 Value 零值")
	fmt.Println(tab, "func (v Value) Call(in []Value) []Value // 使用输入的参数 in 调用 v 持有的函数")
	fmt.Println("---------------")

	fmt.Println("示例见下方的 Animal 结构体:")
	fmt.Println(tab, "reflect.Type NumField() 或 reflect.Value NumField(), 统计结构体的字段数量时, 忽略结构体中字段首字母大小写,")
	fmt.Println(tab, tab, "当类型为指针类型时, 需要先调用 Elem 方法获取 指针类型指向的值的 value 封装")
	fmt.Println(tab, "reflect.Type NumMethod() 或 reflect.Value NumMethod(), 统计结构体的方法数量时")
	fmt.Println(tab, tab, "当为值类型时, 忽略接受者变量类型指针类型的所有方法和值类型的方法名首字母小写")
	fmt.Println(tab, tab, "当为指针类型时, 忽略接受者变量类型(值类型和指针类型)方法名首字母小写")
	fmt.Println("---------------")

	fmt.Println("使用 reflect 获取结构体的方法时, 方法是以方法名的 ascii 码值排序, 和声明位置的顺序没有关系")
	fmt.Println("使用 Call 调用结构体的方法时, 参数为 []reflect.Value 类型切片, 需要处理返回值 []reflect.Value 类型切片")
	fmt.Println(tab, "var params = []reflect.Value // 声明 reflect.Value 类型切片")
	fmt.Println(tab, "params = append(params, reflect.ValueOf(10)) // 添加数据")
	fmt.Println(tab, "params = append(params, reflect.ValueOf(30)) // 添加数据")
	fmt.Println(tab, "res := rVal.Method(1).Call(params)  // 调用结构体方法并传入参数, 如果不需要参数则传入 nil")
	fmt.Println("---------------")

	fmt.Println("使用 reflect 操作结构体, 如果需要使用类型断言时 ")
	fmt.Println(tab, "当为值类型时, 类型断言使用 i.(Animal) 或者 i.(type) case Animal")
	fmt.Println(tab, "当为指针类型时, 类型断言使用 i.(*Animal) 或者 i.(type) case *Animal")
	fmt.Println("-----------------------------------")

	var s = Person{Name: "学生1", Age: 18, Address: "beijing"}
	fmt.Printf("修改前: s = %+v\n", s)               // {Name:学生1 Address:beijing Age:18}
	fmt.Printf("修改后: s = %+v\n", reflectTest1(s)) // &{Name:学生2 Address:shanghai Age:19}
	fmt.Println("---------------")
	rType := reflect.StructOf([]reflect.StructField{
		{
			Name: "Width",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"width"`,
		},
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
	})
	fmt.Printf("rType %+v\n", rType) // struct { Width float64 "json:\"width\""; Height float64 "json:\"height\"" }
	rVal := reflect.New(rType)       // 创建结构体类型实例
	rVal.Elem().FieldByName("Width").SetFloat(float64(80.90))
	rVal.Elem().FieldByName("Height").SetFloat(float64(60.70))
	fmt.Printf("rVal.Elem().FieldByName(\"Width\").Float() 值为 %e\n", rVal.Elem().FieldByName("Width").Float())   // 8.090000e+01
	fmt.Printf("rVal.Elem().FieldByName(\"Height\").Float() 值为 %E\n", rVal.Elem().FieldByName("Height").Float()) // 6.070000E+01
	fmt.Printf("rVal.Interface() %+v\n", rVal.Interface())                                                       // &{Width:80.9 Height:60.7}
	fmt.Printf("rType.Field(0) 为 %+v\n", rType.Field(0))                                                         // {Name:Width PkgPath: Type:float64 Tag:json:"width" Offset:0 Index:[0] Anonymous:false}
	fmt.Printf("rType.Field(1) 为 %+v\n", rType.Field(1))                                                         // {Name:Height PkgPath: Type:float64 Tag:json:"height" Offset:8 Index:[1] Anonymous:false}

	fmt.Println("---------reflectStruct01(a)---------")
	reflectStruct01(Animal{Name: "dog", Color: "yellow", Age: 20})

	fmt.Println("---------reflectStruct02(&a)---------")
	a := Animal{Name: "cat", Color: "black", Age: 18}
	reflectStruct02(&a)
	fmt.Println("-----------------------------------")
}

func reflectTest1(s interface{}) *Person {
	// var stu *Person // 此处需要声明 指针类型变量, reflect.New() 返回的是指针类型
	rType := reflect.TypeOf(s)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem() // 如果为指针类型则需要 Elem 获取真实类型
	}
	rVal := reflect.New(rType)                                                                                      // 根据传入的 reflect.Type 的类型返回新申请的内存的指针指向的值的 value 封装
	fmt.Printf("rType %+v , rVal %+v \n", rType, rVal)                                                              //  main.Person &{Name: Address: Age:0}
	fmt.Printf("rVal.Elem().NumField() is %d, rVal.NumMethod() is %d \n", rVal.Elem().NumField(), rVal.NumMethod()) // 3 1
	for i, len := 0, rType.NumField(); i < len; i++ {
		// main.Person 0 Name
		// main.Person 1 Address
		// main.Person 2 Age
		fmt.Printf("struct %+v %dth field is %+v \n", rType, i, rType.Field(i).Name)
	}
	name, addr, age := "学生2", "shanghai", 19
	fmt.Printf("修改 Name = %s, Address = %s, Age = %d\n", name, addr, age)
	rVal.Elem().FieldByName("Name").SetString(name)
	rVal.Elem().FieldByName("Address").SetString(addr)
	rVal.Elem().FieldByName("Age").SetUint(uint64(age))
	fmt.Printf("rVal.Elem().FieldByName(\"Age\").Uint() 值为 %d\n", rVal.Elem().FieldByName("Age").Uint()) // 19
	for i, len := 0, rVal.NumMethod(); i < len; i++ {
		fmt.Println("-- 通过 rVal.Method(i int).Call(nil) 调用结构体的方法")
		fmt.Print("调用方式1: rVal.Method(i).Call(nil) ")
		rVal.Method(i).Call(nil) // 调用方法
	}
	fmt.Println("-- 通过指定方法名调用结构体的方法")
	fmt.Print("调用方式2: rVal.MethodByName(\"SayHello\").Call(nil) ") // 需要判断返回值是否是 Value 零值
	rVal.MethodByName("SayHello").Call(nil)                        // 调用方法
	fmt.Println("--")
	fmt.Println("new(Person) ", new(Person))
	// reflect.Value => interface{} 后使用 类型断言 真实类型
	if stu, ok := rVal.Interface().(*Person); ok {
		return stu
	}
	return new(Person)
}

func reflectStruct01(a interface{}) {
	fmt.Println(a) // {dog yellow 20 0 {    0 0}}
	rType := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	// main.Animal {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
	fmt.Printf("a 类型为 %T 值为 %+v\n", a, a)
	fmt.Println("-----------")

	fmt.Printf("rType.NumField() 获取到 %d 个字段, rType.NumMethod() 获取到 %d 个方法\n", rType.NumField(), rType.NumMethod()) // 5 2
	fmt.Printf("rVal.NumField() 获取到 %d 个字段, rVal.NumMethod() 获取到 %d 个方法\n", rVal.NumField(), rVal.NumMethod())     // 5 2
	fmt.Printf("rVal.Kind() 类别为 %s\n", rVal.Kind())                                                                // struct
	fmt.Println("-----------")

	var sf = reflect.VisibleFields(rType)
	// []reflect.StructField
	fmt.Printf("reflect.VisibleFields(rType) 类型为 %T 值为 %v\n", sf, sf)
	for i := 0; i < len(sf); i++ {
		tagName, ok := sf[i].Tag.Lookup("json")
		fmt.Print(tab)
		fmt.Printf("是否导出 sf[%d].IsExported() %v  ", i, sf[i].IsExported())
		fmt.Printf("获取Tag: sf[%d].Tag.Lookup(\"json\") 结果为 ", i)
		if ok {
			fmt.Println(tagName)
		} else {
			fmt.Println("not found")
		}
	}
	fmt.Println("-----------")

	for i, len := 0, rType.NumField(); i < len; i++ {
		name := rType.Field(i).Name
		// 0 0 Name 0 dog
		// 1 1 Color 1 yellow
		fmt.Printf("[%d] rType.Field(%d).Name 获取到的 key 为 %s, rVal.Field(%d) 获取到的 value 为 %v\n", i, i, name, i, rVal.Field(i))
		sa := rVal.Interface() // 以 空接口 类型返回当前 v 保存的值, 通过类型断言可以转为实际的对象
		if instance, ok := sa.(Animal); ok {
			if strings.ToLower(name) == "color" {
				fmt.Println(tab, "当 key == color 时, value 修改为 white")
				fmt.Println(tab, "只能类型断言 instance, ok := sa.(Animal); ok 修改原对象的值 ")
				fmt.Print(tab, tab)
				fmt.Printf("修改前为 %+v\n", instance) // {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				instance.Color = "white"
				fmt.Print(tab, tab, "instance.Color = \"white\"\n")
				fmt.Print(tab, tab)
				fmt.Printf("修改后为 %+v\n", instance) // {Name:dog Color:white Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				fmt.Print(tab, tab)
				fmt.Printf("a 的值为 %+v\n", a) // {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
			}
			if strings.ToLower(name) == "age" {
				fmt.Println(tab, "只能类型断言 instance, ok := sa.(Animal); ok 修改原对象的值 ")
				fmt.Print(tab, tab)
				fmt.Printf("修改前为 %+v\n", instance) // {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				instance.Age = 99
				fmt.Print(tab, tab, "instance.Age = 99\n")
				fmt.Print(tab, tab)
				fmt.Printf("修改后为 %+v\n", instance) // {Name:dog Color:yellow Age:99 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				fmt.Print(tab, tab)
				fmt.Printf("a 的值为 %+v\n", a) // {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
			}
		}
	}
	fmt.Println("-----------")

	fmt.Println("结构体的方法名按 ascii 值排序, 和定义位置的顺序没有关系")
	for i, len := 0, rType.NumMethod(); i < len; i++ {
		name := rType.Method(i).Name
		fmt.Print(tab)
		// 0 SayHello
		// 1 Sum
		fmt.Printf("rType.Method(%d).Name 获取的方法名为 %s  ", i, name)
		// 0 true
		// 1 true
		fmt.Printf("rType.Method(%d).IsExported() 是否为导出方法 %v \n", i, rType.Method(i).IsExported())
		if strings.ToLower(name) != "sum" {
			fmt.Print(tab)
			rVal.Method(i).Call(nil) // animal SayHello...(initial uppercase)
		}
	}
	fmt.Printf("a 最终结果为 %+v\n", a) // {Name:dog Color:yellow Age:20 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
}

func reflectStruct02(a interface{}) {
	rVal := reflect.ValueOf(a)
	fmt.Println(rVal.Elem().Kind())
	if rVal.Elem().Kind() != reflect.Struct {
		fmt.Println("参数不是一个结构体")
		return
	}
	rType := reflect.TypeOf(a)
	// *main.Animal &{Name:cat Color:black Age:18 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
	fmt.Printf("a 类型为 %T 值为 %+v\n", a, a)
	fmt.Println("调用实例方法-方式1 类型断言 instance, ok := a.(*Animal);ok ")
	if instance, ok := a.(*Animal); ok {
		fmt.Print("--", tab)
		instance.sayHello() // *ptr animal sayHello...(initial lowercase)
		fmt.Print("--", tab)
		instance.SayHello() // animal SayHello...(initial uppercase)
		fmt.Print("--", tab)
		instance.sayGoodBye() // animal sayGoodBye...(initial lowercase)
		fmt.Print("--", tab)
		instance.SayGoodBye() // *ptr animal SayGoodBye...(initial uppercase)
	}
	fmt.Println("调用实例方法-方式2 类型断言 switch instance := a.(type) case *Animal:")
	switch instance := a.(type) {
	case *Animal:
		fmt.Print("--", tab)
		instance.sayHello() // *ptr animal sayHello...(initial lowercase)
		fmt.Print("--", tab)
		instance.SayHello() // animal SayHello...(initial uppercase)
		fmt.Print("--", tab)
		instance.sayGoodBye() // animal sayGoodBye...(initial lowercase)
		fmt.Print("--", tab)
		instance.SayGoodBye() // *ptr animal SayGoodBye...(initial uppercase)
	default:
		fmt.Println("没有匹配到指定类型")
	}
	fmt.Println("-----------")

	fmt.Printf("rType.Elem().NumField() 获取到 %d 个字段, rType.NumMethod() 获取到 %d 个方法\n", rType.Elem().NumField(), rType.NumMethod()) // 5 3
	fmt.Printf("rVal.Elem().NumField() 获取到 %d 个字段, rVal.NumMethod() 获取到 %d 个方法\n", rVal.Elem().NumField(), rVal.NumMethod())     // 5 3
	fmt.Printf("rVal.Kind() 类别为 %s\n", rVal.Kind())                                                                              // ptr
	fmt.Println("-----------")

	var sf = reflect.VisibleFields(rType.Elem())
	fmt.Printf("reflect.VisibleFields(rType) 类型为 %T 值为 %v\n", sf, sf) // []reflect.StructField
	for i := 0; i < len(sf); i++ {
		tagName, ok := sf[i].Tag.Lookup("json")
		fmt.Print(tab)
		fmt.Printf("是否导出 sf[%d].IsExported() %v  ", i, sf[i].IsExported())
		fmt.Printf("获取Tag: sf[%d].Tag.Lookup(\"json\") 结果为 ", i)
		if ok {
			fmt.Println(tagName)
		} else {
			fmt.Println("not found")
		}
	}
	fmt.Println("-----------")

	for i, len := 0, rType.Elem().NumField(); i < len; i++ {
		name := rType.Elem().Field(i).Name
		fmt.Printf("[%d] rType.Elem().Field(%d).Name 获取到的 key 为 %s, rVal.Elem().Field(%d) 获取到的 value 为 %v\n", i, i, name, i, rVal.Elem().Field(i))
		sa := rVal.Interface() // 以 空接口 类型返回当前 v 保存的值, 通过类型断言可以转为实际的对象
		if instance, ok := sa.(*Animal); ok {
			if strings.ToLower(name) == "color" {
				fmt.Println(tab, "当 key == color 时, value 修改为 white")
				fmt.Println(tab, "方式1: 类型断言 instance, ok := sa.(*Animal); ok 修改原对象的值 ")
				fmt.Print(tab, tab)
				fmt.Printf("修改前为 %+v\n", instance) // &{Name:cat Color:black Age:18 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				instance.Color = "white"           // 等价于 rVal.Elem().Field(i).Set(reflect.ValueOf("white"))
				fmt.Print(tab, tab, "instance.Color = \"white\" 等价于 rVal.Elem().Field(i).Set(reflect.ValueOf(\"white\"))\n")
				fmt.Print(tab, tab)
				fmt.Printf("修改后为 %+v\n", instance) // &{Name:cat Color:white Age:18 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				fmt.Print(tab, tab)
				fmt.Printf("a 的值为 %+v\n", a) // &{Name:cat Color:white Age:18 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
			}
			if strings.ToLower(name) == "age" {
				fmt.Println(tab, "方式2: 使用 reflect API 修改原对象的值")
				fmt.Print(tab, tab)
				fmt.Printf("修改前为 %+v\n", rVal) // &{Name:cat Color:white Age:18 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				fmt.Print(tab, tab, "rVal.Elem().Field(i).SetUint(uint64(99)) // Age 字段类型为 uint8, 传入参数类型需要转为 uint64\n")
				rVal.Elem().Field(i).SetUint(uint64(99)) // Age 字段类型为 uint8, 此处传入参数类型需要转为 uint64
				fmt.Print(tab, tab)
				fmt.Printf("rVal.Elem().Field(\"%d\").Uint() 值为 %d\n", i, rVal.Elem().Field(i).Uint()) // 99
				fmt.Print(tab, tab)
				fmt.Printf("修改后为 %+v\n", rVal) // &{Name:cat Color:white Age:99 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
				fmt.Print(tab, tab)
				fmt.Printf("a 的值为 %+v\n", a) // &{Name:cat Color:white Age:99 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
			}
		}
	}
	fmt.Println("-----------")

	fmt.Println("结构体的方法名按 ascii 值排序, 和定义位置的顺序没有关系")
	for i, len := 0, rType.NumMethod(); i < len; i++ {
		name := rType.Method(i).Name
		fmt.Print(tab)
		// 0 SayGoodBye
		// 1 SayHello
		// 2 Sum
		fmt.Printf("rType.Method(%d).Name 获取的方法名为 %s  ", i, name)
		// 0 true
		// 1 true
		// 2 true
		fmt.Printf("rType.Method(%d).IsExported() 是否为导出方法 %v \n", i, rType.Method(i).IsExported())
		if strings.ToLower(name) == "sum" {
			fmt.Println(tab, "如果方法名为 sum 时,调用该方法, Call 方法接收的参数是一个 reflect.Value 类型的切片, 返回值是一个 reflect.Value 类型的切片")
			// 2 30
			fmt.Println(tab, "rVal.Method(", i, ").Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)})[0] 的结果为", rVal.Method(i).Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)})[0])
		} else {
			rVal.Method(i).Call(nil)
		}
	}
	fmt.Println("-----------")
	fmt.Printf("最终结果 a 为 %+v\n", a) // &{Name:cat Color:white Age:99 height:0 Address:{Province: City: Region: Street: Long:0 Lat:0}}
}

type Person struct {
	Name, Address string
	Age           uint8
}

func (s *Person) SayHello() {
	fmt.Println(s.Name, "is", s.Age, "year old. live in", s.Address)
}

type Address struct {
	Province string  `json:"province"`
	City     string  `json:"city"`
	Region   string  `json:"region"`
	Street   string  `json:"street"`
	Long     float32 `json:"longitude"`
	Lat      float32 `json:"latitude"`
}

type Animal struct {
	Name   string `json:"name" validate:"required"`
	Color  string `json:"color"`
	Age    uint8  `json:"age" validate:"min=18,required"`
	Height uint8  `json:"height"` // 字段名小写
	Address
}

func (a Animal) Sum(i, j int) int {
	return i + j
}
func (a Animal) sayGoodBye() {
	fmt.Println("animal sayGoodBye...(initial lowercase)")
}
func (a Animal) SayHello() {
	fmt.Println("animal SayHello...(initial uppercase)")
}
func (a *Animal) sayHello() {
	fmt.Println("*ptr animal sayHello...(initial lowercase)")
}
func (a *Animal) SayGoodBye() {
	fmt.Println("*ptr animal SayGoodBye...(initial uppercase)")
}
