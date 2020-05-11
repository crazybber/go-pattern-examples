package circuit

import "time"

//Options for breaker
type Options struct {
	Name              string
	Expiry            time.Time
	Interval, Timeout time.Duration
	MaxRequests       uint32
	ReadyToTrip       StateCheckerHandler
	OnStateChanged    StateChangedEventHandler
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

//OnStateChanged set handle of ChangedHandle
func OnStateChanged(handler StateChangedEventHandler) Option {
	return func(opts *Options) {
		opts.OnStateChanged = handler
	}
}

//ReadyToTrip check traffic state ,to see if request can go
func ReadyToTrip(readyToGo StateCheckerHandler) Option {
	return func(opts *Options) {
		opts.ReadyToTrip = readyToGo
	}
}
