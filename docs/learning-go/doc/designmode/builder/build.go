package builder

/*
构建者模式将复杂对象的构建与其表示分离开来, 允许同一构建过程创建不同的表示,
它能解决问题: 复杂的对象通常是一步一步构建的，构建者模式为创建此类对象提供了灵活的解决方案
*/
type HouseBuilder interface {
	SetWindows() HouseBuilder
	SetDoors() HouseBuilder
	SetRoof() HouseBuilder
	Build() *House
}
type House struct {
	windows string
	doors   string
	roof    string
}
type VillaBuilder struct {
	house House
}

func (v *VillaBuilder) SetWindows() HouseBuilder {
	v.house.windows = "villa house"
	return v
}
func (v *VillaBuilder) SetDoors() HouseBuilder {
	v.house.doors = "villa doors"
	return v
}
func (v *VillaBuilder) SetRoof() HouseBuilder {
	v.house.roof = "villa roof"
	return v
}
func (v *VillaBuilder) Build() *House {
	return &v.house
}

type Director struct {
	builder HouseBuilder
}

func (d *Director) Construct() *House {
	return d.builder.SetWindows().SetDoors().SetRoof().Build()
}
