/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-06-05 12:43:39
 * @Last Modified by: Edward
 * @Last Modified time: 2020-06-05 17:34:37
 */

// Package deadline implements deadline (also known as "timeout") resiliency pattern for Go.
package deadline

import (
	"errors"
	"time"
)

// ErrTimedOut is the error returned from Run when the Worker expires.
var ErrTimedOut = errors.New("timed out waiting for function to finish")

// Worker implements the Deadline/timeout resiliency pattern.
// worker do the target job
type Worker struct {
	timeout time.Duration
	action  string
}

// New create a new Worker with the given timeout.and tile
func New(timeout time.Duration, someActionTitle string) *Worker {
	return &Worker{
		timeout: timeout,
		action:  someActionTitle,
	}
}

// Run runs the given function, passing it a stopper channel. If the Worker passes before
// the function finishes executing, Run returns ErrTimeOut to the caller and closes the stopper
// channel so that the work function can attempt to exit gracefully. It does not (and cannot)
// simply kill the running function, so if it doesn't respect the stopper channel then it may
// keep running after the Worker passes. If the function finishes before the Worker, then
// the return value of the function is returned from Run.
func (d *Worker) Run(work func(stopperSignal chan error) error) error {

	result := make(chan error)
	//we can stop the work in advance
	stopper := make(chan error, 1)

	go func() {
		value := work(stopper)
		select {
		case result <- value:
		case stopError := <-stopper:
			result <- stopError
		}
	}()

	//handle result

	select {
	case ret := <-result:
		return ret
	case <-time.After(d.timeout):
		close(stopper)
		return ErrTimedOut
	}
}
