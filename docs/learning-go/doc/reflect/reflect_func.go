package main

import (
	"fmt"
	"reflect"
)

func reflectFunc() {
	fmt.Println("---------------reflectFunc()---------------")
	fmt.Println("func FuncOf(in, out []Type, variadic bool) Type // 返回具有给定参数和结果类型的函数类型, variadic 可变参数控制函数是否可变")
	fmt.Println(tab, "例如，如果 k 表示 int, e 表示字符串, FuncOf([]Type{k}, []Typ{e}, false) 表示 func(int) string")
	fmt.Println("func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value // 返回一个具有给定类型、包装函数fn的函数的Value封装")
	fmt.Println(tab, "函数 fn 的实现可以假设参数 Value切片 匹配 typ 类型指定的参数数目和类型. 如果 typ 表示一个可变参数函数类, 参数切片中最后一个Value本身必须是一个包含所有可变参数的切片. fn 返回的结果 Value切片 也必须匹配 typ 类型指定的结果数目和类型")
	fmt.Println("func (t Type) NumIn() int // 返回 func 类型的参数个数, 如果不是函数, 将会panic ")
	fmt.Println("func (t Type) In(i int) Type // 返回 func 类型的第 i 个参数的类型, 如非函数或者 i 不在 [0, NumIn()) 内将会 panic")
	fmt.Println("func (t Type) NumOut() int // 返回 func 类型的返回值个数, 如果不是函数, 将会 panic")
	fmt.Println("func (t Type) Out(i int) Type // 返回 func 类型的第 i 个返回值的类型, 如非函数或者 i 不在 [0, NumOut()) 内将会 panic")
	fmt.Println("-------------------------------")

	fmt.Println("reflect.FuncOf([]reflect.Type{reflect.TypeOf(0), reflect.TypeOf(123.456)}, []reflect.Type{reflect.TypeOf(\"\")}, false).String()")
	fn := reflect.FuncOf([]reflect.Type{reflect.TypeOf(0), reflect.TypeOf(123.456)}, []reflect.Type{reflect.TypeOf("")}, false)
	fmt.Println(fn.String()) // func(int, int) string
	fmt.Println("-------------------------------")

	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
	fmt.Println("-------------------------------")

	fmt.Println(`
var floatSwap func(float64, float64) (float64, float64)
fnRType := reflect.TypeOf(floatSwap)
in := fnRType.NumIn()
out := fnRType.NumOut()
fmt.Println("fnRType.NumIn()", in)
fmt.Println("fnRType.NumOut()", out)
for i := 0; i < in; i++ {
  fmt.Printf("fnRType.In(%\d).Name() %\v\n", i, fnRType.In(i).Name())
}
for i := 0; i < out; i++ {
  fmt.Printf("fnRType.Out(%\d).Name() %\v\n", i, fnRType.Out(i).Name())
}`)
	fmt.Println("-----------------")
	fnRType := reflect.TypeOf(floatSwap)
	in := fnRType.NumIn()
	out := fnRType.NumOut()
	fmt.Println("fnRType.NumIn()", in)   // 2
	fmt.Println("fnRType.NumOut()", out) // 2
	for i := 0; i < in; i++ {
		fmt.Printf("fnRType.In(%d).Name() %v\n", i, fnRType.In(i).Name()) // 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，Name() 会返回 ""
	}
	for i := 0; i < out; i++ {
		fmt.Printf("fnRType.Out(%d).Name() %v\n", i, fnRType.Out(i).Name()) // 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，Name() 会返回 ""
	}
	fmt.Println("-------------------------------")
}
