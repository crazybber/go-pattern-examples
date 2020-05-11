package ratelimit

import (
	"fmt"
	"time"
)

/*

Rate limiting is an very important mechanism
With limiting you can controll resource utilization and maintain quality of service.
Go  supports rate limiting by using goroutines, channels, and tickers.
*/

func rateLimiting(requestQueueSize, allowedBurstCount int) {

	requests := make(chan int, requestQueueSize)
	for i := 1; i <= requestQueueSize; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//允许的突发数量
	burstyLimiter := make(chan time.Time, allowedBurstCount)

	//初始化允许突发的数量
	for i := 0; i < allowedBurstCount; i++ {
		burstyLimiter <- time.Now()
	}

	//控制请求频率的计时器
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//请求队列
	burstyRequests := make(chan int, requestQueueSize)
	for i := 1; i <= requestQueueSize; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
