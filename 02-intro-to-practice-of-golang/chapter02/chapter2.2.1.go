// 变量声明的多种方式
package main

import (
	"fmt"
)

func main() {
	var a int
	a = 10
	// 先定义整型变量再赋值

	var b int = 10
	// 同时定义整型变量并赋值

	var (
		//c int
		//d string
		// 若变量已定义，但在后续程序中未被使用，编译器将报错。
		_ string
		// _ 空白标识符省略字符串类型的变量声明而不报错
		e int = 10
	)
	// 批量定义变量

	var f, g int
	f, g = 10, 10
	// 批量定义同一类型的变量再赋值
	f = f + g

	h := 10
	// := 短变量声明常用于函数内或局部变量的声明

	fmt.Printf("定义变量，不设置初始值：%d\n", a)
	fmt.Printf("定义变量并设置初始值：%d\n", b)
	fmt.Printf("批量定义变量，可根据决定情况是否设置初始值：%d\n", e)
	fmt.Printf("多个变量同一数据类型：%d\n", f)
	fmt.Printf("定义变量并赋值，通过值设置变量的数据类型，即短变量声明：%d\n", h)
}
