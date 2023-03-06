/* 整型数据的定义与取值
   建议：实际开发过程中使用 int 整型与 uint 无符号整型
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	var number1 int8
	var number2 int16
	var number3 int32
	var number4 int64
	var number5 uint
	// 声明不同类型的整型变量

	n = 15645645613456465
	number1 = 120
	number2 = 1314
	number3 = 30000
	number4 = 111111
	//number5 = -15645645613456465  无符号整型的报错：overflows uint
	number5 = 15645645613456465

	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", number1)
	fmt.Printf("%v\n", number2)
	fmt.Printf("%v\n", number3)
	fmt.Printf("%v\n", number4)
	fmt.Printf("%v\n", number5)
}
