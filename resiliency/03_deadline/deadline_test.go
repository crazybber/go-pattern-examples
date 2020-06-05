package deadline

import (
	"errors"
	"testing"
	"time"
)

func takes5ms(stopper <-chan struct{}) error {
	time.Sleep(5 * time.Millisecond)
	return nil
}

func takes20ms(stopper <-chan struct{}) error {
	time.Sleep(20 * time.Millisecond)
	return nil
}

func returnsError(stopper <-chan struct{}) error {
	return errors.New("foo")
}

func TestMultiDeadline(t *testing.T) {
	dl := New(10*time.Millisecond, "test multi deadline case")

	if err := dl.Run(takes5ms); err != nil {
		t.Error(err)
	}

	if err := dl.Run(takes20ms); err != ErrTimedOut {
		t.Error(err)
	}

	if err := dl.Run(returnsError); err.Error() != "foo" {
		t.Error(err)
	}

	done := make(chan struct{})
	err := dl.Run(func(stopper <-chan struct{}) error {
		<-stopper
		close(done)
		return nil
	})
	if err != ErrTimedOut {
		t.Error(err)
	}
	<-done
}

func TestDeadline(t *testing.T) {
	dl := New(1*time.Second, "one dead line case")

	err := dl.Run(func(stopper <-chan struct{}) error {
		time.Sleep(time.Second * 10)
		return nil
	})

	switch err {
	case ErrTimedOut:
		t.Error("execution took too long, oops")
	default:
		// some other error
		t.Log("done")
	}
}
