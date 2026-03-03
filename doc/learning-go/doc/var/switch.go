package main

import (
	"fmt"
)

func switchNote() {
	fmt.Println("-----------switchNote()-------------")
	fmt.Println("switch 表达式不必须为字面量(能确定表达式的最终结果即可), 有且只有一个 default 分支")
	fmt.Println(tab, "case 使用多个值时中间使用 逗号 分割")
	fmt.Println(tab, "case 为逻辑表达式时, switch 不需要判断条件")
	fmt.Println(tab, "fallthrough 继续执行下一个 case 语句")
	fmt.Println(tab, "使用 .(type) 配合 switch 实现类型断言")
	fmt.Println("-----------------")
	fmt.Println(`
  // case 为逻辑表达式时, switch 不需要判断条件
  var r int = 11
  switch { // case 为表达式 switch 不需要判断条件
  case r > 10 && r < 20:
    fmt.Println(r)
  }`)
	fmt.Println("-----------------")
	fmt.Println(`case 多值情况 
  var a = "mum"
  switch a {
  case "mum", "daddy":
    fmt.Println("family")
  }`)
	fmt.Println("-------------------------------")

	fmt.Println(`使用 fallthrough 穿透下一个 case
  var s = "hello"
  switch {
  case s == "hello":
    fmt.Println("hello")
    fallthrough
  case s != "world":
    fmt.Println("world")
  }`)
	fmt.Println("-------------------------------")

	fmt.Println("使用 .(type) 配合 switch 实现类型断言")
	fmt.Println(`
  switch instance := inter.(type) {
  case Type1:
    fmt.Println("Type1 = ", Type1)
  case Type2:
    fmt.Println("Type2 = ", Type2)
  }`)
	fmt.Println("-------------------------------")

	selectNote()
	fmt.Println("-------------------------------")
}

func selectNote() {
	fmt.Println("-----------selectNote()-------------")
	fmt.Println("公平的选择: 如果有多个 case 同时就绪, select 会随机选择一个执行, 保证每个 case 被公平考虑, 防止某些 case 长期得不到执行")
	fmt.Println("阻塞与非阻塞: 如果所有的 case 都不可执行且没有 default 语句时, select 将会阻塞直到至少有一个 case 可执行")
	fmt.Println(tab, "如果定义了 default 语句, 当没有 case 可以执行时, select 会执行 default 语句")
	fmt.Println("执行后解除阻塞: 当某个 case 执行后, 与其对应的 channel 将不再阻塞, 如果该 channel 是发送操作, 那么等待该 channel 的数据的 goroutine 将被唤醒")
	fmt.Println(tab, "反之, 如果该 channel 是接收操作, 那么向该 channel 发送的数据的 goroutine 将被唤醒")
}
