package abstractfactory

import "fmt"

//IRobot : 机器人能做的事情
type IRobot interface {
	Name() string
	DoWork()
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法做一件事
}

//IBattery 电池能做的事情
type IBattery interface {
	Charge(robot IRobot)
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法做一件事
}

//IProduce 是当抽象工厂模式例子中的关键接口
//IProduce 返回一组产品对象
//IProduce 本质是创建工作对象，但必须以接口方式返回
type IProduce interface {
	CreateRobot() IRobot
	CreateBattery() IBattery
	//.....BALABALA..可以继续写很多接口方法
	//每个接口方法都要返回一个接口
}

////////////////////////////////
//接口定义好了，开始进行实现和应用
////////////////////////////////

//HomeRobot 家用机器人
type HomeRobot struct{}

//DoWork 机器人可以做工作
func (*HomeRobot) DoWork() {
	fmt.Print("robot is cleaning home\n")
}

//Name 机器人的名字
func (*HomeRobot) Name() string {
	return fmt.Sprint("home robot")
}

//HomeBattery 家用电池
type HomeBattery struct{}

// Charge SaveOrderDetail接口,保存订单细节
func (*HomeBattery) Charge(robot IRobot) {

	rn := robot.Name()
	fmt.Print("HomeBattery is charging for:", rn)
	fmt.Println()
}

//HomeRobotFactory 家用机器人工厂
type HomeRobotFactory struct{}

//CreateRobot 创建机器人
func (*HomeRobotFactory) CreateRobot() IRobot {
	return &HomeRobot{}
}

//CreateBattery  创建电池
func (*HomeRobotFactory) CreateBattery() IBattery {
	return &HomeBattery{}
}
