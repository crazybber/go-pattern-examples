package factorymethod

import "fmt"

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
	words string
	a, b  int
}

//Clean 打扫
func (o *BasicRobotModel) Clean(a int) {
	fmt.Printf("%d", a)
}

//Speak 说话
func (o *BasicRobotModel) Speak(b int) {
	o.b = b
}

//Work  main work
func (o *BasicRobotModel) Work() string {
	fmt.Sprint("my main work is do somthing")
}

//FightingRobotFactory 生产各类军工机器人
type FightingRobotFactory struct{}

//Build a robot from FightingRobotFactory
func (FightingRobotFactory) Build() Assistant {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//FightingRobot 实际的战斗机器人
type FightingRobot struct {
	*BasicRobotModel
}

//Result 获取结果
func (o FightingRobot) Result() int {
	return o.a + o.b
}

//HomeRobotFactory 生产各类家用机器人
type HomeRobotFactory struct{}

//Build a robot from HomeRobotFactory
func (HomeRobotFactory) Build() Assistant {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//HomeRobot 实际的家用机器人
type HomeRobot struct {
	*OperatorBase
}

//Result 获取结果
func (o HomeRobot) Result() int {
	return o.a - o.b
}
