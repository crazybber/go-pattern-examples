package fanout

import "sync"

// Split 重复分发数据为多份
// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []chan int {

	//get a queue of chan
	//cs := make([]chan int, n) //创建了个chan 数组但是空的

	cs := []chan int{}
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	//Distributes the work in a round robin fashion among the stated number of channels
	//until the main channel has been closed. In that case, close all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan int) {
		// Close every channel when the execution ends.
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		//this version will block
		for {
			//get a target from ch
			select {
			case val, ok := <-ch:
				if !ok {
					return // channel closed
				}
				//send value to all channels
				for _, c := range cs {
					c <- val
				}
			}
		}
	}

	// a worker to distribute message
	go distributeToChannels(ch, cs)

	return cs
}

// Split2  多工作者重复分发
// Split2 a channel into n channels that receive messages in a round-robin fashion.
func Split2(ch <-chan int, n int) []chan int {

	cs := make([]chan int, 0)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	distributeToChannels := func(ch <-chan int, cs []chan int) {
		// Close every channel when the execution ends.
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()
		var wg sync.WaitGroup
		for {
			//get a target from ch
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}
				wg.Add(1)
				go func(v int) {
					defer wg.Done()
					//send value to all channels
					for _, c := range cs {
						c <- val
					}

				}(val)

			}
		}

		wg.Wait()
	}

	go distributeToChannels(ch, cs)

	return cs
}

//Split3  随机分发到不同的目的地
//Split3 a channel into n channels that receive messages in a round-robin fashion.
func Split3(ch <-chan int, n int) []chan int {

	cs := make([]chan int, 0)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan int) {
		// Close every channel when the execution ends.
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}

//The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
// The gen function starts a goroutine that sends the integers on the channel and closes the channel when all
// the values have been sent:
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

//The second stage, sq, receives integers from a channel and returns a channel that emits the square of
// each received integer. After the inbound channel is closed and this stage has sent all the values downstream,
// it closes the outbound channel:
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
