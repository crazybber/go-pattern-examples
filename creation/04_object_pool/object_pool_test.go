package objectpool

import (
	"fmt"
	"testing"
)

func TestObjectPool(t *testing.T) {

	p := newPool(3)

	doc1 := <-p
	doc1.surgery("tom")

	doc2 := <-p
	doc2.surgery("rose")

	doc3 := <-p
	doc3.surgery("kate")

	select {
	case obj := <-p:
		obj.surgery("lily")
		p <- obj
	default:
		// No more objects left — retry later or fail
		fmt.Println("No more objects left, this moment")
		return
	}
}
