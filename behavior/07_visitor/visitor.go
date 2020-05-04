package visitor

import "fmt"

////////////////////////////////
//使用石油的例子

//IGasResource 作为资源提供接口
type IGasResource interface {
	Accept(IGasVisitor)
}

//gas 汽油
type gas struct {
	density int
}

//IGasVisitor 访问者接口
type IGasVisitor interface {
	Visit(gas)
}

//Accept 接待汽油客户
func (g gas) Accept(visitor IGasVisitor) {
	visitor.Visit(g)
}

//diesel 柴油
type diesel struct {
	energy int
}

//IDieselVisitor 访问者接口
type IDieselVisitor interface {
	Visit(diesel)
}

//Accept 提供柴油
func (d diesel) Accept(visitor IDieselVisitor) {
	visitor.Visit(d)
}

//militaryFactory 军工厂，消费石油，制造务器
type militaryFactory struct {
	name string
}

//Visit 军工厂只够买柴油，制造武器
func (m *militaryFactory) Visit(d diesel) {
	fmt.Println("militaryFactory: use diesel with inner energy", d.energy)
}

// clothFactory 服务装类工厂，购买汽油，制造化纤物品
type clothFactory struct{}

//Visit 购买汽油
func (c *clothFactory) Visit(g gas) {
	fmt.Println("clothFactory: use gas with density", g.density)
}
