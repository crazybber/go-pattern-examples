package objectpool

import (
	"testing"
)

func TestObjectPool(t *testing.T) {

	p := newPool(2)

	select {
	case obj := <-p:
		obj.surgery( /*...*/ )

		p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}
}
