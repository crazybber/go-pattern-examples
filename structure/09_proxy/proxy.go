package proxy

import "errors"

////////////////////////////////
//用一个代理注册公司的例子，正向代理
////////////////////////////////

//Nation 国籍
type Nation int

//nation catalog
const (
	CN Nation = iota
	UK
	JP
)

//IPioneer 一个创业者，要注册公司
type IPioneer interface {
	RegisterCompany(companyName string) (name, enterpriseNo string, err error)

	Conndition() (money int, kind Nation)
}

//Pioneer 一个创业者
type Pioneer struct {
	AccountMoney int
	NationKind   Nation
}

//RegisterCompany 创业者要注册公司
func (p Pioneer) RegisterCompany(companyName string) (name, enterpriseNo string, err error) {

	return
}

//Conndition 注册条件
func (p Pioneer) Conndition() (money int, kind Nation) {
	money = p.AccountMoney
	kind = p.NationKind
	return
}

//RegistryProxyCompany 代注公司，帮用户注册，对用户来讲，可以当成工商局来看待.
type RegistryProxyCompany struct {
	p IPioneer
}

//RegisterCompany 代表用户注册公司
func (r RegistryProxyCompany) RegisterCompany(companyName string) (name, enterpriseNo string, err error) {

	//检查注册人的，资金,姓名，
	money, nation := r.p.Conndition()

	if money < 10000 || nation != CN {
		return "", "", errors.New("Condition not OK")
	}
	////////////////////////////////
	///发送请求到工商局
	////////////////////////////////

	name = companyName
	enterpriseNo = "abvdefe12450"

	return name, enterpriseNo, err

}
