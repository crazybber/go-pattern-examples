package circuit

/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-10 22:00:58
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-11 17:46:20
 */

import (
	"context"
	"errors"
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
	FailureThreshold      = 10
)

//StateCheckerHandler check state
type StateCheckerHandler func(counts counters) bool

//StateChangedEventHandler set event handle
type StateChangedEventHandler func(name string, from State, to State)

//Option set Options
type Option func(opts *Options)

//RequestBreaker for protection
type RequestBreaker struct {
	options    Options
	mutex      sync.Mutex
	state      State
	generation uint64
	counts     ICounter
}

//NewRequestBreaker return a breaker
func NewRequestBreaker(opts ...Option) *RequestBreaker {

	defaultOptions := Options{
		Name:           "defaultBreakerName",
		Expiry:         time.Now().Add(time.Second * 20),
		Interval:       time.Second * 2,
		Timeout:        time.Second * 60, //default to 60 seconds
		MaxRequests:    5,
		ReadyToTrip:    func(counts counters) bool { return true },
		OnStateChanged: func(name string, from State, to State) {},
	}

	for _, setOption := range opts {
		setOption(&defaultOptions)

	}

	return &RequestBreaker{
		options:    defaultOptions,
		counts:     nil,
		generation: 0,
	}
}

//State of current switch
type State int

//states of CircuitBreaker
const (
	UnknownState State = iota
	FailureState
	SuccessState
)

//Circuit of action stream
type Circuit func(context.Context) error

//ICounter interface
type ICounter interface {
	Count(State)
	ConsecutiveFailures() uint32
	LastActivity() time.Time
	Reset()
}

type counters struct {
	Requests     uint32
	state        State
	lastActivity time.Time
	counts       uint32 //counts of failures
}

func (c *counters) Count(state State) {

}

func (c *counters) ConsecutiveFailures() uint32 {

	return 0
}

func (c *counters) LastActivity() time.Time {
	return c.lastActivity
}

func (c *counters) Reset() {

}

//NewCounter New Counter for Circuit Breaker
func NewCounter() ICounter {
	return &counters{}
}

//Breaker of circuit
func Breaker(c Circuit, failureThreshold uint32) Circuit {

	cnt := NewCounter()

	return func(ctx context.Context) error {
		if cnt.ConsecutiveFailures() >= failureThreshold {

			canRetry := func(cnt ICounter) bool {
				backoffLevel := cnt.ConsecutiveFailures() - failureThreshold

				// Calculates when should the circuit breaker resume propagating requests
				// to the service
				shouldRetryAt := cnt.LastActivity().Add(time.Second * 2 << backoffLevel)

				return time.Now().After(shouldRetryAt)
			}

			if !canRetry(cnt) {
				// Fails fast instead of propagating requests to the circuit since
				// not enough time has passed since the last failure to retry
				return ErrServiceUnavailable
			}
		}

		// Unless the failure threshold is exceeded the wrapped service mimics the
		// old behavior and the difference in behavior is seen after consecutive failures
		if err := c(ctx); err != nil {
			cnt.Count(FailureState)
			return err
		}

		cnt.Count(SuccessState)
		return nil
	}
}
