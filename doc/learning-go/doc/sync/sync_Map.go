package main

import (
	"fmt"
	"sync"
)

func syncMapNote() {
	fmt.Println("-------------syncMapNote()-------------")
	fmt.Println("sync.Map(需要先导入 sync 包) 并发安全的 map 不需要初始化可直接声明使用, 操作时需要通过指定方法 Store 保存, Load 获取, Delete 删除, Range 遍历, LoadOrStore 获取或者保存, LoadAndDelete 获取并删除指定键值对")
	var sm sync.Map
	sm.Store("green", 97)
	sm.Store("london", 100)
	sm.Store(1, "1")
	fmt.Println(tab, "var sm sync.Map")
	fmt.Println(tab, "sm.Store(\"green\", 97)")
	fmt.Println(tab, "sm.Store(\"london\", 100)")
	fmt.Println(tab, "sm.Store(1, \"1\")")
	// fmt.Println("sm =", sm) // {{0 0} {{map[] true}} map[1:0xc000110018 green:0xc000110008 london:0xc000110010] 0}
	value, ok := sm.Load("london")
	fmt.Printf("sm.Load(key interface{}) (value interface{}, ok bool) 获取指定key的值 value, ok := sm.Load(\"london\") 操作 %t 值为 %v\n", ok, value) // true 100
	sm.Delete("london")
	// fmt.Printf("sm.Delete(key interface{}) 删除指定的键值对 sm.Delete(\"london\") 后 map 为 %v\n", sm) // {{0 0} {{map[1:0xc000110018 green:0xc000110008] false}} map[] 0}
	fmt.Println("---------------")

	fmt.Println("sync.Map 获取 map 数量需要在遍历时自己统计")
	fmt.Println(`sm.Range(f func(key, value interface{}) bool) 遍历方法接收一个回调函数, 回调函数的第一个参数 key, 第二个参数 value, 回调函数的返回值 bool 决定是否继续遍历
  count := 0
  sm.Range(func(k, v interface{}) bool {
    count += 1 // 手动统计 map 长度
    fmt.Println("iterate:", "k=", k, " v=", v)
    return true // 返回值决定是否继续遍历
  })`)
	count := 0
	sm.Range(func(k, v interface{}) bool {
		count += 1 // 手动统计 map 长度
		fmt.Println("iterate:", "k=", k, " v=", v)
		return true // 返回值决定是否继续遍历
	})
	fmt.Printf("sm 的长度为 %d\n", count) // 2
	fmt.Println("---------------")

	res1, ok := sm.LoadOrStore("blue", "BLUE")
	fmt.Println("sm.LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) 获取或者保存键值对, 如果已存在则返回 已获取的值 和 true, 不存在则保存当前键值对并返回 当前值 和 false")
	fmt.Printf("sm.LoadOrStore(\"blue\", \"BLUE\") 的返回值 %+v %v\n", res1, ok) // BLUE false
	fmt.Print(`sm.LoadOrStore("blue", "BLUE") `)
	fmt.Println(sm.LoadOrStore("blue", "BLUE")) // BLUE true
	// fmt.Println("sm =", sm)                     // {{map[1:0xc000110018 green:0xc000110008] true}} map[1:0xc000110018 blue:0xc000110020 green:0xc000110008] 1}
	fmt.Println("---------------")

	res2, ok := sm.LoadAndDelete("blue")
	fmt.Println("sm.LoadAndDelete(key interface{}) (value interface{}, loaded bool) 获取并删除指定键值对, 如果已存在删除指定键值对并返回 已获取的值 和 true, 不存在返回 nil 和 false")
	fmt.Printf("sm.LoadAndDelete(\"blue\") 的返回值 %v %v\n", res2, ok) // BLUE true
	fmt.Print(`sm.LoadAndDelete("blue") `)
	fmt.Println(sm.LoadAndDelete("blue")) // <nil> false
	// fmt.Println("sm =", sm)               // {{0 0} {{map[1:0xc000110018 green:0xc000110008] false}} map[] 0}
	fmt.Println("------------------------------")
}
