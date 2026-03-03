package main

import (
	"fmt"
	"reflect"
	"time"
)

func reflectMap() {
	fmt.Println("------------reflectMap()----------------")
	fmt.Println("func MapOf(key, elem Type) Type // 返回具有给定 key 和 elem 类型的 map 类型")
	fmt.Println(tab, "func (t Type) Key() Type // 返回 map 类型的键的类型. 如非映射类型将 panic")
	fmt.Println("func MakeMap(t Type) Value // 创建一个具有指定类型的新 map")
	fmt.Println(tab, "func (v Value) SetMapIndex(key, val Value)")
	fmt.Println(tab, "  // 用来给v的映射类型持有值添加/修改键值对, 如果val是Value零值, 则是删除键值对")
	fmt.Println(tab, "func (v Value) MapIndex(key Value) Value // 获取指定 key 的值")
	fmt.Println(tab, "func (v Value) MapKeys() []Value // 获取所有的 key")
	fmt.Println(tab, "func (v Value) MapRange() *MapIter // 获取 map 的迭代器")
	fmt.Println(tab, tab, "func (it *MapIter) Key() Value // 获取 key")
	fmt.Println(tab, tab, "func (it *MapIter) Next() bool // 是否有个下一个")
	fmt.Println(tab, tab, "func (it *MapIter) Value() Value // 获取 value")
	fmt.Println("----------------------------")

	// 方式1:
	// var mVal = reflect.ValueOf(map[string]int{"Sunday": 0})
	// var mType = mVal.Type()
	// var mrval = reflect.MakeMap(mType)
	// 方式2:
	// var mVal = reflect.ValueOf(map[string]int{"Sunday": 0})
	// var rMapIndirect = reflect.Indirect(mVal)
	// var mrval = reflect.MakeMap(rMapIndirect.Type())
	// 方式3:
	var mrt = reflect.MapOf(reflect.TypeOf("hello"), reflect.TypeOf(1)) // 返回具有给定 key 和 elem 类型的 map 类型
	var mrval = reflect.MakeMap(mrt)                                    // 创建一个具有指定类型的新 map

	// 类型断言, 必须和 MapOf 的参数元素类型一致, 否则 runtime panic: interface conversion: interface {} is map[string]int, not map[string]string
	var mv = mrval.Interface().(map[string]int)

	fmt.Println(`
  // 方式1:
  // var mVal = reflect.ValueOf(map[string]int{"Sunday": 0})
  // var mType = mVal.Type()
  // var mrval = reflect.MakeMap(mType)
  // 方式2:
  // var mVal = reflect.ValueOf(map[string]int{"Sunday": 0})
  // var rMapIndirect = reflect.Indirect(mVal)
  // var mrval = reflect.MakeMap(rMapIndirect.Type())
  // 方式3:
  var mrt = reflect.MapOf(reflect.TypeOf("hello"), reflect.TypeOf(1)) // 返回具有给定 key 和 elem 类型的 map 类型
  var mrval = reflect.MakeMap(mrt)                                      // 创建一个具有指定类型的新 map

  // 类型断言, 必须和 MapOf 的参数元素类型一致, 否则 runtime panic: interface conversion: interface {} is map[string]int, not map[string]string
  var mv = mrval.Interface().(map[string]int) 

  mv["Sunday"] = int(time.Sunday)
  mv["Monday"] = int(time.Monday)
  mv["Tuesday"] = int(time.Tuesday)
  mv["Wednesday"] = int(time.Wednesday)
  mv["Thursday"] = int(time.Thursday)
  mv["Friday"] = int(time.Friday)
  mv["Saturday"] = int(time.Saturday)`)

	mv["Sunday"] = int(time.Sunday)
	mv["Monday"] = int(time.Monday)
	mv["Tuesday"] = int(time.Tuesday)
	mv["Wednesday"] = int(time.Wednesday)
	mv["Thursday"] = int(time.Thursday)
	mv["Friday"] = int(time.Friday)
	mv["Saturday"] = int(time.Saturday)
	fmt.Printf("mv = %+v, length = %d\n", mv, len(mv)) // map[Friday:5 Monday:1 Saturday:6 Sunday:0 Thursday:4 Tuesday:2 Wednesday:3] 7
	fmt.Println("mrval.MapKeys()", mrval.MapKeys())    // [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	fmt.Println("mrval.MapIndex(reflect.ValueOf(\"Saturday\"))", mrval.MapIndex(reflect.ValueOf("Saturday")))
	fmt.Println("----------------------------")

	fmt.Println(`
  var mIter = mrval.MapRange()
  for mIter.Next() {
    fmt.Println(tab, "key", mIter.Key(), "value", mIter.Value())
  }`)
	var mIter = mrval.MapRange()
	for mIter.Next() {
		fmt.Println(tab, "key =>", mIter.Key(), "value =>", mIter.Value()) // key => ** value => **
	}
	fmt.Println("----------------------------")

	var mrval2 = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(1)))
	mrval2.SetMapIndex(reflect.ValueOf("age"), reflect.ValueOf(18))
	mrval2.SetMapIndex(reflect.ValueOf("score"), reflect.ValueOf(99))
	fmt.Println("var mrval2 = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(\"\"), reflect.TypeOf(1)))")
	fmt.Println("mrval2.SetMapIndex(reflect.ValueOf(\"age\"), reflect.ValueOf(18))")
	fmt.Println("mrval2.SetMapIndex(reflect.ValueOf(\"score\"), reflect.ValueOf(99))")
	fmt.Printf("mrval2 = %v\n", mrval2) // map[age:18 score:99]
	mrval2.SetMapIndex(reflect.ValueOf("age"), reflect.ValueOf(20))
	fmt.Println("mrval2.SetMapIndex(reflect.ValueOf(\"age\"), reflect.ValueOf(20))")
	fmt.Printf("mrval2 = %v\n", mrval2) // map[age:20 score:99]
	fmt.Println("----------------------------")

	fmt.Println(`
  var tom = struct {
    Name string
    Age  uint8
  }{Name: "Tom", Age: 18} // 匿名结构体
  var mRType = reflect.MapOf(reflect.TypeOf(tom), reflect.TypeOf(3.14))
  fmt.Println("mRType.Key().Name()", mRType.Key().Name()) // 匿名结构体输出空, 普通结构体打印结构体名称
  fmt.Println("mRType.Elem().Name()", mRType.Elem().Name()) // float64
  fmt.Println("mRType.Key().Kind()", mRType.Key().Kind()) // struct`)
	fmt.Println("--------------")
	var tom = struct {
		Name string
		Age  uint8
	}{Name: "Tom", Age: 18} // 匿名结构体
	var mRType = reflect.MapOf(reflect.TypeOf(tom), reflect.TypeOf(3.14))
	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，Name() 会返回 ""
	fmt.Println("mRType.Key().Name()", mRType.Key().Name())   // 匿名结构体输出空, 普通结构体打印结构体名称
	fmt.Println("mRType.Elem().Name()", mRType.Elem().Name()) // float64
	fmt.Println("mRType.Key().Kind()", mRType.Key().Kind())   // struct
	fmt.Println("----------------------------")
}
