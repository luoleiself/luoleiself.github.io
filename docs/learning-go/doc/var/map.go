package main

import (
	"fmt"
	_ "sync"
)

type Person struct {
	Name string
	Age  uint8
}

func mapNote() {
	// 非线程安全
	// 1. 典型的使用场景，大多数情况下不需要从多个 goroutine 安全访问
	// 2. 性能考量，如果内置锁机制，会降低所有程序的性能
	// 3. 灵活性，可以根据实际需求选择最适合的同步策略
	fmt.Println("--------------mapNote()----------------")
	fmt.Println("声明: map 是引用类型, map 是一种特殊的数据结构, 一种元素对的无序集合")
	fmt.Println(tab, "var 变量名 map[key类型]value类型")
	fmt.Println(tab, "map 只声明未初始化的默认值为 nil, 有内存地址, length 为 0, 但不能进行 map 操作")
	fmt.Println(tab, "map 的 key 的数据类型只能是可参与比较运算的类型(即 == 或 != ), 如布尔类型、数值类型(整、浮、复)、字符串类型、数组、结构体类型。 引用类型(切片、map、函数)等不能作为 key 的数据类型")
	fmt.Println(tab, "map 的 value 可以是任意数据类型")
	fmt.Println(tab, "对 map 中不存在的 key 操作时, 是对 key 对应的 value 的零值操作")

	fmt.Println("------------------------------")
	fmt.Println("\033[1;32m声明方式1\033[0m: 声明 map 时初始化")
	fmt.Println(tab, "对 map 中不存在的 key 操作时, 是对 key 对应的 value 的零值操作")
	var m1 = map[string]int{"hello": 5, "world": 10}
	fmt.Print("var m1 = map[string]int{\"hello\": 5, \"world\": 10}\n")
	fmt.Printf("m1 的类型为 %T 长度为 %d 值为 %v\n", m1, len(m1), m1) // map[string]int 2 map[hello:5 world:10]
	fmt.Printf("m1[\"hello\"] 的值为 %d\n", m1["hello"])        // 5
	fmt.Printf("m1[\"world\"] 的值为 %d\n", m1["world"])        // 10
	fmt.Printf("m1[\"gg\"] 的值为 %d\n", m1["gg"])              // gg 不存在, 0
	m1["he"]++
	fmt.Print("m1[\"he\"]++\n")   // he 不存在
	fmt.Printf("m1 的值为 %v\n", m1) // map[he:1 hello:5 world:10]
	fmt.Println("-----------")

	fmt.Println("\033[1;32m声明方式2\033[0m: 内置函数 make 声明的 map 默认值不为 nil, 可正常操作, 等价于 map[string]float32{}")
	m4 := make(map[string]float32, 2) // make 方法创建 map, 第二个参数没有具体意义
	fmt.Println("m4 := make(map[string]float32, 2)")
	fmt.Printf("m4 的类型为 %T 元素个数为 %d 值为 %v\n", m4, len(m4), m4) // map[string]float32 0 map[]
	m4["chinese"] = 90.52
	m4["english"] = 95.18
	m4["math"] = 97
	fmt.Printf("m4 添加元素后元素个数为 %d 值为 %v \n", len(m4), m4) // 3 map[chinese:90.52 english:95.18 math:97]
	fmt.Println("------------------------------")

	var test1 = map[string]Month{"January": January, "February": February, "March": March, "April": April, "May": May, "June": June, "July": July, "August": August, "September": September, "October": October, "November": November, "December": December}
	fmt.Printf("test1 的类型为 %T 元素个数为 %d \n值为 %v\n", test1, len(test1), test1) // map[string]main.Month  12
	fmt.Println(`test1["February"] = `, test1["February"])                   // 2
	fmt.Println(`test1["September"] = `, test1["September"])                 // 9
	fmt.Println(`test1["December"] = `, test1["December"])                   // 12
	fmt.Println("获取元素, 未定义的 key 将返回声明时的类型默认值, 通过 map[key] 返回的第二个值判断是否获取成功")
	fmt.Println(`test1["Julyy"] = `, test1["julyy"]) // 0
	if res, ok := test1["July"]; ok {
		fmt.Println("获取成功, 值为 ", res) // 7
	} else {
		fmt.Println("获取失败, 未找到指定key ", res)
	}
	fmt.Println("----------")
	fmt.Println("map的遍历, 输出的顺序随机不固定")
	for key, val := range test1 {
		fmt.Printf("test1[%s] = %d\n", key, val)
	}
	delete(test1, "March")
	// 11 map[April:4 August:8 December:12 February:2 January:1 July:7 June:6 May:5 November:11 October:10 September:9]
	fmt.Printf("删除 March 后 test1 元素个数为 %d 的值为 %v\n", len(test1), test1)

	fmt.Println("------map作为函数参数传递----------")
	var test2 = map[string]int{"chinese": 100, "english": 99}
	fmt.Print("var test2 = map[string]int{\"chinese\": 100, \"english\": 99}\n")
	fmt.Printf("test2 的值为 %v\n", test2) // map[chinese:100 english:99]
	fmt.Println("调用函数 testMap(test2)")
	testMap(test2)
	fmt.Printf("main() 函数内的 test2 的值为 %v\n", test2) //  map[chinese:100 english:60]
	fmt.Println("------------------------------")

	var person1 = Person{"张三", 18}
	var person2 = Person{"李四", 22}
	var map1 = map[Person]int{person1: 100, person2: 225}
	var map2 = map[Person]int{person1: 200, person2: 250}
	fmt.Print("var person1 = Person{\"张三\", 18}\n")
	fmt.Print("var person2 = Person{\"李四\", 22}\n")
	fmt.Print("var map1 = map[Person]int{person1: 100, person2: 225}\n")
	fmt.Printf("map1 的类型为 %T 值为 %+v\n", map1, map1)                                // map[main.Person]int map[{Name:张三 Age:18}:100 {Name:李四 Age:22}:225]
	fmt.Println("map1[person2] = ", map1[person2])                                 // 225
	fmt.Println("map1[person2] == map1[person2] ", map1[person2] == map2[person2]) // true
	fmt.Println("------------------------------")
}
func testMap(test2 map[string]int) {
	test2["english"] = 60
	fmt.Print("testMap 函数内修改 test2[\"english\"] = 60\n")
	fmt.Printf("testMap 函数内 test2 的值为 %v\n", test2) // map[chinese:100 english:60]
}
