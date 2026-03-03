package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string
	Age  uint8
}

func reflectTest01(i interface{}) {
	// 获取 reflect.Type 类型
	// 方式2 使用 ValueOf 获取, 然后通过值的 Type 方法 获取 reflect.Type 类型
	rType := reflect.TypeOf(i)
	fmt.Printf("reflect.TypeOf(%v) 的类型为 %T 值为 %v\n", i, rType, rType)

	// 获取 reflect.Value 类型
	rVal := reflect.ValueOf(i)
	// rType := rVal.Type() // 获取 i 的 reflect.Type 类型
	fmt.Printf("reflect.ValueOf(%v) 的类型为 %T 值为 %v\n", i, rVal, rVal)

	// 获取 reflect.Kind 类型
	rKind := rVal.Kind()
	fmt.Printf("rVal.Kind() 的类型为 %T 值为 %v\n", rKind, rKind)
}

func TestDemo1(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		// t.Parallel()
		fmt.Println(`
		var num2 = 100
		rVal := reflect.ValueOf(&num2)
		rVal.Elem().SetInt(250) // Elem 方法返回 接口v 包含的值或 指针v 指向的值. 如果v的种类不是接口或Ptr, 它会崩溃. 如果v为零, 则返回零值`)
		var num2 = 100
		fmt.Println("num2=", num2) // 100
		rVal := reflect.ValueOf(&num2)
		rVal.Elem().SetInt(250)                                     // Elem 方法返回 接口v 包含的值或 指针v 指向的值. 如果v的种类不是接口或Ptr, 它会崩溃. 如果v为零, 则返回零值
		fmt.Printf("rVal.Elem().Int() 的值为 %d\n", rVal.Elem().Int()) // 250
		fmt.Println("\nnum2=", num2)                                // 250
		fmt.Println("-----------------------------------")

		var r1 = []reflect.Type{reflect.TypeOf(1), reflect.TypeOf(2.2)}
		fmt.Println("var r1 = []reflect.Type{reflect.TypeOf(1), reflect.TypeOf(2.2)}")
		fmt.Printf("r1 的类型为 %T 值为 %v\n", r1, r1) // []reflect.Type [int float64]
		var r2 = []reflect.Value{reflect.ValueOf("hello"), reflect.ValueOf(18)}
		fmt.Println("var r2 = []reflect.Value{reflect.ValueOf(\"hello\"), reflect.ValueOf(18)}")
		fmt.Printf("r2 的类型为 %T 值为 %v\n", r2, r2) // []reflect.Value [hello <int Value>]
		fmt.Println("-----------------------------------")
	})
}

func TestDemo2(t *testing.T) {
	zto1 := reflect.Zero(reflect.TypeOf("hello"))
	zto2 := reflect.Zero(reflect.TypeOf(-0.9))
	zto3 := reflect.Zero(reflect.TypeOf(18))
	zto4 := reflect.Zero(reflect.TypeOf('g'))
	zto5 := reflect.Zero(reflect.TypeOf(true))
	zto6 := reflect.Zero(reflect.TypeOf([...]byte{'a', 'b', 'c'}))
	fmt.Printf("reflect.Zero(reflect.TypeOf(\"hello\")) 的类型为 %T 值为 %v\n", zto1, zto1)                // reflect.Value
	fmt.Printf("reflect.Zero(reflect.TypeOf(-0.9)) 类型为 %T 值为 %v\n", zto2, zto2)                      // reflect.Value 0
	fmt.Printf("reflect.Zero(reflect.TypeOf(18)) 的类型为 %T 值为 %v\n", zto3, zto3)                       // reflect.Value 0
	fmt.Printf("reflect.Zero(reflect.TypeOf('g')) 的类型为 %T 值为 %v\n", zto4, zto4)                      // reflect.Value 0
	fmt.Printf("reflect.Zero(reflect.TypeOf(true)) 的类型为 %T 值为 %v\n", zto5, zto5)                     // reflect.Value false
	fmt.Printf("reflect.Zero(reflect.TypeOf([...]byte{'a', 'b', 'c'})) 的类型为 %T 值为 %v\n", zto6, zto6) // reflect.Value [0 0 0]
	fmt.Printf("-----------------------------------")
}

func TestDemo3(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		// t.Parallel()
		var num = 100
		fmt.Println("var num = 100")
		// *reflect.rtype int
		// reflect.Value 100
		// reflect.Kind int
		reflectTest01(num)
	})
	t.Run("string", func(t *testing.T) {
		// t.Parallel()
		var st = "hello world"
		fmt.Println("var st = \"hello world\"")
		// *reflect.rtype string
		// reflect.Value hello world
		// reflect.Kind string
		reflectTest01(st)
	})
	t.Run("array", func(t *testing.T) {
		// t.Parallel()
		var arr = [5]byte{'h', 'e', 'l', 'l', 'o'}
		fmt.Println("var arr = [5]byte{'h', 'e', 'l', 'l', 'o'}")
		// *reflect.rtype [5]uint8
		// reflect.Value [104 101 108 108 111]
		// reflect.Kind array
		reflectTest01(arr)
	})
	t.Run("map", func(t *testing.T) {
		// t.Parallel()
		var m = map[string]uint8{"sunday": 0, "Monday": 1, "Tuesday ": 2, "Wednesday ": 3, "Thursday": 4, "Friday": 5, "Saturday": 6}
		fmt.Println("var m = map[string]uint8{\"sunday\": 0, \"Monday\": 1, \"Tuesday\": 2, \"Wednesday\": 3, \"Thursday\": 4, \"Friday\": 5, \"Saturday\": 6}")
		// *reflect.rtype map[string]uint8
		// reflect.Value map[Friday:5 Monday:1 Saturday:6 Thursday:4 Tuesday :2 Wednesday :3 sunday:0]
		// reflect.Kind map
		reflectTest01(m)
	})
	t.Run("slice", func(t *testing.T) {
		// t.Parallel()
		var sl = []bool{true, false, true, false}
		fmt.Println("var sl = []bool{true, false, true, false}")
		// *reflect.rtype []bool
		// reflect.Value [true false true false]
		// reflect.Kind slice
		reflectTest01(sl)
	})
	t.Run("struct", func(t *testing.T) {
		// t.Parallel()
		var stu = Student{Name: "student", Age: 18}
		fmt.Println("var stu = Student{Name: \"student\", Age: 18}")
		// *reflect.rtype main.Student
		// reflect.Value {student 18}
		// reflect.Kind struct
		reflectTest01(stu)
	})
}
