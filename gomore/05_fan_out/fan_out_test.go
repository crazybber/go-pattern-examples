package fanout

import (
	"fmt"
	"sync"
	"testing"
)

//多工作者,重复分发
func TestFanOutDuplicateMultiWorkers(t *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	//generator integer stream
	inputChan := gen(dataStreams...)

	// transfer to
	ch := sq(inputChan)

	//	split it to 3 channel
	// 重复分发
	outArray := Split2(ch, 3)

	length := len(outArray)
	t.Log("length of out channel:", length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {

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

//单个工作者，重复分发
func TestFanOutDuplicate(t *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	//generator integer stream
	inputChan := gen(dataStreams...)

	// transfer to
	ch := sq(inputChan)

	//	split it to 3 channel
	// 重复分发
	outArray := Split(ch, 3)

	length := len(outArray)
	t.Log("length of out channel:", length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {

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

//随机分发
// worker: 2 11245
// worker: 0 14988
// worker: 1 10117
func TestFanOutRandom(t *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	//generator integer stream
	inputChan := gen(dataStreams...)

	// transfer to
	ch := sq(inputChan)

	//	split it to 3 channel
	// 重复分发
	outArray := Split3(ch, 3)

	length := len(outArray)
	t.Log("length of out channel:", length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {

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
