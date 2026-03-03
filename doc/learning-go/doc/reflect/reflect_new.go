package main

import (
	"fmt"
	"reflect"
	// _ "time"
)

func reflectNew() {
	fmt.Println("------------reflectNew()-------------")
	fmt.Println("func ArrayOf(length int, elem Type) Type // 使用 reflect 创建 array")
	fmt.Println("-------------")

	var art = reflect.ArrayOf(4, reflect.TypeOf(2)) // 根据数组长度和 reflect.Type 类型返回 reflect.Type 类型
	var arv = reflect.New(art)
	var arr *[4]int
	if ok := arv.CanInterface(); ok {
		arr = arv.Interface().(*[4]int) // reflect.Value => interface{} 后使用 类型断言 真实类型
		for i := 0; i < 4; i++ {
			arr[i] = i * i
		}
	} else {
		fmt.Println("arv can not use Interface method...")
	}
	fmt.Println(`
  var art = reflect.ArrayOf(4, reflect.TypeOf(2)) // 根据数组长度和 reflect.Type 类型返回 reflect.Type 类型
  var arv = reflect.New(art)
  var arr *[4]int
  if ok := arv.CanInterface();ok{
    arr = arv.Interface().(*[4]int) // reflect.Value => interface{} 后使用 类型断言 真实类型
    for i := 0; i < 4; i++ {
      arr[i] = i * i
    }
  } else {
    fmt.Println("arv can not use Interface method...")
  }`)
	fmt.Printf("arr = %+v, length = %d, cap = %d\n", arr, len(arr), cap(arr))                   // &[0 1 4 9]  4  4
	fmt.Printf("arv.Elem().Index(3) 的类型为 %T 值为 %v\n", arv.Elem().Index(3), arv.Elem().Index(3)) // reflect.Value 9
	fmt.Println("----------------------------")

	var rVal = reflect.ValueOf(123.456)
	if ok := rVal.CanConvert(reflect.TypeOf(110)); ok {
		var rv = rVal.Convert(reflect.TypeOf(110))
		fmt.Printf("var rVal = reflect.ValueOf(123.456) convert to reflect.TypeOf(110) 的类型为 %T 值为 %+v\n", rv, rv) // reflect.Value 123
	} else {
		fmt.Printf("var rVal = reflect.ValueOf(123.456) can not convert to reflect.TypeOf(110) %v\n", false)
	}
	fmt.Println("-------------")
	var rVal2 = reflect.ValueOf(123.456)
	if ok := rVal2.CanConvert(reflect.TypeOf("hello")); ok {
		var rv2 = rVal2.Convert(reflect.TypeOf("hello"))
		fmt.Printf("var rVal2 = reflect.ValueOf(123.456) convert to reflect.TypeOf(\"hello\") 的类型为 %T 值为 %v\n", rv2, rv2)
	} else {
		fmt.Printf("var rVal2 = reflect.ValueOf(123.456) can not convert to reflect.TypeOf(\"hello\") %v\n", false)
	}
	fmt.Println("----------------------------")

	var b = true
	var rVal3 = reflect.ValueOf(&b)
	fmt.Print("var b = true\nvar rVal3 = reflect.ValueOf(&b)\n")
	fmt.Printf("rVal3.Elem().Bool() 值为 %v\n", rVal3.Elem().Bool()) // true
	if ok := rVal3.Elem().CanSet(); ok {
		rVal3.Elem().SetBool(false)
		fmt.Printf("rVal3.Elem().SetBool(false) 后的值为 %v\n", rVal3.Elem().Bool()) // false
	}
	fmt.Println("----------------------------")

	var s = []int{1, 2, 3, 4}
	var sRVal = reflect.ValueOf(&s)
	fmt.Println("var s = []int{1, 2, 3, 4}\nvar sRVal = reflect.ValueOf(&s)")
	fmt.Printf("sRVal 的类型为 %T 值为 %v\n", sRVal, sRVal)  // reflect.Value  &[1 2 3 4]
	fmt.Printf("sRVal.IsNil() 值为 %t\n", sRVal.IsNil()) // false
	fmt.Println("-------------")
	var s2 *[]int
	var s2RVal = reflect.ValueOf(s2)
	fmt.Println("var s2 *[]int\nvar s2RVal = reflect.ValueOf(s2)")
	fmt.Printf("s2RVal 的类型为 %T 值为 %v\n", s2RVal, s2RVal) // reflect.Value <nil>
	fmt.Printf("s2RVal.IsNil() 值为 %t\n", s2RVal.IsNil()) // true
	fmt.Println("----------------------------")
}
