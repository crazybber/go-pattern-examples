package fanin

import (
	"fmt"
	"testing"
)

func TestFanIn(T *testing.T) {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan := generatePipeline(randomNumbers)

	// Fan-out to 2 Go-routine
	c1 := squareNumber(inputChan)
	c2 := squareNumber(inputChan)

	// Fan-in the resulting squared numbers
	c := fanIn(c1, c2)
	sum := 0

	// Do the summation
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("Total Sum of Squares: %d", sum)
}
