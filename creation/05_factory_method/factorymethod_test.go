package factorymethod

import "testing"

func doWork(factory IRobotFactory, cleanhour int) string {
	robot := factory.Build()
	robot.Clean(cleanhour)

	robot.Speak("robot name")

	return robot.Work()

}

func TestRobotFactory(t *testing.T) {
	var factory IRobotFactory

	factory = FightingRobotFactory{}
	if doWork(factory, 2) != "i can fighting2" {
		t.Fatal("error with factory method pattern")
	}

	factory = HomeRobotFactory{}
	if doWork(factory, 1) != "i can do homework1" {
		t.Fatal("error with factory method pattern")
	}
}
