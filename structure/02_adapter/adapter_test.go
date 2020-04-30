package adapter

import (
	"testing"
)

func TestPowerSocketAdapter(t *testing.T) {

	plug := &TwoPinPlugin{}

	threePinSocket := ThreePinPowerSocket{}

	//三孔插头是不能为两针插头充电的,可以试试看
	threePinSocket.ThreePinCharging(plug)

	//只好加一个电源适配器
	powersocket := NewPowerAdapter(&threePinSocket)

	//再试试能不能充电
	powersocket.PlugCharging(plug)

}
