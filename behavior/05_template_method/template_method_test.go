package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {

	//打印机
	aprinter := Printer{}

	//这个是被复合的work流程
	aprinter.DoPrintWork()

	//连接PDF打印机
	aprinter.printer = &PDF{output: "./home"}

	aprinter.Set("---PDF--")
	//打印
	aprinter.DoPrintWork()

	//连接纸质打印机
	aprinter.printer = &DevicePrinter{quality: 5}

	aprinter.Set("---Paper--")
	//打印
	aprinter.DoPrintWork()

}
