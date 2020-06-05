/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-10 22:00:58
 * @Last Modified by: Edward
 * @Last Modified time: 2020-06-03 23:54:29
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
const (
	defaultTimeout          = time.Second * 30
	defaultSuccessThreshold = 2
)

////////////////////////////////
//way 2 对象式断路器
// 高级模式
// 支持多工作者模式
////////////////////////////////

//RequestBreaker for protection
type RequestBreaker struct {
	options  Options
	mutex    sync.Mutex
	state    State
	cnter    counters
	preState State
}

//NewRequestBreaker return a breaker
func NewRequestBreaker(opts ...Option) *RequestBreaker {

	defaultOptions := Options{
		Name:           "defaultBreakerName",
		Expiry:         time.Now().Add(time.Second * 20),
		Interval:       time.Second * 10, // interval to check  closed status,default 10 seconds
		Timeout:        time.Second * 60, //timeout to check open, default 60 seconds
		MaxRequests:    5,
		CanOpen:        func(current State, cnter counters) bool { return cnter.ConsecutiveFailures > 2 },
		CanClose:       func(current State, cnter counters) bool { return cnter.ConsecutiveSuccesses > 2 },
		OnStateChanged: func(name string, fromPre State, toCurrent State) {},
	}

	for _, setOption := range opts {
		setOption(&defaultOptions)

	}

	return &RequestBreaker{
		options:  defaultOptions,
		cnter:    counters{},
		state:    StateUnknown,
		preState: StateUnknown,
	}
}

func (rb *RequestBreaker) changeStateTo(state State) {
	rb.preState = rb.state
	rb.state = state
	rb.cnter.Reset()
}

func (rb *RequestBreaker) beforeRequest() error {

	rb.mutex.Lock()
	defer rb.mutex.Unlock()
	fmt.Println("before do request:", rb.cnter.Total())

	switch rb.state {
	case StateOpen:
		//如果是断开状态，并且超时了，转到半开状态，记录
		if rb.options.Expiry.Before(time.Now()) {
			rb.changeStateTo(StateHalfOpen)
			rb.options.Expiry = time.Now().Add(rb.options.Timeout)
			return nil
		}
		return ErrTooManyRequests
	case StateClosed:
		if rb.options.Expiry.Before(time.Now()) {
			rb.cnter.Reset()
			rb.options.Expiry = time.Now().Add(rb.options.Interval)
		}
	}

	return nil

}

// Do the given requested work if the RequestBreaker accepts it.
// Do returns an error instantly if the RequestBreaker rejects the request.
// Otherwise, Execute returns the result of the request.
// If a panic occurs in the request, the RequestBreaker handles it as an error and causes the same panic again.
func (rb *RequestBreaker) Do(work func(ctx context.Context) (interface{}, error)) (interface{}, error) {

	//before

	if err := rb.beforeRequest(); err != nil {
		return nil, err
	}

	//do work
	//do work from requested user
	result, err := work(rb.options.Ctx)

	//after work
	rb.afterRequest(err)

	return result, err
}

func (rb *RequestBreaker) afterRequest(resultErr error) {

	rb.mutex.Lock()
	defer rb.mutex.Unlock()
	//after
	fmt.Println("after do request:", rb.cnter.Total())

	if resultErr != nil {
		//失败了,handle 失败
		rb.cnter.Count(FailureState, rb.preState == rb.state)
		switch rb.state {
		case StateHalfOpen, StateClosed:
			if rb.options.CanOpen(rb.state, rb.cnter) {
				rb.changeStateTo(StateOpen)                                     //打开开关
				rb.options.OnStateChanged(rb.options.Name, rb.state, StateOpen) //关闭到打开
			}
		}
	} else {
		//success !
		rb.cnter.Count(SuccessState, rb.preState == rb.state)

		switch rb.state {
		case StateHalfOpen:
			if rb.preState == StateOpen {
				//	rb.changeStateTo(StateHalfOpen) //already handled in beforeRequest,Only fire StateChange Event
				rb.options.OnStateChanged(rb.options.Name, StateOpen, StateHalfOpen) //打开到半开
			}
			if rb.cnter.ConsecutiveSuccesses >= rb.options.ShoulderHalfToOpen {
				rb.changeStateTo(StateClosed)
				rb.options.Expiry = time.Now().Add(rb.options.Interval)
				rb.options.OnStateChanged(rb.options.Name, StateHalfOpen, StateClosed) //半开到关闭
			}

		}

	}

}
