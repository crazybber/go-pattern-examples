package circuit

import "time"

//Options for breaker
type Options struct {
	Name           string
	Expiry         time.Time
	Interval       time.Duration
	Timeout        time.Duration
	MaxRequests    uint32
	ReadyToTrip    StateCheckerHandler
	OnStateChanged StateChangedEventHandler
}

//SetName of breaker
func SetName(name string) Option {
	return func(opts *Options) {
		opts.Name = name
	}
}

//SetExpiry of breaker
func SetExpiry(expiry time.Time) Option {
	return func(opts *Options) {
		opts.Expiry = expiry
	}
}

//SetStateChangedHandle set handle of ChangedHandle
func SetStateChangedHandle(handler StateChangedEventHandler) Option {
	return func(opts *Options) {
		opts.OnStateChanged = handler
	}
}

//SetReadyToTrip check traffic state ,to see if request can go
func SetReadyToTrip(readyToGo StateCheckerHandler) Option {
	return func(opts *Options) {
		opts.ReadyToTrip = readyToGo
	}
}
