package prototype

//Cloneable is the key interface
type Cloneable interface {
	Clone() Cloneable
}

//cloneLab 克隆实验室,可以克隆很多动物
type cloneLab struct {
	animals map[string]Cloneable
}

//new 返回一个
func newCloneLab() *cloneLab {
	return &cloneLab{
		animals: make(map[string]Cloneable),
	}
}

func (c *cloneLab) Get(name string) Cloneable {
	return c.animals[name]
}

func (c *cloneLab) Set(name string, newObject Cloneable) {
	c.animals[name] = newObject
}
