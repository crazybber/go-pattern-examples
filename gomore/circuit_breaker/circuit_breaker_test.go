package circuit

import (
	"context"
	"errors"
	"time"
)

var (
	ErrServiceUnavailable = errors.New("Service Unavailable")
)

type State int

const (
	UnknownState State = iota
	FailureState
	SuccessState
)

//Counter interface
type Counter interface {
	Count(State)
	ConsecutiveFailures() uint32
	LastActivity() time.Time
	Reset()
}

type counters struct {
	state        State
	lastActivity time.Time
}

func (c *counters) Count(State) {

}

func (c *counters) ConsecutiveFailures() uint32 {

	return 0
}

func (c *counters) LastActivity() time.Time {
	return c.lastActivity
}

func (c *counters) Reset() {

}

func NewCounter() Counter {
	var i Counter
	return i
}

type Circuit func(context.Context) error

func Breaker(c Circuit, failureThreshold uint32) Circuit {

	cnt := NewCounter()

	return func(ctx context.Context) error {
		if cnt.ConsecutiveFailures() >= failureThreshold {
			canRetry := func(cnt Counter) bool {
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
