package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

/*

Rate limiting is an very important mechanism
With limiting you can controll resource utilization and maintain quality of service.
Go  supports rate limiting by using goroutines, channels, and tickers.
*/

func TestRateLimiting(t *testing.T) {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		t.Log("Sev request by 2000 Millisecond", req, time.Now())
	}

	burstyLimiter := make(chan struct{}, 3)

	//init burstyLimiter
	for i := 0; i < 3; i++ {
		burstyLimiter <- struct{}{}
	}
	go func() {
		for {
			select {
			case <-time.Tick(200 * time.Millisecond):
				burstyLimiter <- struct{}{}
			}

		}
	}()

	//max request queue
	burstyRequestsQueue := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequestsQueue <- i
	}
	close(burstyRequestsQueue)

	for req := range burstyRequestsQueue {
		<-burstyLimiter
		if len(burstyLimiter) > 0 {
			fmt.Println("working current in bursting status!")
		} else {
			fmt.Println("working current in normal status!")
		}
		fmt.Println("request handled", req, time.Now())
	}

	//rateLimiting()

}
