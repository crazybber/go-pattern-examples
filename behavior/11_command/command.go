package command

import "fmt"

////////////////////////////////
//这里使用军训的例子，使用队列向左转，向右转，向后转的口令
//命令发出者是教官,命令执行者是队列

//ICommand 命令接口，队列要进行响应
type ICommand interface {
	Execute()
}

//Troop 队列
type Troop struct{ name string }

//Execute cmd
func (t *Troop) Execute() {
	fmt.Println("cmd had been executed by", t.name)
}

//TurnLeftCommand 向左转
type TurnLeftCommand struct {
	//可以携带参数,每个命令有一个接收者，去执行
	receiver ICommand
}

//Execute 执行向左转
func (t *TurnLeftCommand) Execute() {
	fmt.Print("Troop Turn Left\n")
	t.receiver.Execute()
}

//TurnRightCommand 向右转
type TurnRightCommand struct {
	//可以携带参数
	receiver ICommand
}

//Execute 执行向右转
func (t *TurnRightCommand) Execute() {
	fmt.Print("Troop Turn Right\n")
	t.receiver.Execute()
}

//TurnBackCommand 向后转
type TurnBackCommand struct {
	//可以携带参数
	holdTime int //保持时间
	receiver ICommand
}

//Execute 执行向后转
func (t *TurnBackCommand) Execute() {
	fmt.Print("Troop Turn Back\n")
	t.receiver.Execute()
}

// Instructor 教官
type Instructor struct {
	commands []ICommand
}

//AddCommand 教官喊口令一般都是一连串
func (i *Instructor) AddCommand(c ICommand) {
	i.commands = append(i.commands, c)
}

//Execute  教官的Execute是发出命令
func (i *Instructor) Execute() {

	for _, cmd := range i.commands {
		cmd.Execute()
	}

}
