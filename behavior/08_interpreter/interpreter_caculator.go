package interpreter

import (
	"strconv"
	"strings"
)

//解释自定义的加减法运算
//输入：字符串
//输出：整数值
//将一个包含加减运算的字符串，正常解析出结果

//Element 每个元素的解释接口
type Element interface {
	Interpret() int
}

//ValElement 值节点
type ValElement struct {
	val int
}

//Interpret 值解析单元的返回值
func (n *ValElement) Interpret() int {
	return n.val
}

//AddOperate Operation(+)
type AddOperate struct {
	left, right Element
}

//Interpret AddOperate
func (n *AddOperate) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

//MinOperate Operation(-)
type MinOperate struct {
	left, right Element
}

//Interpret MinOperate
func (n *MinOperate) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

//Parser machine
type Parser struct {
	exp   []string
	index int
	prev  Element
}

//Parse content
func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddOperte()
		case "-":
			p.prev = p.newMinOperte()
		default:
			p.prev = p.newValElement()
		}
	}
}

func (p *Parser) newAddOperte() Element {
	p.index++
	return &AddOperate{
		left:  p.prev,
		right: p.newValElement(),
	}
}

func (p *Parser) newMinOperte() Element {
	p.index++
	return &MinOperate{
		left:  p.prev,
		right: p.newValElement(),
	}
}

func (p *Parser) newValElement() Element {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValElement{
		val: v,
	}
}

//Result of parsing result
func (p *Parser) Result() Element {
	return p.prev
}
