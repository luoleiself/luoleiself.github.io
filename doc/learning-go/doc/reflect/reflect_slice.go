package main

import (
	"fmt"
	"reflect"
)

func reflectSlice() {
	fmt.Println("-----------reflectSlice()-------------")
	fmt.Println("func SliceOf(t Type) Type // 返回元素类型为 t 的切片类型")
	fmt.Println("func MakeSlice(t Type, len, cap int) Value // 为指定的切片类型、长度和容量创建一个新的零初始化 slice")
	fmt.Println("func Append(s Value, x ...Value) Value // 向 reflect.Value 切片中追加 reflect.Value")
	fmt.Println("func AppendSlice(s, t Value) Value // 向 reflect.Value 切片中追加 reflect.Value ")
	fmt.Println("-------------")
	var srt = reflect.SliceOf(reflect.TypeOf("hello")) // 返回元素类型为 t 的切片类型
	var sms = reflect.MakeSlice(srt, 5, 5)             // 创建新的切片
	// 类型断言, 必须和 SliceOf 参数元素类型一致, 否则 runtime panic: interface conversion: interface {} is []string, not []int
	// var sv = sms.Interface().([]string)
	// sv[0] = "hello"
	// sv[1] = "world"
	sms = reflect.AppendSlice(sms, reflect.ValueOf([]string{"hello", "world"}))
	sms = reflect.Append(sms, reflect.ValueOf("你好"))
	var sv = sms.Interface().([]string)
	fmt.Println(`
  var srt = reflect.SliceOf(reflect.TypeOf("hello")) // 返回元素类型为 t 的切片类型
  var sms = reflect.MakeSlice(srt, 5, 5)             // 创建新的切片
	// 类型断言, 必须和 SliceOf 参数元素类型一致, 否则 runtime panic: interface conversion: interface {} is []string, not []int
  var sv = sms.Interface().([]string)                
  sv[0] = "hello"
  sv[1] = "world"
  // sms = reflect.AppendSlice(sms, reflect.ValueOf([]string{"hello", "world"}))
  // sms = reflect.Append(sms, reflect.ValueOf("你好"))
  // var sv = sms.Interface().([]string)	`)
	fmt.Printf("sv = %+s, length = %d, cap = %d\n", sv, len(sv), cap(sv)) // [     hello world 你好] 8 10
	fmt.Println("----------------------------")
}
