package chain

import "fmt"

////////////////////////////////
//使用费用申请审批的例子

//FeeRequest 这就要要处理的对象
type FeeRequest struct {
	Mount         int    //申请的金额
	RequiredLevel int    //审批人等级要求
	Name          string //申请人
}

//IApprove 审批动作
type IApprove interface {
	SetNext(next IApprove)    //设置下一个审批流转
	HaveRight(level int) bool //
	HandleApproval(request FeeRequest) bool
}

////////////////////////////////
//实现方式1
////////////////////////////////

//FeeRequestChainFlow 定义一个流程管理器
type FeeRequestChainFlow struct {
	approvers []IApprove //下一个流程节点
}

//AddApprover 添加一个审批对象
func (f *FeeRequestChainFlow) AddApprover(approver IApprove) {
	f.approvers = append(f.approvers, approver)
}

//RunApprovalFlow to deal request by chains
func (f *FeeRequestChainFlow) RunApprovalFlow(request FeeRequest) {

	for i := 0; i < len(f.approvers); i++ {
		result := f.approvers[i].HandleApproval(request)
		if !result {
			//中间有一个环节出问题，流程就终止
			break
		}
	}

}

////////////////////////////////
//实现方式2
////////////////////////////////

//GM 总经理要审批
type GM struct {
	nextHandler IApprove //下一个流程节点
	level       int
}

//NewGM 总经理审批
func NewGM() IApprove {
	return &GM{level: 8}
}

//SetNext 设置下一个审批节点
func (g *GM) SetNext(next IApprove) {
	g.nextHandler = next
}

//HaveRight 处理审批所需要的权限级别
func (g *GM) HaveRight(RequiredLevel int) bool {
	return g.level > RequiredLevel
}

//HandleApproval 进行审批
func (g *GM) HandleApproval(request FeeRequest) bool {
	if g.HaveRight(request.RequiredLevel) {
		fmt.Printf("GM permit %s %d fee request\n", request.Name, request.Mount)
		return true
	}
	fmt.Printf("GM NO right to approve %s %d fee request\n", request.Name, request.Mount)
	//direct forward to Next One
	if g.nextHandler != nil {
		return g.nextHandler.HandleApproval(request)
	}
	return true
}

//CFO 需要审批
type CFO struct {
	nextHandler IApprove //下一个流程节点
	level       int
}

//NewCFO 对象
func NewCFO() IApprove {
	return &CFO{}

}

//HaveRight CFO总是有权限的
func (*CFO) HaveRight(RequiredLevel int) bool {
	return true
}

//SetNext 设置下一个审批节点
func (c *CFO) SetNext(next IApprove) {
	c.nextHandler = next
}

//HandleApproval 进行审批
func (c *CFO) HandleApproval(request FeeRequest) bool {
	if request.Mount < 1e+10 {
		fmt.Printf("CFO permit %s %d fee request\n", request.Name, request.Mount)
		return true
	}
	fmt.Printf("CFO No right to approve %s %d fee request \n", request.Name, request.Mount)
	if c.nextHandler != nil {
		return c.nextHandler.HandleApproval(request)
	}
	return true
}

//CEO 需要审批
type CEO struct {
}

//NewCEO 对象
func NewCEO() IApprove {
	return &CEO{}
}

//HaveRight CEO总是有权限的
func (*CEO) HaveRight(RequiredLevel int) bool {
	return true
}

//SetNext 设置下一个审批节点
func (c *CEO) SetNext(next IApprove) {
	//no thing to do
}

//HandleApproval 进行审批
func (*CEO) HandleApproval(request FeeRequest) bool {
	if request.Mount < 1e+15 {
		fmt.Printf("CEO permit %s %d fee request\n", request.Name, request.Mount)
		return true
	}
	fmt.Printf("CEO deny %s %d fee request \n", request.Name, request.Mount)
	return false
}
