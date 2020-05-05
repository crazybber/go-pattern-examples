package interpreter

import (
	"fmt"
	"strings"
)

//用会议交流的例子，规则如下:
// A表示左边发言者，B表示右边发言者
// A "->"  B 表示  A说，B听,此时B不能发言。
// A "<-"  B 表示  B说，A听,此时B不能发言。
// A "<->" B 表示  A 和 B 可以自由发言。

//IActionInterpret 解释器
type IActionInterpret interface {
	Interpret()
}

//SpeakerUnit 每个参与动作的单元都是，会议发言者
type SpeakerUnit struct {
	Name     string
	CanSpeak bool //每个参会者有两种基本的行为，发言或者安静
}

//Interpret 解释自己的基本行为
func (s *SpeakerUnit) Interpret() {
	if s.CanSpeak {
		fmt.Println("i'm", s.Name, "i’m speaking")
		return
	}
	fmt.Println("i'm", s.Name, "already silent")
}

//LeftSpeakAction ：A "->"  B 表示  A说，B听,此时B不能发言。
type LeftSpeakAction struct {
	leftSpeaker, rightSpeaker IActionInterpret
}

//Interpret 解释执行
func (l *LeftSpeakAction) Interpret() {

	l.leftSpeaker.Interpret()
	l.rightSpeaker.Interpret()

}

//RightSpeakAction ：A "<-"  B 表示  B说，A听,此时B不能发言。
type RightSpeakAction struct {
	leftSpeaker, rightSpeaker IActionInterpret
}

//Interpret 解释执行
func (r *RightSpeakAction) Interpret() {
	r.leftSpeaker.Interpret()
	r.rightSpeaker.Interpret()
}

//BothSpeakAction : A "<->" B 表示  A 和 B 可以自由发言。
type BothSpeakAction struct {
	leftSpeaker, rightSpeaker IActionInterpret
}

//Interpret 解释执行
func (b *BothSpeakAction) Interpret() {
	b.leftSpeaker.Interpret()
	b.rightSpeaker.Interpret()
}

//SignParser 我们自己的DSL解析器
type SignParser struct {
	actionUnits []string         //要解析的内容
	result      IActionInterpret //上一个也是解释器单元
}

//解析 ->
func (s *SignParser) newLeftSpeakAction() IActionInterpret {

	left := &SpeakerUnit{
		Name:     s.actionUnits[0],
		CanSpeak: true,
	}
	right := &SpeakerUnit{
		Name: s.actionUnits[2],
	}
	return &LeftSpeakAction{
		leftSpeaker:  left,
		rightSpeaker: right,
	}
}

//解析 <-
func (s *SignParser) newRightSpeakAction() IActionInterpret {

	left := &SpeakerUnit{
		Name: s.actionUnits[0],
	}
	right := &SpeakerUnit{
		Name:     s.actionUnits[2],
		CanSpeak: true,
	}
	return &LeftSpeakAction{
		leftSpeaker:  left,
		rightSpeaker: right,
	}
}

//解析 <->
func (s *SignParser) newBothSpeakAction() IActionInterpret {

	left := &SpeakerUnit{
		Name:     s.actionUnits[0],
		CanSpeak: true,
	}
	right := &SpeakerUnit{
		Name:     s.actionUnits[2],
		CanSpeak: true,
	}
	return &LeftSpeakAction{
		leftSpeaker:  left,
		rightSpeaker: right,
	}
}

//Parse 标识解析器进行解析，exp就是要解释的内容
func (s *SignParser) Parse(exp string) {

	s.actionUnits = strings.Split(exp, " ") //单元分割符

	switch s.actionUnits[1] {
	case "->":
		s.result = s.newLeftSpeakAction()
	case "<-":
		s.result = s.newRightSpeakAction()
	case "<->":
		s.result = s.newBothSpeakAction()
	default:
		fmt.Println("some error raised")
	}

}

//Result 就是两边正确执行了动作
func (s *SignParser) Result() {

	s.result.Interpret()

}
