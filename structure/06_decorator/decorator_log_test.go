package decorator

import (
	"log"
	"testing"
)

//装饰一个函数，或者说包装一个函数
type FuncToDecorate func(int) int

func TestDecorateLog(t *testing.T) {
	f := LogDecorate(Double)
	f(5)
}

func Double(n int) int {
	return n * 2
}

func LogDecorate(fn FuncToDecorate) FuncToDecorate {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
