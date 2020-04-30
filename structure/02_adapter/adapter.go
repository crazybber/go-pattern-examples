package adapter

import "fmt"

//IPlug 插头
type IPlug interface {
	GetPin() int
}

//TwoPinPlugin 造一个两针的插头
type TwoPinPlugin struct {
}

//GetPin 获取插头针数
func (t *TwoPinPlugin) GetPin() int {
	return 2
}

//IPowerSocket 电源插座
type IPowerSocket interface {
	PlugCharging(p IPlug)
}

//IThreePinPowerSocket 三孔插座
type IThreePinPowerSocket interface {
	ThreePinCharging(p IPlug)
}

//ThreePinPowerSocket 是被适配的目标类
type ThreePinPowerSocket struct{}

//ThreePinCharging 只能为三针的插头通电
func (*ThreePinPowerSocket) ThreePinCharging(p IPlug) {
	if p.GetPin() != 3 {
		fmt.Println("i can not charge for this type")
		return
	}
	fmt.Println("charging for three pin plug")
}

//NewPowerAdapter  生产一个新的电源适配器
func NewPowerAdapter(threePinPowerSocket IThreePinPowerSocket) IPowerSocket {
	return &PowerAdapter{IThreePinPowerSocket(threePinPowerSocket), 0}
}

//PowerAdapter 是能充电的关键电源转换器
type PowerAdapter struct {
	IThreePinPowerSocket
	pin int
}

//GetPin Adapter 的兼容能力
func (p *PowerAdapter) GetPin() int {
	return p.pin
}

//PlugCharging 在PowerAdapter上进行实现
func (p *PowerAdapter) PlugCharging(ip IPlug) {
	if ip.GetPin() == 2 {
		p.pin = 3
		p.ThreePinCharging(p)
	}

}
