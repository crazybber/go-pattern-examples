package state

import (
	"fmt"
	"time"
)

////////////////////////////////
//使用机器人状态转换的例子
//简化示例，假定机器人有：走，跑，静止三种活动状态

//IRobotState 代表机器人此时的活动状态
type IRobotState interface {
	//上下文可以用于传递参数
	Move(*SportRobot)
}

//SportRobot 定义一个运动机器人
type SportRobot struct {
	Name         string
	stateMachine IRobotState //状态机
}

//Move action，主动行为
func (s *SportRobot) Move() {
	s.stateMachine.Move(s)
}

//UpdateState 更新状态
func (s *SportRobot) UpdateState(newState IRobotState) {
	s.stateMachine = newState
	//触发行为
	s.stateMachine.Move(s)
}

//NewSportRobot 生产一个机器人
func NewSportRobot(name string) *SportRobot {
	//默认一个状态
	return &SportRobot{
		Name:         name,
		stateMachine: &StillState{},
	}

}

//RuningState 奔跑状态
type RuningState struct {
}

//Move 实现跑的状态
func (r *RuningState) Move(robot *SportRobot) {

	fmt.Println("i'm robot:", robot.Name)
	fmt.Println("i am running", time.Now())

}

//WalkState 行走状态
type WalkState struct {
}

//Move 实现行走状态
func (w *WalkState) Move(robot *SportRobot) {
	fmt.Println("i'm robot:", robot.Name)
	fmt.Println("i am waling", time.Now())

}

//StillState 静止状态
type StillState struct {
}

//Move 实现静止状态
func (s *StillState) Move(robot *SportRobot) {
	fmt.Println("i'm robot:", robot.Name)
	fmt.Println("i am sitting", time.Now())

}
