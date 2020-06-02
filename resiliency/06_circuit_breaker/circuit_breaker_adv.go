/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-10 22:00:58
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 17:25:56
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
	options    Options
	mutex      sync.Mutex
	state      State
	generation uint64
	cnter      counters
}

//NewRequestBreaker return a breaker
func NewRequestBreaker(opts ...Option) *RequestBreaker {

	defaultOptions := Options{
		Name:           "defaultBreakerName",
		Expiry:         time.Now().Add(time.Second * 20),
		Interval:       time.Second * 2,  // interval to check status
		Timeout:        time.Second * 60, //default to 60 seconds
		MaxRequests:    5,
		WhenToBreak:    func(cnter counters) bool { return cnter.ConsecutiveFailures > 2 },
		OnStateChanged: func(name string, fromPre State, toCurrent State) {},
	}

	for _, setOption := range opts {
		setOption(&defaultOptions)

	}

	return &RequestBreaker{
		options:    defaultOptions,
		cnter:      counters{},
		generation: 0,
	}
}

// Do the given requested work if the RequestBreaker accepts it.
// Do returns an error instantly if the RequestBreaker rejects the request.
// Otherwise, Execute returns the result of the request.
// If a panic occurs in the request, the RequestBreaker handles it as an error and causes the same panic again.
func (rb *RequestBreaker) Do(work func(ctx context.Context) (interface{}, error)) (interface{}, error) {

	preState := StateUnknown

	//before
	fmt.Println("before do : request:", rb.cnter.Total())
	rb.mutex.Lock()
	switch rb.state {
	case StateOpen:
		return nil, ErrTooManyRequests
	case StateHalfOpen:
		//如果是断开状态，并且超时了，转到半开状态，记录
		if rb.options.Expiry.Before(time.Now()) {
			rb.state = StateHalfOpen
			preState = rb.state
			rb.cnter.Reset()

		}

	}

	rb.mutex.Unlock()

	//do work
	//do work from requested user
	result, err := work(rb.options.Ctx)

	rb.mutex.Lock()
	//失败了
	if err != nil {
		rb.cnter.Count(FailureState)
		//如果是在半开状态下的失败，立即打开开关
		if rb.state == StateHalfOpen {
			rb.state = StateOpen //转为打开
		} else if rb.state == StateClosed {
			if rb.options.WhenToBreak(rb.cnter) {
				rb.state = StateOpen //打开开关
				rb.cnter.Reset()
			}
		}

	} else {
		rb.cnter.Count(SuccessState)
		if preState == StateOpen && rb.state == StateHalfOpen {
			rb.options.OnStateChanged(rb.options.Name, StateOpen, StateHalfOpen)
		}

	}
	rb.mutex.Unlock()
	//after
	fmt.Println("after do : request:", rb.cnter.Total())

	return result, err
}
