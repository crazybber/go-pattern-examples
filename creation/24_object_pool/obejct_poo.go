package objectpool

type doctor struct {
	name string
	kind int //科室
}

type pool chan *doctor

func newPool(total int) pool {
	p := make(pool, total)

	for i := 0; i < total; i++ {
		p <- new(doctor)
	}

	return p
}

//surgery
func (d doctor) surgery() {

}
