package mediator

import (
	"fmt"
)

//在这个模式里面，我们演示一个场景用例
//房东通过中介向房客收房租
//房客通过中介向房东要求换新家具
//中介每次服务都收100块服务费

//Person 定义一个本身的人
type Person struct {
	Name         string
	WalletAssets int //每个人都有钱包
}

//Tenant 租客 继承Person
type Tenant struct {
	Person
	furniture string
}

//ITenant 租户能做的事情
type ITenant interface {
	AskRepair(mediator IMediator)
}

//Landlord 房东，也继承Person，要收房租
type Landlord struct {
	Person
	RentAccout int //房东的租金账户
}

//ILandlord 房东能做的事情
type ILandlord interface {
	CollectRent(mediator IMediator)
}

//Mediator 中介也继承Person，比如某居客，某家，某壳，即代表租客跟房东谈条件，又代表房东对付租客
//Mediator 中介一定会持有两方的必要信息
//Mediator 这里简化一下，假设中介只为一个房东和一个租客服务
type Mediator struct {
	Person
	tenant      ITenant   //中介持有房客的信息
	landlord    ILandlord //中介持有房东的信息
	feelandlord int
	feetenant   int
}

//IMediator 中介能做的事情,中介能代表任何一方，
//所以理论上他需要实现所代表对象的所有能力
//实际设计中，中介对象本身也会成为问题的所在，可能会比较臃肿
type IMediator interface {
	RegisterRoom(landlord ILandlord)
	Serve(client interface{}) //服务日常活动
	RentOutRoom(tenant ITenant)
}

//AskRepair 要求房东修家具，只需要向中介提要求，中介会提代替房客提要求
func (t *Tenant) AskRepair(mediator IMediator) {
	fmt.Println("Tenant: i need landlord fix furniture:")
	mediator.Serve(t)
}

//CollectRent 房东收租金,只需要向中介收，中介会提代替房东收租金
func (l *Landlord) CollectRent(mediator IMediator) {
	fmt.Println("Landlord: collect money")
	fmt.Printf("Landlord: RentAccout %d, WalletAssets %d\n", l.RentAccout, l.WalletAssets)
	mediator.Serve(l)

}

//RegisterRoom 可以在中介这里发布房源
func (m *Mediator) RegisterRoom(landlord ILandlord) {
	m.landlord = landlord
}

//RentOutRoom 可以从中介租房子
func (m *Mediator) RentOutRoom(tenant ITenant) {
	m.tenant = tenant
}

//Serve 中介要替两边或者多边办事，所以它很累,所有事情都要做
//这是关键过程
//简单起见，1代表租客，2代表房东
func (m *Mediator) Serve(client interface{}) {

	switch client.(type) {
	case ITenant:
		fmt.Println("i am serving tenant")
	case ILandlord:
		fmt.Println("i am serving landlord")
	}

}
