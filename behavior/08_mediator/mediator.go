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
}

//Landlord 房东，要收房租
type Landlord struct {
	Person
	accout int //房东的租金账户
}

//mediator 中介，比如某居客，某家，某壳，即代表租客跟房东谈条件，又代表房东对付租客
//mediator  所以中介一定会持有两方的信息，最好用接口代表对象
//mediator  这里简化一下，直接用类型的引用，表示拥有关系
type mediator struct {
	tenant      *Tenant
	landlord    *Landlord
	feeLandlord int //向房东收费的账户
	feeTenant   int //向租客收费的账户
}

//AskRepair 要求房东修家具
func (c *tenant) AskRepair(furniture string) {
	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

//CollectRent 房东收租金了
func (l *landlord) CollectRent(moneyCount int) {

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}
