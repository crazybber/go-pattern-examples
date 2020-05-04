package iterator

////////////////////////////////
//使用景点例子

//Iterator 游客地图是一个迭代器Iterator，为用户提供当前景区中不同景点的统一访问能力
type Iterator interface {
	Reset()          //重置
	FirstPot() IPot  //首个景点
	IsLastPot() bool //当前景点是否是最后一个
	Next() IPot      //下一个景点
}

//IPot 景点对象的接口
type IPot interface {
	Visit() //景点可以参观
}

//ScenicArea 景区包含所有的景点
type ScenicArea struct {
	count int    //景点的数量
	pots  []IPot //景点列表，景区可能一直在开发新的景点，所以景区的数量可能一直在增长
}

//PotsIterator 实现景区景点迭代器的对象
//PotsIterator 该对象的目的就是为了隐藏景区本身
//PotsIterator 实现为一个游标迭代器
type PotsIterator struct {
	cursor, count int
	potsSlice     []IPot
}

//Iterator 返回一个接待
func (s *ScenicArea) Iterator() Iterator {
	return &PotsIterator{
		cursor:    0,
		count:     s.count,
		potsSlice: s.pots,
	}
}

//AddPot 添加景点
func (s *ScenicArea) AddPot(pots ...IPot) {
	for i := range pots {
		s.count++
		s.pots = append(s.pots, pots[i])
	}
}

//PotsCount 添加景点
func (s *ScenicArea) PotsCount() int {
	return s.count
}

//Reset 重置
func (s *PotsIterator) Reset() {
	s.cursor = 0
}

//FirstPot 第一个景点
func (s *PotsIterator) FirstPot() IPot {
	return s.potsSlice[0]
}

//IsLastPot 判断游标的位置
func (s *PotsIterator) IsLastPot() bool {
	return s.cursor == s.count
}

//Next 去路线上的下一个景点
func (s *PotsIterator) Next() IPot {
	s.cursor++
	if s.IsLastPot() {
		return nil
	}
	return s.potsSlice[s.cursor]

}
