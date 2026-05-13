package main

import (
	"fmt"
)

type Student struct {
	Name     string             `json:"name"`
	Sex      byte               `json:"sex"`
	Age      uint8              `json:"age"`
	Address  string             `json:"addr"`
	Weight   float32            `json:"weight"`
	Hobbies  []string           `json:"hobbies"`
	Score    map[string]float32 `json:"score"`
	Avatar   any                `json:"avatar"`
	HomePage string             `json:"homePage"`
}

func JSONDemo() {
	fmt.Println("--------json---------")
	fmt.Println("func Marshal(v interface{}) ([]byte, error)")
	fmt.Println("func Unmarshal(data []byte, v interface{}) error")
	fmt.Println("`json:\"-\"` // 字段被忽略")
	fmt.Println("`json:\"myName,omitempty\"` // 使用 myName 作为键名, 如果字段为空值则跳过")
	fmt.Println("`json:\",omitempty\"` // 使用字段名作为键名, 如果字段为空值则跳过")
	fmt.Println("------")
	fmt.Println("多个匿名字段:")
	fmt.Println("  json 标签为 \"-\" 的匿名字段强行忽略, 不作考虑")
	fmt.Println("  json 标签提供了键名的匿名字段, 视为非匿名字段")
	fmt.Println("  其余字段中如果只有一个匿名字段, 则使用该字段")
	fmt.Println("  其余字段中如果有多个匿名字段, 但压平后不会出现冲突, 所有匿名字段压平")
	fmt.Println("  其余字段中如果有多个匿名字段, 但压平后出现冲突, 全部忽略, 不产生错误")
	fmt.Println("------")
	fmt.Println("布尔类型的值编码为布尔类型")
	fmt.Println("数值类型的值编码为数字类型")
	fmt.Println("数组和切片类型的值编码为 json 数组")
	fmt.Println("映射类型的值编码为 json 对象, 映射的键必须是字符串, 对象的键直接使用映射的键")
	fmt.Println("结构体类型的值编码为 json 对象")
	fmt.Println("指针类型的值编码为其指向的值(的 json 编码), nil 指针编码为 null")
	fmt.Println("接口类型的值编码为接口内保存的具体类型的值(的 json 编码), nil 接口编码为 null")
	fmt.Println("通道、复数、函数类型的值不能编码进 json 会返回 UnSupportedTypeError")
	fmt.Println("-----------------")
}
