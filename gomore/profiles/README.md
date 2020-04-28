# 时间差模式

算是一个技巧：

`defer`后面跟的函数参数会被第一时间计算并存储在函数计算过程本地
//当函数在被调用时会利用前面已经存储的值进行计算

  ```go
  func YourMainFunction(input SomeType) error {

	defer YourFunc(time.Now(), ....) //一定要放在函数的第一行或者你想要计算时差的代码的前面.

  //其他代码..


}
  ```