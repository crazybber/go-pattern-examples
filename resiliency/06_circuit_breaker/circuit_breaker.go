/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-10 22:00:58
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-21 17:41:46
 */

package circuit

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

////////////////////////////////
///使用HTTP请求的例子
//每个搜索引擎时时刻刻都会遇到超大规模的请求的流量.
//这里演示一个复杂一点的例子，同时使用Option 模式

//ErrServiceUnavailable for error
var (
	ErrTooManyRequests    = errors.New("too many requests")
	ErrServiceUnavailable = errors.New("service unavailable")
	FailureThreshold      = 10 //最大失败次数--->失败阈值
)

// 默认的超时时间
const defaultTimeout = time.Second * 30

////////////////////////////////
/// 状态计数器 用以维护断路器内部的状态
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

////////////////////////////////
//way 1 对象式断路器
////////////////////////////////

//RequestBreaker for protection
type RequestBreaker struct {
	options    Options
	mutex      sync.Mutex
	state      OperationState //断路器的当前状态
	generation uint64
	counts     ICounter
}

//NewRequestBreaker return a breaker
func NewRequestBreaker(opts ...Option) *RequestBreaker {

	defaultOptions := Options{
		Name:           "defaultBreakerName",
		Expiry:         time.Now().Add(time.Second * 20),
		Interval:       time.Second * 2,  // interval to check status
		Timeout:        time.Second * 60, //default to 60 seconds
		MaxRequests:    5,
		WhenToBreak:    func(counts counters) bool { return counts.ConsecutiveFailures > 2 },
		OnStateChanged: func(name string, fromPre State, toCurrent State) {},
	}

	for _, setOption := range opts {
		setOption(&defaultOptions)

	}

	return &RequestBreaker{
		options:    defaultOptions,
		counts:     &counters{},
		generation: 0,
	}
}

// Do the given requested work if the RequestBreaker accepts it.
// Do returns an error instantly if the RequestBreaker rejects the request.
// Otherwise, Execute returns the result of the request.
// If a panic occurs in the request, the RequestBreaker handles it as an error and causes the same panic again.
func (rb *RequestBreaker) Do(work func() (interface{}, error)) (interface{}, error) {

	//before
	fmt.Println("before do : request:", rb.counts.Total())

	//do work from requested user
	result, err := work()

	fmt.Println("after do : request:", rb.counts.Total())

	return result, err
}

////////////////////////////////
//way 2 简单的函数式断路器
////////////////////////////////

//Circuit of action stream,this is actually to do something.
//Circuit hold the really action
type Circuit func(context.Context) error

//Breaker return a closure wrapper to hold request,达到指定的失败次数后电路断开
func Breaker(c Circuit, failureThreshold uint32) Circuit {

	//内部计数器
	cnt := counters{}
	expired := time.Now()
	currentState := StateClosed //默认是闭合状态

	//ctx can be used hold parameters
	return func(ctx context.Context) error {

		if cnt.ConsecutiveFailures >= failureThreshold {

			//断路器在half open状态下的控制逻辑
			canRetry := func(cnt counters) bool {
				//间歇时间，多个线程时候会存在同步文件需要lock操作
				backoffLevel := cnt.ConsecutiveFailures - failureThreshold
				// Calculates when should the circuit breaker resume propagating requests
				// to the service
				shouldRetryAt := cnt.LastActivity().Add(time.Second << backoffLevel)
				return time.Now().After(shouldRetryAt)
			}

			//如果仍然不能执行，直接返回失败
			if !canRetry(cnt) {
				// Fails fast instead of propagating requests to the circuit since
				// not enough time has passed since the last failure to retry
				return ErrServiceUnavailable
			}
		}

		// 可以执行，则执行，并累计成功和失败次数
		// Unless the failure threshold is exceeded the wrapped service mimics the
		// old behavior and the difference in behavior is seen after consecutive failures
		// do the job

		switch currentState {
		case StateOpen:
			if time.Now().Before(expired) {
				currentState = StateHalfOpen //转为半开
			}
			return ErrServiceUnavailable
		case StateClosed:
		case StateHalfOpen:

		}

		if err := c(ctx); err != nil {
			//统计状态
			cnt.Count(FailureState)

			return err
		}

		//统计成功状态
		cnt.Count(SuccessState)
		return nil
	}
}
