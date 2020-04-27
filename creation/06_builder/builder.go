package builder

//ICar 汽车，我们要造车了
//ICar 车具有以下能力
type ICar interface {
	Start()
	Stop()
	Speed() int
	Brand() string
}

//ICarBuilder 用来造车
type ICarBuilder interface {
	Wheel(wheel int) ICarBuilder
	Engine(engine string) ICarBuilder
	Speed(max int) ICarBuilder
	Brand(brand string) ICarBuilder
	Build() ICar
}

//CarProto 车的原型
type CarProto struct {
	Wheel    int
	Engine   string
	MaxSpeed int
	Brand    string
}

//CarStudio 打算通过成立造车实验室进行造车
type CarStudio struct {
	prototype CarProto
}

// NewCarStudio 造车工作室
func NewCarStudio() ICarBuilder {
	return &CarStudio{}
}

// Wheel of car
func (c *CarStudio) Wheel(wheel int) ICarBuilder {
	c.prototype.Wheel = wheel
	return c
}

// Engine of car
func (c *CarStudio) Engine(engine string) ICarBuilder {
	return c
}

// Speed of car
func (c *CarStudio) Speed(max int) ICarBuilder {
	return c
}

// Brand of car
func (c *CarStudio) Brand(brand string) ICarBuilder {
	return c
}

// Build return a car
func (c *CarStudio) Build() ICar {
	return c.prototype
}
