package iterator

////////////////////////////////
//使用景点例子

//ITouristMap 游客地图是一个迭代器Iterator，为用户提供当前景区中不同景点的统一访问能力
type ITouristMap interface {
	FirstPot()         //首个景点
	IsLastPot() bool   //当前景点是否是最后一个
	Next() interface{} //下一个景点
}

//IScenicArea 是一个针对景区的Aggregate聚合类型接口，返回一个迭代器接口
//IScenicArea 返回一个游客访问接口
type IScenicArea interface {
	Iterator() ITouristMap
}

//ScenicArea 景区包含所有的景点
type ScenicArea struct {
	count int           //景点的数量
	pots  []interface{} //景点列表，景区可能一直在开发新的景点，所以景区的数量可能一直在增长
}

//Iterator 通过
func (s *ScenicArea) Iterator() ITouristMap {
	return &ScenicAreaPotsMap{
		numbers: n,
		next:    n.start,
	}
}

//ScenicAreaPotsMap 就是景区提供的迭代器类型，要实现具体的景区景点的迭代访问能力
type ScenicAreaPotsMap struct {
	numbers *Numbers
	next    int
}

//FirstPot 第一个景点
func (i *ScenicAreaPotsMap) FirstPot() {
	i.next = i.numbers.start
}

//IsLastPot 是否是最后一个
func (i *ScenicAreaPotsMap) IsLastPot() bool {
	return i.next > i.numbers.end
}

//Next 去路线上的下一个景点
func (i *ScenicAreaPotsMap) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}
