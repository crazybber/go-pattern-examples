package composite

import (
	"fmt"
	"reflect"
)

////////////////////////////////
//使用集装箱与货物的例子
//下面三个是对象的该模式的关键结构
////////////////////////////////

//Cargo class 基本的货物类型可以被继承
type Cargo struct {
	Volume      uint   //货物需要要尺寸
	Description string //描述
}

//GetCargoType return type
func (c *Cargo) GetCargoType() string {
	return reflect.TypeOf(c).String()
}

//ShowContent return type
func (c *Cargo) ShowContent() {
	typeName := reflect.TypeOf(c).String()
	fmt.Println("Type: ", typeName, " Content ", c.Description)

}

//Box 复合类型，表示集装箱
//Box 复合类型，集装箱里面装具体的货物,也可以继续放箱子
type Box struct {
	Cargo                    //继承货物类
	InnerSpace uint          //内部空间
	Children   []interface{} //有子对象
}

//PutInCargo 增加新的能力
//PutInCargo (cargo ICargo) //放入一个子类型
func (b *Box) PutInCargo(cargo interface{}) {

	switch cargo.(type) {
	case *Box:
		fmt.Println("get a Box: Type: ", cargo.(*Box).GetCargoType())
	case *SingleCargo:
		fmt.Println("get a SingleCargo Type: ", cargo.(*SingleCargo).GetCargoType())
	}

	b.Children = append(b.Children, cargo)

}

//GetChildren () []ICompositedCargo
func (b *Box) GetChildren() []interface{} {
	return b.Children
}

//ShowContent 覆盖继承实现
//ShowContent display children content
func (b *Box) ShowContent() {
	typeName := reflect.TypeOf(b).String()
	fmt.Println("Type: ", typeName, " InnerSpace ", b.InnerSpace, " Children: ", b.Description)
	count := len(b.Children)
	fmt.Println("Children Count: ", count, " Description ", b.Description)
	for _, child := range b.Children {
		//判断类型
		switch child.(type) {
		case *Box:
			fmt.Println("Current Child is a Box: Type cast to (*Box) ")
			child.(*Box).ShowContent()
		case *SingleCargo:
			fmt.Println("Current Child is a Single Cargo: Type cast to (*SingleCargo) ")
			child.(*SingleCargo).ShowContent()
		}
	}

}

//SingleCargo 具体的非复合类型，对应多叉树的叶子节点
type SingleCargo struct {
	Cargo           //继承货物类，具有基本的货物熟悉
	From, To string //货是从谁发到谁的
}

//ShowContent return type
func (s *SingleCargo) ShowContent() {
	typeName := reflect.TypeOf(s).String()
	fmt.Println("Type: ", typeName, " From ", s.From, " To ", s.To, " Content ", s.Description)

}
