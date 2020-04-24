package factorymethod

import (
	"fmt"
	"strconv"
)

//Assistant 是robot能做的事情
type Assistant interface {
	Clean(int)
	Speak(string)
	Work() string
}

//IRobotFactory must be implemented by Factory
//different Factory create different robot
type IRobotFactory interface {
	Build() Assistant
}

//BasicRobotModel 是基本的机器人模型
type BasicRobotModel struct {
	words    string
	workTime int
}

//Clean 打扫
func (b *BasicRobotModel) Clean(a int) {
	b.workTime = a
	fmt.Printf("i can clean :%d hours\n", a)
}

//Speak 说话
func (b *BasicRobotModel) Speak(w string) {
	b.words = w
	fmt.Printf("my name is: %s\n", w)
}

//Work  main work
func (b *BasicRobotModel) Work() string {
	return fmt.Sprint("my main work is do somthing")
}

//FightingRobotFactory 生产各类军工机器人
type FightingRobotFactory struct{}

//Build a robot from FightingRobotFactory
func (FightingRobotFactory) Build() Assistant {
	return &FightingRobot{
		BasicRobotModel: &BasicRobotModel{},
	}
}

//FightingRobot 实际的战斗机器人
type FightingRobot struct {
	*BasicRobotModel
}

//Work for FightingRobot to do some fighting
func (f FightingRobot) Work() string {
	fmt.Printf("%s\n", "i can fighting")
	return "i can fighting" + strconv.Itoa(f.workTime)
}

//HomeRobotFactory 生产各类家用机器人
type HomeRobotFactory struct{}

//Build a robot from HomeRobotFactory
func (HomeRobotFactory) Build() Assistant {
	return &HomeRobot{
		BasicRobotModel: &BasicRobotModel{},
	}
}

//HomeRobot 实际的家用机器人
type HomeRobot struct {
	*BasicRobotModel
}

//Work robot do some work
func (h HomeRobot) Work() string {
	fmt.Printf("%s\n", "i can do homework")
	return "i can do homework" + strconv.Itoa(h.workTime)
}
