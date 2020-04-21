package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	buttion1 Command
	buttion2 Command
}

func NewBox(buttion1, buttion2 Command) *Box {
	return &Box{
		buttion1: buttion1,
		buttion2: buttion2,
	}
}

func (b *Box) PressButtion1() {
	b.buttion1.Execute()
}

func (b *Box) PressButtion2() {
	b.buttion2.Execute()
}
