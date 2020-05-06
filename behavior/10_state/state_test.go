package state

import "testing"

func TestState(t *testing.T) {

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
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
