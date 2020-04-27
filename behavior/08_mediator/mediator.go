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
	Name       string
	MoneyCount int //每个人都有钱包
}

//Tenant 租客
type Tenant struct {
	Person
	m *mediator
}

//Landlord 房东，要收房租
type Landlord struct {
	Person
	accout int //房东的租金账户
	m      *mediator
}

//Mediator 中介，比如某居客，某家，某壳，即代表租客跟房东谈条件，又代表房东对付租客
//Mediator  所以中介一定会持有两方的信息，最好用接口代表对象
//Mediator  这里简化一下，假设中介只为一个房东和一个租客服务，直接用类型的引用，表示拥有关系
type Mediator struct {
	tenant        interface{}
	landlord      interface{}
	feelandlord   int
	feelandtenant int
}

//AskRepair 要求房东修家具
func (t *Tenant) AskRepair(furniture string) {
	fmt.Println("i need landlord fix the:", furniture)
	t.m.Changed()
}

//CollectRent 房东收租金了
func (l *Landlord) CollectRent(moneyCount int) {

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
}

//Changed 中介要提两边或者多边办事，所以它
func (m *Mediator) Changed(i interface{}) {

}

//PublishRoom 可以在中介这里发布房源
func (m *Mediator) PublishRoom(landlord interface{}) {

	m.landlord = landlord

}

//RentRom 可以从中介租房子
func (m *Mediator) RentRom(tenant interface{}) {

	m.tenant = tenant
}
