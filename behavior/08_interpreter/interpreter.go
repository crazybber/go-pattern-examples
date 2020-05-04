package interpreter

import (
	"strconv"
	"strings"
)

//用会议交流的例子
// A表示左边发言者，B表示右边发言者
// A "->"  B 表示  A说，B听,此时B不能发言。
// A "<-"  B 表示  B说，A听,此时B不能发言。
// A "<->" B 表示  A 和 B 可以自由发言。
// A  "-"  B 表示  A 和 B 都不能发言，只能倾听。



type IActionInterpret interface {
	Interpret()
}

//Speaker 发言者
type Speaker struct {
	Name string
	Age int

}

func (s *Speaker) Interpret()  {

}

//LeftIntroduce A向B介绍自己
type LeftIntroduce struct {
	leftSpeaker, rightSpeaker Speaker
}


func (l *LeftIntroduce) Interpret()  {

}

//RightIntroduce B向A介绍自己
type RightIntroduce struct {
	leftSpeaker, rightSpeaker Speaker
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

//标识解析器
type SignParser struct {
	actionsMark   []string
}

//Parse 标识解析器进行解析
func (p *SignParser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	s := p.actionsMark[0]
	for len(s) >0 {
		switch s.actionsMark[0] {

		}
	}
}



