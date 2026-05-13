package main

import (
	"errors"
	"fmt"
)

func rangeNote() {
	fmt.Println("-----------rangeNote()------------")

	fmt.Println("for 循环的4种使用方式")
	fmt.Println("  1. for i :=1;i < len;i++{ /*...*/ }")
	fmt.Println("  2. for ;; { /*...*/ }")
	fmt.Println("  3. for i < len { /*...*/ }")
	fmt.Println("  4. for k,v := range T { /*...*/ }")
	fmt.Println("------------------------------")
	fmt.Println("for range 是 go 特有的一种迭代结构, 可以遍历数组、切片、字符串、map、通道(channel)")
	fmt.Println("val 始终为集合中对应索引的值的拷贝, 对它所做的修改不会影响到集合中原有的值")
	fmt.Println(`str := "he ll o"
  for i, val := range str {
    val = 'g'
    fmt.Printf("str= %\q i= %\d val= %\q\n", str, i, val)
  }

  slice := []int{1, 2, 3, 4, 5, 6}
  for i, val := range slice {
    val = 100
    fmt.Printf("slice= %\v i= %\d val= %\v\n", slice, i, val)
  }`)
	fmt.Println("----------------")
	str := "he ll o"
	for i, val := range str {
		val = 'g'
		fmt.Printf("str= %q i= %d val= %q\n", str, i, val)
	}
	fmt.Println("----------------")
	slice := []int{1, 2, 3, 4, 5, 6}
	for i, val := range slice {
		val = 100
		fmt.Printf("slice= %v i= %d val= %v\n", slice, i, val)
	}
	fmt.Println("------------------------------")

	fmt.Println("range 遍历数组、切片、字符串返回索引和值")
	fmt.Println("range 遍历 map 返回键和值")
	fmt.Println("range 遍历 通道(channel)返回通道内的值")
	fmt.Println("------------------------------")

	breakNote()
	fmt.Println("------------------------------")
}

func breakNote() {
	fmt.Println("-----------------breakNote()-------------------")
	fmt.Println("流程控制是每种编程语言控制逻辑走向和执行次序的重要部分 go 语言提供了 if、switch、for、goto、break、continue等流程控制关键字, goto 和 break 可以跳转到定义的 label 处")
	fmt.Println("continue 结束当前循环, 开始下一次循环, 仅限在 for 循环中使用, 后面可以跟标签名跳转到指定的标签名位置")
	fmt.Println("break 后面可以跟标签名中断任意层级循环, 标签名必须定义在 for、switch、select 的代码块上")
	fmt.Println("goto 后面可以跟标签名跳转到程序指定位置, 禁止跳过变量声明、跳转到内部或平行代码块")
	fmt.Println("continue")
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
	fmt.Println("-----------------------------")
	fmt.Println(`OuterLoop:
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                continue OuterLoop
            }
        }
    }`)
	fmt.Println("-----------------------------------------------")

JLoop: // 这里是 break 跳转的位置
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break JLoop
			}
		}
	}
	name := "hello go"
	fmt.Println("name= ", name)
	fmt.Println("-----------------------------")
	fmt.Println(`JLoop: // 这里是 break 跳转的位置
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            break JLoop
        }
    }
    name := "hello go"
    fmt.Println("name= ", name)`)
	fmt.Println("-----------------------------------------------")

	fmt.Println("goto 跳出循环")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i= %d j= %d \n", i, j)
			if i == 3 && j == 3 {
				goto gotoHere
			}
		}
	}
	return // 手动返回, 避免执行进入标签
gotoHere:
	fmt.Println("done")
	fmt.Println("-----------------------------")
	fmt.Println(`for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        fmt.Printf("i= %\d j= %\d \n", i, j)
        if i == 3 && j == 3 {
            goto gotoHere
        }
    }
}
return // 手动返回, 避免执行进入标签
gotoHere:
    fmt.Println("done")
	`)
	fmt.Println("-----------------------------------------------")

	fmt.Println("goto 集中处理错误")
	err := errors.New("errors.New method return")
	if err != nil {
		goto onExit
	}

	err = errors.New("errors.New method return 2")
	if err != nil {
		goto onExit
	}

	fmt.Println("done")
	return
onExit:
	fmt.Println("onExit ", err)
	fmt.Println("-----------------------------")
	fmt.Println(`
    err := errors.New("errors.New method return")
    if err != nil {
        goto onExit
    }

    err = errors.New("errors.New method return 2")
    if err != nil {
        goto onExit
    }

    fmt.Println("done")
    return
onExit:
    fmt.Println("onExit ", err)`)
	fmt.Println("-----------------------------------------------")

	var s *int
	fmt.Printf("s 的类型为 %T 值为 %v\n", s, s)
	fmt.Println(s)
	fmt.Println("s == nil ", s == nil)
	fmt.Println("-----------------------------")
}
