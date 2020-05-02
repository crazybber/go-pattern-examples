package composite

import "testing"

//我们目的是:使用一致性的方式访问一类具有共性并相互组合起来的对象

func TestCompositeconnections(t *testing.T) {

	box1 := &Box{Cargo: Cargo{1, "Big Box"}, InnerSpace: 130000}

	box2 := &Box{Cargo{2, "Middle Box"}, 80000, nil}

	box1.PutInCargo(box2)

	cargo1 := &SingleCargo{Cargo{3, "Hat"}, "CN", "CA"}

	box1.PutInCargo(cargo1)

	cargo2 := &SingleCargo{Cargo: Cargo{4, "Men Cloth"}, From: "China", To: "UK"}

	cargo3 := &SingleCargo{Cargo: Cargo{5, "Women Cloth"}, From: "HK", To: "TW"}

	box2.PutInCargo(cargo2)

	box2.PutInCargo(cargo3)

	box1.ShowContent()

	//Box1
	// -Box2
	//   -Cargo 2
	//   -Cargo 3
	// -Cargo1

}

func TestComposite(t *testing.T) {

	box1 := Box{Cargo: Cargo{1, "Big Box"}, InnerSpace: 130000}
	box1.ShowContent()

	box2 := Box{Cargo{2, "Middle Box"}, 80000, nil}

	box2.ShowContent()

	cargo1 := SingleCargo{Cargo{1, "Hat"}, "CN", "CA"}

	cargo1.ShowContent()

	cargo2 := SingleCargo{Cargo: Cargo{2, "Men Cloth"}, From: "China", To: "UK"}

	cargo2.ShowContent()

	cargo3 := &SingleCargo{Cargo: Cargo{2, "Women Cloth"}, From: "HK", To: "TW"}

	cargo3.ShowContent()

	//Call base class
	cargo3.Cargo.ShowContent()
}
