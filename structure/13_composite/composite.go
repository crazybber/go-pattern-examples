package composite

//ItemType 具体物品的类型
type ItemType int

//定义两种数据类型，叶子和复合节点
const (
	ManCloth ItemType = iota
	WomanCloth
	HatOrShose
)

//IWearingItem 定义复合类型的接口,表示穿戴物品
type IWearingItem interface {
	GetWearingType() ItemType
}

//CompositedComponet 复合类型，复合类型知道所有的具体类型
type CompositedComponet struct {
}

//NewComponent return a concrate
func NewComponent(kind int, name string) IWearingItem {

	return nil
}
