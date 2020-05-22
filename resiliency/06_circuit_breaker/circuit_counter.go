/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-22 12:41:54
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 17:22:35
 */

package circuit

import "time"

////////////////////////////////
/// 计数器 用以维护断路器内部的状态
/// 无论是对象式断路器还是函数式断路器
/// 都要用到计数器
////////////////////////////////

//State  断路器本身的状态的
//State  of switch int
type State int

// state for breaker
const (
	StateClosed State = iota //默认的闭合状态，可以正常执行业务
	StateHalfOpen
	StateOpen
)

//OperationState of current 某一次操作的结果状态
type OperationState int

//ICounter interface
type ICounter interface {
	Count(OperationState)
	LastActivity() time.Time
	Reset()
	Total() uint32
}

type counters struct {
	Requests             uint32 //连续的请求次数
	lastActivity         time.Time
	TotalFailures        uint32
	TotalSuccesses       uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

func (c *counters) Total() uint32 {
	return c.Requests
}

func (c *counters) LastActivity() time.Time {
	return c.lastActivity
}

func (c *counters) Reset() {
	ct := &counters{}
	ct.lastActivity = c.lastActivity
	c = ct
}

//Count the failure and success
func (c *counters) Count(statue OperationState) {

	switch statue {
	case FailureState:
		c.ConsecutiveFailures++
	case SuccessState:
		c.ConsecutiveSuccesses++
	}
	c.Requests++
	c.lastActivity = time.Now() //更新活动时间
	//	c.lastOpResult = statue
	//handle status change

}
