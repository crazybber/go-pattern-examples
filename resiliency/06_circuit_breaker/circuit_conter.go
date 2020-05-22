/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-22 12:41:54
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 12:41:54
 */

package circuit

import "time"

//OperationState of current 某一次操作的结果状态
type OperationState int

//states of CircuitBreaker
//states: closed --->open ---> half open --> closed
const (
	UnknownState OperationState = iota
	FailureState
	SuccessState
)

//ICounter interface
type ICounter interface {
	Count(OperationState)
	LastActivity() time.Time
	Reset()
	Total() uint32
}

type counters struct {
	Requests             uint32 //连续的请求次数
	lastState            OperationState
	lastActivity         time.Time
	counts               uint32 //counts of failures
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
	ct.lastState = c.lastState
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
	c.lastState = statue
	//handle status change

}
