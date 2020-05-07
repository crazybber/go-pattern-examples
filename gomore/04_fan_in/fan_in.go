package fanin

import "sync"

// Merge operate a FanIn to compose different channels into one
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int, 3)

	wg.Add(len(cs))

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	//启动多个 go routine 开始工作
	for _, c := range cs {
		go send(c)
	}
	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	//关闭动作,放在发送一方，会更好
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateNumbersPipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		//发送完成之后关闭
		close(out)
	}()
	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		//发送完成之后关闭
		close(out)
	}()
	return out
}
