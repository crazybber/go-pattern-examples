package deadline

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func workerTakes5ms(stopper chan error) error {
	fmt.Println("i'm doing this job in 5ms")
	time.Sleep(5 * time.Millisecond)
	return nil
}

func workerTakes20ms(stopper chan error) error {
	fmt.Println("i'm doing this job in 20ms,so work will timeout")
	time.Sleep(20 * time.Millisecond)
	return nil
}

func cancelWork(stopper chan error) error {
	fmt.Println("i'm doing this job")
	stopper <- errors.New("canceled job") //cancel job
	time.Sleep(5 * time.Millisecond)
	fmt.Println("job canceled")
	return nil
}

func returnsError(stopper chan error) error {
	fmt.Println("i'm doing this job but error occurred")
	return errors.New("foo")
}

func TestMultiDeadline(t *testing.T) {

	dl := New(15*time.Millisecond, "test multi deadline case")

	if err := dl.Run(workerTakes5ms); err != nil {
		t.Error(err)
	}

	err := dl.Run(cancelWork)

	t.Log("cancelWork  error:", err)

	if err.Error() != "canceled job" {
		t.Error(err)
	}

	err = dl.Run(workerTakes20ms)

	if err != ErrTimedOut {
		t.Error(err)
	}
}

func TestDeadline(t *testing.T) {

	dl := New(1*time.Second, "one time deadline case worker")

	done := make(chan error)

	err := dl.Run(func(stopper chan error) error {
		fmt.Println("i am doing something here")
		time.Sleep(time.Second * 2)
		close(done)
		return nil
	})

	if err != ErrTimedOut {
		t.Error(err)
	}
	<-done

}
