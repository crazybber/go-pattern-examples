package flyweight

/*
 * @Description: github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-01 18:10:57
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-01 19:38:21
 */
import (
	"fmt"
)

//IDeliverCompany 公司的能力
type IDeliverCompany interface {
	//雇人
	Hire(name string)
	//送货任务
	DeliverTask(name string, packets []string)

	GetDeliver(name string) IDeliver
}

//DeliverCompany 快递公司
type DeliverCompany struct {
	Employees map[string]IDeliver
}

//IDeliver 快递员能做的事情
type IDeliver interface {
	//送货
	DeliverPackets(packets []string)
}

//Deliver 快递员工,员工是一个享元对象
type Deliver struct {
	Name    string   //快递员的名字
	Packets []string //快递员的携带的快递，这一部分是变化的
}

//Hire 雇佣新员工，员工是全公司共享
func (d *DeliverCompany) Hire(name string) {
	if d.Employees == nil || len(d.Employees) == 0 {
		d.Employees = make(map[string]IDeliver)
	}
	if _, ok := d.Employees[name]; ok {
		fmt.Println("already hired")
		return
	}
	d.Employees[name] = &Deliver{Name: name}

	fmt.Println("hired")
}

//GetDeliver return Deliver
func (d *DeliverCompany) GetDeliver(name string) IDeliver {

	return d.Employees[name]

}

//DeliverTask 派员工送货
func (d *DeliverCompany) DeliverTask(name string, packets []string) {

	d.Employees[name].DeliverPackets(packets)

}

//DeliverPackets 送货了
func (d *Deliver) DeliverPackets(packets []string) {

	fmt.Println(d.Name, ": Delivered:", packets)
}
