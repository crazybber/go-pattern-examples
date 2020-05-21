package semaphore

import (
	"errors"
	"time"
)

//error info
var (
	ErrNoTickets      = errors.New("could not acquire semaphore")
	ErrIllegalRelease = errors.New("can't release the semaphore without acquiring it first")
)

// ISemaphore contains the behavior of a semaphore that can be acquired and/or released.
type ISemaphore interface {
	Acquire() error
	Release() error
}

type semp struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *semp) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *semp) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}

}

//New return a new Semaphore
func New(tickets int, timeout time.Duration) ISemaphore {
	return &semp{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
