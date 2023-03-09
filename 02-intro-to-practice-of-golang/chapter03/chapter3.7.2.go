package main

import (
	"fmt"
)

func main() {
	var n rune
	n = '你'
	p := &n // 指针 p 的声明与赋值，其值为 n 变量的虚拟内存地址。
	q := &p // 指针 q 的声明与赋值，其值为指针 p 的虚拟内存地址。

	fmt.Printf("格式化符号 v：%+v\n", 'a')
	fmt.Printf("格式化符号 T：%T\n", 123)
	fmt.Printf("格式化符号 q：%q\n", "Hello go")
	fmt.Printf("格式化符号 s：%s\n", "Hello go")
	fmt.Printf("格式化符号 f：%.2f\n", 3.14159)
	fmt.Printf("格式化符号 d：%d\n", 3121)
	fmt.Printf("格式化符号 p：%p\n", p)
	fmt.Printf("指针 p 的虚拟内存地址：%p\n", q)
	fmt.Printf("UTF-8 字符 a 的虚拟内存地址：%v\n", &n)
}
