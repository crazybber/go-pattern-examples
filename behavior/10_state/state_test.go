package state

import (
	"testing"
	"time"
)

func TestRobotState(t *testing.T) {

	robot := NewSportRobot("home keeper")

	//主动驱动行为
	robot.Move()

	//跑
	robot.UpdateState(&RuningState{})

	//跑一会
	time.Sleep(time.Millisecond * 234)

	//走
	robot.UpdateState(&WalkState{})

	//走一会儿
	time.Sleep(time.Millisecond * 544)
	//继续跑
	robot.UpdateState(&RuningState{})

}
func TestAlarmState(t *testing.T) {

	expect := "vibrating humming ... vibrating humming...vibrating humming..." +
		"vibrating humming ... vibrating humming...vibrating humming..." +
		"sun rise ,get up ,get up get up..."

	//创建一个手机闹铃
	mobile := NewAlert()

	ringsSounds := mobile.Alert()

	//叠加振铃声音，振铃响两遍
	ringsSounds += mobile.Alert()

	//设置振铃的铃声
	mobile.SetState(&AlertSong{})

	ringsSounds += mobile.Alert()

	if ringsSounds != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, ringsSounds)
	}
}
