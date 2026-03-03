package factoryfunc

import (
	"fmt"
	"testing"
)

func TestDoc(t *testing.T) {
	t.Run("doc", func(t *testing.T) {
		fmt.Println("工厂方法模式: 定义一个创建对象的接口, 但由实现这个接口的结构体决定实例化哪个结构体, 工厂方法把结构体的实例化推迟到子结构体中进行.")
		fmt.Println("优点:")
		fmt.Println("  1. 灵活性增强, 对于新产品的创建, 只需要多写一个相应的结构体工厂")
		fmt.Println("  2. 典型的解耦框架, 高层只需要知道产品的抽象类,无须关心其他实现类, 满足LKP(狄米特法则), DIP(依赖倒置原则), LSP(里氏替换原则)")
		fmt.Println("缺点:")
		fmt.Println("  1. 结构体的个数容易过多, 增加复杂度")
		fmt.Println("  2. 增加了系统的抽象性和理解难度")
		fmt.Println("  3. 抽象产品只能生产一种产品, 此弊端可使用抽象工厂模式解决")
	})
}
func TestFactoryFunc(t *testing.T) {
	t.Run("factoryFunc", func(t *testing.T) {
		/*
			var cf ShapeFactory = CircleFactory{}
			cannot use CircleFactory{} (type CircleFactory) as type ShapeFactory in assignment:
			 CircleFactory does not implement ShapeFactory (Create method has pointer receiver)
		*/
		var cf ShapeFactory = &CircleFactory{}    // *factoryfunc.CircleFactory &factoryfunc.CircleFactory{}
		var sf ShapeFactory = &SquareFactory{}    // *factoryfunc.SquareFactory &factoryfunc.SquareFactory{}
		var rf ShapeFactory = &RectangleFactory{} // *factoryfunc.RectangleFactory &factoryfunc.RectangleFactory{}
		fmt.Println("-------------------------")
		cf.Create().Draw()
		sf.Create().Draw()
		rf.Create().Draw()
		fmt.Println("-------------------------")
		var c1 Shape = cf.Create()
		var c2 Shape = cf.Create()
		fmt.Printf("c1 的类型为 %T 地址为 %p 值为 %#v\n", c1, c1, c1) // *factoryfunc.Circle &factoryfunc.Circle{Name:"circle"}
		fmt.Printf("c2 的类型为 %T 地址为 %p 值为 %#v\n", c2, c2, c2) // *factoryfunc.Circle &factoryfunc.Circle{Name:"circle"}
		c1.Draw()
		c2.Draw()
		fmt.Println("-------------------------")
		var s1 Shape = sf.Create()
		var s2 Shape = sf.Create()
		fmt.Printf("s1 的类型为 %T 地址为 %p 值为 %#v\n", s1, s1, s1) // *factoryfunc.Square &factoryfunc.Square{Name:"square"}
		fmt.Printf("s2 的类型为 %T 地址为 %p 值为 %#v\n", s2, s2, s2) // *factoryfunc.Square &factoryfunc.Square{Name:"square"}
		s1.Draw()
		s2.Draw()
		fmt.Println("-------------------------")
		var r1 Shape = rf.Create()
		var r2 Shape = rf.Create()
		fmt.Printf("r1 的类型为 %T 地址为 %p 值为 %#v\n", r1, r1, r1) // *factoryfunc.Rectangle &factoryfunc.Rectangle{Name:"rectangle"}
		fmt.Printf("r2 的类型为 %T 地址为 %p 值为 %#v\n", r2, r2, r2) // *factoryfunc.Rectangle &factoryfunc.Rectangle{Name:"rectangle"}
		r1.Draw()
		r2.Draw()
		fmt.Println("-------------------------")
	})
}
