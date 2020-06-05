/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-05-21 12:14:27
 * @Last Modified by: Edward
 * @Last Modified time: 2020-05-21 12:27:56
 */

package ratelimit

import (
	"testing"
)

/*

Rate limiting is an very important mechanism
With limiting you can controll resource utilization and maintain quality of service.
Go  supports rate limiting by using goroutines, channels, and tickers.
*/

var (
	requestQueueSize = 10
)

func TestRateLimiting(t *testing.T) {

	//请求队列
	burstyRequests := make(chan int, requestQueueSize)

	for i := 1; i <= requestQueueSize; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//对请求进行限流

	//200ms允许一次请求,最多同时3个突发
	rateLimiting(burstyRequests, 200, 3)

}
