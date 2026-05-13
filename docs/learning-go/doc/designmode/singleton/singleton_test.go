package singleton_test

import (
	"fmt"
	"testing"

	"github.com/luoleiself/learning-go/designmode/singleton"
)

func TestDoc(t *testing.T) {
	t.Run("singleTon", func(t *testing.T) {
		fmt.Println("包外测试: ------")
		fmt.Println("单例模式: 确保一个类在任何情况下都绝对只有一个实例, 并提供一个全局访问点, 属于创建型设计模式.")
	})
}

func TestSingleTon(t *testing.T) {
	t.Run("singleTon", func(t *testing.T) {
		o1 := singleton.GetInstance()
		fmt.Printf("o1 的类型为 %T 地址为 %p 值为 %#v\n", o1, o1, o1)
		o2 := singleton.GetInstance()
		fmt.Printf("o2 的类型为 %T 地址为 %p 值为 %#v\n", o2, o2, o2)
		o3 := singleton.GetInstance()
		fmt.Printf("o3 的类型为 %T 地址为 %p 值为 %#v\n", o3, o3, o3)
	})
}
