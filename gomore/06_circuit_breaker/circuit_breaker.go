package circuit

import (
	"context"
	"errors"
	"time"
)

//ErrServiceUnavailable for error
var (
	ErrTooManyRequests    = errors.New("too many requests")
	ErrServiceUnavailable = errors.New("service unavailable")
	FailureThreshold      = 10
)

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
func NewCounter() Counter {
	return &counters{}
}

//Breaker of circuit
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
