package fanout

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func run() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	var errcList []chan error

	//	res := fanin(fanout(generateNumbers(ctx), squareNumbers))

	nChan, errc, err := generateNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	errcList = append(errcList, errc)

	fChanl, errcl, err := fanout(ctx, nChan, squareNumbers)
	errcList = append(errcList, errcl...)

	res, errRes, err := fanin(fChanl, errcl)

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
	errc := make(chan error, 100)

	go func() {
		defer close(out)
		defer close(errc)

		for i := 0; i < 10; i++ {
			fmt.Println("gen: ", i)
			out <- i
		}
	}()

	return out, errc, nil
}

func fanout(ctx context.Context, in chan interface{}, fct Op) ([]chan interface{}, []chan error, error) {
	fmt.Println("fanout called")
	var out []chan interface{}
	var errcl []chan error

	for i := 0; i < 5; i++ {
		res, errc, err := fct(ctx, in)
		// Todo: manage error
		_ = err
		out = append(out, res)
		errcl = append(errcl, errc)
	}

	return out, errcl, nil
}

func squareNumbers(ctx context.Context, in chan interface{}) (chan interface{}, chan error, error) {
	fmt.Println("squareNumbers called")
	out := make(chan interface{}, 100)
	errc := make(chan error, 100)

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

	return out, errc, nil
}

func fanin(ins []chan interface{}, errcl []chan error) (chan interface{}, chan error, error) {
	fmt.Println("fanin called")
	out := make(chan interface{}, 100)
	errcout := make(chan error, 100)
	var waitgroup sync.WaitGroup

	length := len(ins) + len(errcl)
	waitgroup.Add(length)

	for _, v := range ins {
		go func(w chan interface{}) {
			for val := range w {
				out <- val
			}
			waitgroup.Done()
		}(v)
	}

	for _, e := range errcl {
		go func(errc chan error) {
			for v := range errc {
				errcout <- v
			}
			waitgroup.Done()
		}(e)
	}

	go func() {
		waitgroup.Wait()
		close(out)
		close(errcout)
	}()

	return out, errcout, nil
}
func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	var errcList []chan error

	//	res := fanin(fanout(generateNumbers(ctx), squareNumbers))

	nChan, errc, err := generateNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	errcList = append(errcList, errc)

	fChanl, errcl, err := fanout(ctx, nChan, squareNumbers)
	errcList = append(errcList, errcl...)

	res, errRes, err := fanin(fChanl, errcl)

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
	errc := make(chan error, 100)

	go func() {
		defer close(out)
		defer close(errc)

		for i := 0; i < 10; i++ {
			fmt.Println("gen: ", i)
			out <- i
		}
	}()

	return out, errc, nil
}

func fanout(ctx context.Context, in chan interface{}, fct Op) ([]chan interface{}, []chan error, error) {
	fmt.Println("fanout called")
	var out []chan interface{}
	var errcl []chan error

	for i := 0; i < 5; i++ {
		res, errc, err := fct(ctx, in)
		// Todo: manage error
		_ = err
		out = append(out, res)
		errcl = append(errcl, errc)
	}

	return out, errcl, nil
}

func squareNumbers(ctx context.Context, in chan interface{}) (chan interface{}, chan error, error) {
	fmt.Println("squareNumbers called")
	out := make(chan interface{}, 100)
	errc := make(chan error, 100)

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

	return out, errc, nil
}

func fanin(ins []chan interface{}, errcl []chan error) (chan interface{}, chan error, error) {
	fmt.Println("fanin called")
	out := make(chan interface{}, 100)
	errcout := make(chan error, 100)
	var waitgroup sync.WaitGroup

	length := len(ins) + len(errcl)
	waitgroup.Add(length)

	for _, v := range ins {
		go func(w chan interface{}) {
			for val := range w {
				out <- val
			}
			waitgroup.Done()
		}(v)
	}

	for _, e := range errcl {
		go func(errc chan error) {
			for v := range errc {
				errcout <- v
			}
			waitgroup.Done()
		}(e)
	}

	go func() {
		waitgroup.Wait()
		close(out)
		close(errcout)
	}()

	return out, errcout, nil
}
