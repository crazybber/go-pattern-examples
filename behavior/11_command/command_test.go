package command

import "testing"

func TestTroopCommand(t *testing.T) {

	//教官
	in := Instructor{}

	//受训队伍1
	tr1 := &Troop{name: "girl team"}
	turnLeft := &TurnLeftCommand{receiver: tr1}

	//受训队伍2
	tr2 := &Troop{name: "boy team"}

	turnRight := &TurnLeftCommand{receiver: tr2}

	//准备命令发给不同的队伍
	in.AddCommand(turnRight)
	in.AddCommand(turnLeft)

	turnback := &TurnBackCommand{receiver: tr2}

	in.AddCommand(turnback)

	turnback = &TurnBackCommand{receiver: tr1}

	in.AddCommand(turnback)

	//开始发布指令
	in.Execute()

}
