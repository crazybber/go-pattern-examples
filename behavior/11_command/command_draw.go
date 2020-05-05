package command

import (
	"strconv"
)

type command interface {
	Execute() string
}

//PathPainter invoke an draw command
type PathPainter struct {
	commands []command
}

//Execute run cmd
func (p *PathPainter) Execute() string {
	var result string
	for _, command := range p.commands {
		result += command.Execute() + "\n"
	}
	return result
}

//Append new cmd PathPainter
func (p *PathPainter) Append(command command) {
	p.commands = append(p.commands, command)
}

//Undo last step cmd
func (p *PathPainter) Undo() {
	if len(p.commands) != 0 {
		p.commands = p.commands[:len(p.commands)-1]
	}
}

//Clear all
func (p *PathPainter) Clear() {
	p.commands = []command{}
}

//Position pos
type Position struct {
	X, Y int
}

//DrawCommand 命令执行者
//DrawCommand line to
type DrawCommand struct {
	Position *Position
}

//Execute cmd
func (d *DrawCommand) Execute() string {
	return strconv.Itoa(d.Position.X) + "." + strconv.Itoa(d.Position.Y)
}
