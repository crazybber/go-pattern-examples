/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-22 12:42:34
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 14:35:00
 */

package circuit

import (
	"context"
	"time"
)

////////////////////////////////
//way 2 简单的函数式断路器
// 一个func实例作为一个Breaker 允许多个worker(即goroutine)同时访问
// 理论上讲也需要考虑同步问题
// 当前的设计思路，非常简单：
//     1、不考虑三种状态之间的转换，只靠两种状态，电路打开与电路关闭
//     2、当累计失败到达一定失败次数就端开请求(打开电路)，并延迟一定的时间再允许请求
////////////////////////////////

//states of CircuitBreaker
//states: closed --->open ---> half open --> closed
const (
	UnknownState OperationState = iota
	FailureState
	SuccessState
)

//Circuit of action stream,this is actually to do something.
//Circuit hold the really action
type Circuit func(context.Context) error

var canRetry = func(cnt counters, failureThreshold uint32) bool {
	backoffLevel := cnt.ConsecutiveFailures - failureThreshold
	// Calculates when should the circuit breaker resume propagating requests
	// to the service
	shouldRetryAt := cnt.LastActivity().Add(time.Second << backoffLevel)
	return time.Now().After(shouldRetryAt)
}

//Breaker return a closure wrapper to hold request
func Breaker(c Circuit, failureThreshold uint32) Circuit {

	//闭包内部的全局计数器 和状态标志
	cnt := counters{}

	//ctx can be used hold parameters
	return func(ctx context.Context) error {

		if cnt.ConsecutiveFailures >= failureThreshold {
			if !canRetry(cnt, failureThreshold) {
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
