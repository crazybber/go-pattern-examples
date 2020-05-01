package main

import (
	"fmt"

	"github.com/logrusorgru/aurora" //这是一个控制台可以多种颜色输出的颜色库
)

func main() {

	startGo := letsGo()
	fmt.Sprintln(aurora.Green(startGo))
}

func letsGo() string {
	fmt.Println("start go!")

	return fmt.Sprintln(contents)
}
