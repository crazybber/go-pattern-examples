/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-06-02 23:57:40
 * @Last Modified by: Edward
 * @Last Modified time: 2020-06-03 23:50:17
 */

package circuit

import (
	"context"
	"time"
)

//BreakConditionWatcher check state
type BreakConditionWatcher func(state State, cnter counters) bool

//StateChangedEventHandler set event handle
type StateChangedEventHandler func(name string, from State, to State)

//Option set Options
type Option func(opts *Options)

//Options for breaker
type Options struct {
	Name               string
	Expiry             time.Time
	Interval, Timeout  time.Duration
	MaxRequests        uint32
	CanOpen            BreakConditionWatcher //是否应该断开电路(打开电路开关)
	CanClose           BreakConditionWatcher //if we should close switch
	OnStateChanged     StateChangedEventHandler
	ShoulderHalfToOpen uint32
	Ctx                context.Context
}

//ActionName of breaker
func ActionName(name string) Option {
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

//WithShoulderHalfToOpen of breaker
func WithShoulderHalfToOpen(shoulderHalfToOpen uint32) Option {
	return func(opts *Options) {
		opts.ShoulderHalfToOpen = shoulderHalfToOpen
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

//WithBreakCondition check traffic state ,to see if request can go
func WithBreakCondition(whenCondition BreakConditionWatcher) Option {
	return func(opts *Options) {
		opts.CanOpen = whenCondition
	}
}

//WithCloseCondition check traffic state ,to see if request can go
func WithCloseCondition(whenCondition BreakConditionWatcher) Option {
	return func(opts *Options) {
		opts.CanClose = whenCondition
	}
}
