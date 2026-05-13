package factoryfunc

import (
	"fmt"
)

// 接口 Shape
type Shape interface {
	Draw()
}

// 结构体 Circle
type Circle struct {
	Name string
}

// 结构体 Circle 实现 Shape 接口
func (c *Circle) Draw() {
	fmt.Println(c.Name, "Draw()...")
}

// 结构体 Square
type Square struct {
	Name string
}

// 结构体 Square 实现 Shape 接口
func (s *Square) Draw() {
	fmt.Println(s.Name, "Draw()...")
}

// 结构体 Rectangle
type Rectangle struct {
	Name string
}

// 结构体 Rectangle 实现 Shape 接口
func (r *Rectangle) Draw() {
	fmt.Println(r.Name, "Draw()...")
}

// 重点, 所有的结构体子工厂都实现此接口, Create 方法创建一个结构体对象
// 工厂方法接口 ShapeFactory
type ShapeFactory interface {
	Create() Shape
}

// 结构体 Circle 工厂 CircleFactory
type CircleFactory struct{}

// 结构体 Circle 工厂 CircleFactory 实现 ShapeFactory 接口
func (c *CircleFactory) Create() Shape {
	// 接口为引用类型, 如果使用链式调用方法, 此处则需要返回指针类型
	/*
		return Circle{Name: "circle"}
		cannot use Circle{...} (type Circle) as type Shape in return argument:
			Circle does not implement Shape (Draw method has pointer receiver)
	*/
	return &Circle{Name: "circle"}
}

// 结构体 Square 工厂 SquareFactory
type SquareFactory struct{}

// 结构体 Square 工厂 SquareFactory 实现 ShapeFactory 接口
func (s *SquareFactory) Create() Shape {
	return &Square{Name: "square"}
}

// 结构体 Rectangle 工厂 RectangleFactory
type RectangleFactory struct{}

// 结构体 Rectangle 工厂 RectangleFactory 实现 ShapeFactory 接口
func (r *RectangleFactory) Create() Shape {
	return &Rectangle{Name: "rectangle"}
}
