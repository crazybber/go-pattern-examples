package fanout

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestFanInOut(T *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	var errList []chan error

	nChan, err, err := generateNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	errList = append(errList, err)

	fChannel, errList, err := fanout(ctx, nChan, squareNumbers)
	errList = append(errList, errList...)

	res, errRes, err := fanin(fChannel, errList)

	go func() {
		for {
			select {
			case r := <-res:
				v, ok := r.(int)
				if !ok {
					// TODO
					_ = ok
				}
				fmt.Println(v)
			case e := <-errRes:
				fmt.Println(e)
			}
		}
	}()

	fmt.Println("finished")
	time.Sleep(10 * time.Duration(time.Second))
}

type Op func(context.Context, chan interface{}) (chan interface{}, chan error, error)

func generateNumbers(ctx context.Context) (chan interface{}, chan error, error) {
	fmt.Println("generateNumbers called")
	out := make(chan interface{}, 100)
	err := make(chan error, 100)

	go func() {
		defer close(out)
		defer close(err)

		for i := 0; i < 10; i++ {
			fmt.Println("gen: ", i)
			out <- i
		}
	}()

	return out, err, nil
}

func fanout(ctx context.Context, in chan interface{}, fct Op) ([]chan interface{}, []chan error, error) {
	fmt.Println("fanout called")
	var out []chan interface{}
	var errList []chan error

	for i := 0; i < 5; i++ {
		res, err, err := fct(ctx, in)
		// Todo: manage error
		_ = err
		out = append(out, res)
		errList = append(errList, err)
	}

	return out, errList, nil
}

func squareNumbers(ctx context.Context, in chan interface{}) (chan interface{}, chan error, error) {
	fmt.Println("squareNumbers called")
	out := make(chan interface{}, 100)
	err := make(chan error, 100)

	go func() {
		for v := range in {
			num, _ := v.(int)
			squared := num * num
			select {
			case out <- squared:
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			}
		}
		close(out)
	}()

	return out, err, nil
}

func fanin(ins []chan interface{}, errList []chan error) (chan interface{}, chan error, error) {
	fmt.Println("fanin called")
	out := make(chan interface{}, 100)
	errout := make(chan error, 100)
	var waitgroup sync.WaitGroup

	length := len(ins) + len(errList)
	waitgroup.Add(length)

	for _, v := range ins {
		go func(w chan interface{}) {
			for val := range w {
				out <- val
			}
			waitgroup.Done()
		}(v)
	}

	for _, e := range errList {
		go func(err chan error) {
			for v := range err {
				errout <- v
			}
			waitgroup.Done()
		}(e)
	}

	go func() {
		waitgroup.Wait()
		close(out)
		close(errout)
	}()

	return out, errout, nil
}
