/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-10 22:00:58
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 14:22:42
 */

package circuit

import (
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
