package fanout

import (
	"fmt"
	"sync"
	"testing"
)

func TestMultiFanOutNumbersSeq(T *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	//generator integer stream
	inputChan := gen(dataStreams...)

	// transfer to
	ch := sq(inputChan)

	//	split it to 3 channel
	// 重复分发
	outArray := Split(ch, 3)

	var wg sync.WaitGroup
	wg.Add(len(outArray))
	for i := 0; i < len(outArray); i++ {

		go func(in <-chan int, index int) {
			sum := 0
			for item := range in {
				sum += item
			}
			fmt.Println("worker:", index, sum)

			wg.Done()
		}(outArray[i], i)
	}
	wg.Wait()
}

func TestManualFanOutNumbersSeq(T *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan1 := gen(dataStreams...)
	inputChan2 := gen(dataStreams...)

	//Manual Fan-out to 2 Go-routine
	c1 := sq(inputChan1)
	c2 := sq(inputChan2)

	fmt.Print("c1 queue: ")
	for n := range c1 {
		fmt.Print(n, " ")
	}
	fmt.Println()

	fmt.Print("c2 queue: ")
	for n := range c2 {
		fmt.Print(n, " ")
	}
	fmt.Println()

}
