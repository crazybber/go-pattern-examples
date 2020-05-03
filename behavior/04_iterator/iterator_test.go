package iterator

import (
	"fmt"
	"testing"
)

//ChildPot 儿童景点
type ChildPot struct {
	Name string
}

func (c *ChildPot) Visit() {
	fmt.Println("i am: ", c.Name)
}
func TestIterator(t *testing.T) {

	scenicArea := ScenicArea{}

	scenicArea.AddPot(&ChildPot{Name: "monkey garden"}, &ChildPot{Name: "fairy country"}, &ChildPot{Name: "future space"})

	t.Log("pots count:", scenicArea.PotsCount())

	potInterator := scenicArea.Iterator()

	pot := potInterator.FirstPot()

	t.Logf("first pot: %#v\n", pot)

	VisitAllPots(potInterator)

	t.Log("add a pot", "pot: count", scenicArea.PotsCount())

	scenicArea.AddPot(&ChildPot{Name: "virtual world"})

	t.Log("pots count:", scenicArea.PotsCount())

	//切片变了，要重新获取
	potInterator = scenicArea.Iterator()

	potInterator.Reset()

	VisitAllPots(potInterator)

}

func VisitAllPots(i Iterator) {
	for c := i.FirstPot(); !i.IsLastPot(); c = i.Next() {
		c.Visit()
		fmt.Printf("finish visit pot : %#v\n", c)
	}
}
