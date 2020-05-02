package abstractfactory

import "testing"

func TestRobotBatteryFactory(t *testing.T) {

	factory := &HomeRobotFactory{}
	robot := factory.CreateRobot()
	robot.DoWork()
	battery := factory.CreateBattery()
	battery.Charge(robot)
}

func TestSQLFactory(t *testing.T) {

	factory := &SQLFactory{}
	orderWorker := factory.CreateOrderWorker()
	orderWorker.SaveOrder()
	detailWorker := factory.CreateOrderDetailWorker()
	detailWorker.SaveOrderDetail()
}

func TestNoSqlFactory(t *testing.T) {

	factory := &NoSQLFactory{}
	orderWorker := factory.CreateOrderWorker()
	orderWorker.SaveOrder()
	detailWorker := factory.CreateOrderDetailWorker()
	detailWorker.SaveOrderDetail()
}
