/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-06-02 23:57:40
 * @Last Modified by: Edward
 * @Last Modified time: 2020-06-02 23:57:40
 */

package circuit

import (
	"context"
	"time"
)

//BreakConditionWatcher check state
type BreakConditionWatcher func(cnter counters) bool

//StateChangedEventHandler set event handle
type StateChangedEventHandler func(name string, from State, to State)

//Option set Options
type Option func(opts *Options)

//Options for breaker
type Options struct {
	Name              string
	Expiry            time.Time
	Interval, Timeout time.Duration
	MaxRequests       uint32
	WhenToBreak       BreakConditionWatcher //是否应该断开电路(打开电路开关)
	OnStateChanged    StateChangedEventHandler
	Ctx               context.Context
}

//Name of breaker
func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}

//Interval of breaker
func Interval(interval time.Duration) Option {
	return func(opts *Options) {
		opts.Interval = interval
	}
}

//Timeout of breaker
func Timeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

// MaxRequests is the maximum number of requests allowed to pass through
// when the CircuitBreaker is half-open.
// If MaxRequests is 0, the CircuitBreaker allows only 1 request.

//MaxRequests of breaker
func MaxRequests(maxRequests uint32) Option {
	return func(opts *Options) {
		opts.MaxRequests = maxRequests
	}
}

//Expiry of breaker
func Expiry(expiry time.Time) Option {
	return func(opts *Options) {
		opts.Expiry = expiry
	}
}

//WithStateChanged set handle of ChangedHandle
func WithStateChanged(handler StateChangedEventHandler) Option {
	return func(opts *Options) {
		opts.OnStateChanged = handler
	}
}

//BreakIf check traffic state ,to see if request can go
func BreakIf(whenCondition BreakConditionWatcher) Option {
	return func(opts *Options) {
		opts.WhenToBreak = whenCondition
	}
}
