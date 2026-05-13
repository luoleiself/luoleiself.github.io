package singleton

import (
	"fmt"
	"testing"
)

func TestSingleTon(t *testing.T) {
	t.Run("包内测试: ------", func(t *testing.T) {
		i1 := newSingleTon()
		fmt.Printf("i1 的类型为 %T 地址为 %p 值为 %#v\n", i1, i1, i1)
		i2 := newSingleTon()
		fmt.Printf("i2 的类型为 %T 地址为 %p 值为 %#v\n", i2, i2, i2)
		i3 := newSingleTon()
		fmt.Printf("i3 的类型为 %T 地址为 %p 值为 %#v\n", i3, i3, i3)
	})
}
