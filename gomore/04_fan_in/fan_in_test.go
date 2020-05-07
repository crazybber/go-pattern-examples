package fanin

import (
	"fmt"
	"testing"
)

func TestFanInNumbersSeq(T *testing.T) {

	//第一路输入源
	dataStreams1 := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan1 := generateNumbersPipeline(dataStreams1)

	//第二路输入源
	dataStreams2 := []int{2, 4, 6, 9, 1, 1, 2, 3, 7, 8}

	inputChan2 := generateNumbersPipeline(dataStreams2)

	c1 := squareNumber(inputChan1)

	c2 := squareNumber(inputChan2)

	//fanIn data for the squared numbers
	out := Merge(c1, c2)

	sum := 0

	for c := range out {
		sum += c
	}

	fmt.Printf("Total Sum of Squares by FanIn : %d\n", sum)
}
