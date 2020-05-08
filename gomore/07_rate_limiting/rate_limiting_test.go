package ratelimit

import (
	"testing"
)

/*

Rate limiting is an very important mechanism
With limiting you can controll resource utilization and maintain quality of service.
Go  supports rate limiting by using goroutines, channels, and tickers.
*/

func TestRateLimiting(t *testing.T) {

	rateLimiting(5, 3)

}
