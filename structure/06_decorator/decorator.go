package decorator

////////////////////////////////
//拿化妆做例子

//IFaceLooks 脸部的颜值
type IFaceLooks interface {
	FaceLooks() int
}

//NatureGirl 天然无化妆的小姐姐
type NatureGirl struct {
	faceValue int
}

//FaceLooks 获取小姐姐的颜值
func (n *NatureGirl) FaceLooks() int {
	return n.faceValue
}

//GirlWithMakeups 化妆后的小姐姐
type GirlWithMakeups struct {
	origin   IFaceLooks //这就是那个自然美的小姐姐
	facePlus int        //脸部加成，说，你想化成什么样子吧。
}

//NewGirlWithMakeup 返回一个化妆后的小姐姐
func NewGirlWithMakeup(origin IFaceLooks, facePlus int) IFaceLooks {
	return &GirlWithMakeups{
		origin:   origin,
		facePlus: facePlus,
	}
}

//FaceLooks 我要开始化妆了..
func (g *GirlWithMakeups) FaceLooks() int {
	return g.origin.FaceLooks() + g.facePlus
}

//FaceReal 实际的颜值..
func (g *GirlWithMakeups) FaceReal() int {
	return g.origin.FaceLooks()
}
