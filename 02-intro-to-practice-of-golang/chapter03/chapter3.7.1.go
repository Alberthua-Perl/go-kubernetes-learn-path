// 双引号与反引号在字符串赋值中的区别
package main

import (
	"fmt"
)

func main() {
	var a string
	var b string

	a = "hello world\nGo"
	// 双引号的字符串赋值：支持转义，不支持字面量打印与多行打印。
	b = `hello
	     \n
		 world`
	// 反引号的字符串赋值：不支持转义，支持字面量打印与多行打印。

	fmt.Printf("双引号的字符串：%v\n", a)
	fmt.Printf("反引号的字符串：%v\n", b)
}
