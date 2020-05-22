/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-22 12:42:34
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-22 12:42:34
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

		//handle statue transformation for timeout
		if currentState == StateOpen {
			nowt := time.Now()
			if expired.Before(nowt) || expired.Equal(nowt) {
				currentState = StateHalfOpen //端开状态的计时器过期了，转为半开
				cnt.ConsecutiveSuccesses = 0 //重置用于累计成功调用的计数器
			}
		}

		switch currentState {
		case StateOpen:
			return ErrServiceUnavailable //直接失败
		case StateHalfOpen:
			if err := c(ctx); err != nil {
				//统计状态
				cnt.Count(FailureState)
				currentState = StateOpen
				expired = time.Now().Add(defaultTimeout) //Reset默认的超时时间
				return err
			}
			//统计成功状态
			cnt.Count(SuccessState)
			//处理状态转换
			if cnt.ConsecutiveSuccesses > defaultSuccessThreshold {
				currentState = StateClosed
				cnt.ConsecutiveFailures = 0
			}
			return nil

		case StateClosed:

		}
		return nil
	}
}
