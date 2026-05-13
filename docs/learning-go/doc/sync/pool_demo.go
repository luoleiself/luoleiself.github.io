package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"unsafe"
)

func PoolNote() {
	fmt.Println("-------------PoolNote()-------------")
	fmt.Println(tab, `type Pool struct {
    // New optionally specifies a function to generate
    // a value when Get would otherwise return nil.
    // It may not be changed concurrently with calls to Get.
    New func() interface{}
    // contains filtered or unexported fields
  }`)
	fmt.Println(tab, "func (p *Pool) Get() interface{} // 从池中选择任意一个对象, 删除其在池中的引用计数, 并提供给调用者")
	fmt.Println(tab, "func (p *Pool) Put(x interface{}) // 将 x 放入池中")
	fmt.Println("------------------------------------")

	fmt.Println(`
  var StudentPool = sync.Pool{
    New: func() interface{} {
      // return &Student{}
      return new(Student)
    },
  }
  var s1 = StudentPool.Get().(*Student) // 类型断言
  defer StudentPool.Put(s1)
  s1.Name = "hello world"
  s1.Age = 18`)
	fmt.Println("------------------")
	var s1 = StudentPool.Get().(*Student) // 类型断言
	defer StudentPool.Put(s1)
	s1.Name = "hello world"
	s1.Age = 18
	// *main.Student 0xc00000a038	8	8	&main.Student{Name:"hello world", Age:0x12}
	fmt.Printf("s1 的类型为 %T 地址为 %p 占用内存大小为 %d 内存对齐系数为 %d 值为 %#v\n", s1, &s1, unsafe.Sizeof(s1), unsafe.Alignof(s1), s1)
	var s2 = StudentPool.Get().(*Student) // 类型断言
	defer StudentPool.Put(s2)
	s2.Name = "hello golang"
	s2.Age = 18
	// *main.Student 0xc00000a040	8	8	&main.Student{Name:"hello golang", Age:0x12}
	fmt.Printf("s2 的类型为 %T 地址为 %p 占用内存大小为 %d 内存对齐系数为 %d 值为 %#v\n", s2, &s2, unsafe.Sizeof(s2), unsafe.Alignof(s2), s2)

	fmt.Println("------------------------------------")

	fmt.Println(`
  var bufPool = sync.Pool{
    New: func() interface{} {
      return new(bytes.Buffer)
    },
  }
  bf, _ := bufPool.Get().(*bytes.Buffer)
  bf.Write([]byte("hello"))
  bf.WriteByte(' ')
  bf.Write([]byte("world"))
  bf.WriteString("=")
  bf.WriteRune('中')
  bf.WriteByte('\n')
  bf.WriteTo(os.Stdout)`)
	fmt.Println("------------------")
	var bufPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	bf, _ := bufPool.Get().(*bytes.Buffer)
	bf.Write([]byte("hello"))
	bf.WriteByte(' ')
	bf.Write([]byte("world"))
	bf.WriteString("=")
	bf.WriteRune('中')
	bf.WriteByte('\n')
	bf.WriteTo(os.Stdout)
	fmt.Println("------------------------------------")
}

type Student struct {
	Name string
	Age  uint8
}

var StudentPool = sync.Pool{
	// 可选参数 New 指定一个函数在 Get 方法可能返回 nil 时来生成一个值
	// 该参数不能在调用 Get 方法时被修改
	New: func() interface{} {
		// return &Student{}
		return new(Student)
	},
	// 包含隐藏或非导出字段
}
