package main

import (
	"fmt"
	"sort"
)

func CustomSort() {
	fmt.Println("自定义排序")
	fmt.Println(`// 定义结构体
  type Person struct {
    Name string
    Age  uint8
  }

  // 定义类型别名
  type ByAge []Person

  // 任何数据类型只要实现了这三个方法, 默认就实现了 sort.Interface 接口, 就可以被 sort 包的函数进行排序, 方法要求集合中的元素可以被整数索引
  func (b ByAge) Len() int {
    return len(b) // Len方法返回集合中的元素个数
  }
  func (b ByAge) Less(i, j int) bool {
    return b[i].Age < b[j].Age // Less方法报告索引i的元素是否比索引j的元素小
  }
  func (b ByAge) Swap(i, j int) {
    b[i], b[j] = b[j], b[i] // Swap方法交换索引i和j的两个元素
  }

  person := []Person{
    {"Bob", 31},
    {"John", 42},
    {"Michael", 17},
    {"Jenny", 26},
  }
  fmt.Println(person)
  sort.Sort(ByAge(person)) // person 类型为 person 切片, 需要转换为 ByAge 类型
  fmt.Println(person)`)
	fmt.Println("-------------")
	person := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(person)
	sort.Sort(ByAge(person)) // person 类型为 person 切片, 需要转换为 ByAge 类型
	fmt.Println(person)
	fmt.Println("-----------------------------")
}

// 定义结构体
type Person struct {
	Name string
	Age  uint8
}

// 定义类型别名
type ByAge []Person

// 任何数据类型只要实现了这三个方法, 默认就实现了 sort.Interface 接口, 就可以被 sort 包的函数进行排序, 方法要求集合中的元素可以被整数索引
func (b ByAge) Len() int {
	return len(b) // Len方法返回集合中的元素个数
}
func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age // Less方法报告索引i的元素是否比索引j的元素小
}
func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i] // Swap方法交换索引i和j的两个元素
}
