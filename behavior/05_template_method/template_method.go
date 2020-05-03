package templatemethod

import "fmt"

////////////////////////////////
//使用打印的例子

//IPrinter 定义打印的流程
type IPrinter interface {
	Set(mark string)
	Print()
}

//Printer 定义基本结构类型
type Printer struct {
	workerMark string
	printer    IPrinter //指项实际工作的类型
}

//LoadDrive 载入驱动
func (p *Printer) LoadDrive() {
	fmt.Print("init print drive\n")
}

//UnLoadDrive 卸载驱动
func (p *Printer) UnLoadDrive() {
	fmt.Print("unload drive\n")
}

//Set 设置参数，这是变化的部分
func (p *Printer) Set(mark string) {
	p.workerMark = mark
	//调用实现
	if p.printer != nil {
		p.printer.Set(mark)
	}
}

//Print 执行打印，这是变化的部分
func (p *Printer) Print() {
	//调用实现
	fmt.Print("print with task mark: ", p.workerMark, "\n")
	if p.printer != nil {
		p.printer.Print()
	}

}

//DoPrintWork 打印
//DoPrintWork 定义了打印的流程
func (p *Printer) DoPrintWork() {
	p.LoadDrive()
	p.Set(p.workerMark)
	p.Print()
	p.UnLoadDrive()
}

//PDF 虚拟打印
type PDF struct {
	Printer
	output string
}

//Print to a PDF
func (p *PDF) Print() {
	fmt.Print("print to PDF ,save to ", p.output, "\n")

}

//DevicePrinter 设备打印机
type DevicePrinter struct {
	Printer
	quality int //1,2,3表示打印高中低
}

//Print to a Paper
func (d *DevicePrinter) Print() {
	fmt.Print("print to Paper ,with quality: ", d.quality, "\n")
}
