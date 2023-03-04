// & 与 * 在指针变量中的使用
package main

import (
	"fmt"
)

func main() {
	a := "golang"
	var ptr *string
	// 短变量声明字符串类型变量与声明字符串类型的指针
	fmt.Printf("a 变量的虚拟内存地址：%v\n", &a)
	ptr = &a
	// 指针即为变量的虚拟内存地址
	fmt.Printf("指针 ptr 的虚拟内存地址：%v\n", ptr)
	fmt.Printf("指针 ptr 的值：%s\n", *ptr)
	// 引用指针即为引用虚拟内存地址中的变量值
}
