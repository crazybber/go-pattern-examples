package prototype

import "testing"

//cloneLab 克隆实验室
var lab *cloneLab

type sheep struct {
	name   string
	weight int
}

func (s *sheep) Clone() Cloneable {
	tc := *s
	return &tc
}

type cow struct {
	name   string
	gender bool
}

func (c *cow) Clone() Cloneable {
	newCow := &cow{
		gender: c.gender,
		name:   c.name,
	}
	return newCow
}

func TestClone(t *testing.T) {

	sheep1 := &sheep{
		name:   "sheep",
		weight: 10,
	}

	sheep2 := sheep1.Clone()

	if sheep1 == sheep2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromLab(t *testing.T) {

	lab := newCloneLab()

	lab.Set("cow", &cow{name: "i am cow", gender: true})

	c := lab.Get("cow").Clone()

	cw := c.(*cow)
	if cw.name != "i am cow" {
		t.Fatal("error")
	}

}
