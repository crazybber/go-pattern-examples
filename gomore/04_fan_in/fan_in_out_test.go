package fanin

import (
	"fmt"
	"testing"
)

func TestFanInOutNumbersSeq(t *testing.T) {

	//一路输入源
	dataStreams := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan1 := generateNumbersPipeline(dataStreams)
	inputChan2 := generateNumbersPipeline(dataStreams)

	//this is a fanout operation
	// Fan-out to 2 Go-routine
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
