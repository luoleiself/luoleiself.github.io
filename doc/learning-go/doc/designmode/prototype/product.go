package prototype

import "fmt"

/*
原型模式允许通过复制现有对象(原型)来创建新对象, 而不是从头开始创建.
当创建一个新对象的成本很高, 而现有对象又可以克隆重用时，原型模式很有用
*/
type Cloneable interface {
	Clone() Cloneable
}

type Product struct {
	name     string
	category string
}

func (p *Product) Clone() Cloneable {
	return &Product{name: p.name, category: p.category}
}

func (p *Product) SetName(name string) {
	p.name = name
}
func (p *Product) SetCategory(category string) {
	p.category = category
}

func (p *Product) GetDetails() string {
	return fmt.Sprintf("Product Name: %s, Category %s", p.name, p.category)
}
