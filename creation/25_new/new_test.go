package newpattern

import (
	"fmt"
	"testing"
)
type homecat interface {
	sleep()
}

type blackCat struct {
	name string
}

func newBlackCat(name string) homecat{
	return &blackCat{name}
}


func (b blackCat) sleep() {
	fmt.Print( b.name + " is sleeping")
}

func TestNewMode(t *testing.T){

	cat := newBlackCat("pi")
	cat.sleep()
}