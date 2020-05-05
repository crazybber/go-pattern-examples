package objectpool

import (
	"fmt"
	"strconv"
)

type doctor struct {
	name string
	kind int //科室,1内科，2外科
}

type pool chan *doctor

func newPool(total int) pool {
	p := make(pool, total)

	for i := 0; i < total; i++ {

		dc := new(doctor)
		dc.name = "doctor: " + strconv.Itoa(i)
		p <- dc
	}

	return p
}

//surgery
func (d doctor) surgery(someone string) {
	fmt.Println("doctor:", d.name, "do surgery for", someone)
}
