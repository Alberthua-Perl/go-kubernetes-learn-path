// Go 中的字符类型：uint8 与 rune
package main

import (
	"fmt"
)

func main() {
	var a byte // uint8 数据类型为 byte 类型
	var b rune // rune 数据类型等价于 int32 数据类型

	a = 'u'
	b = 'a'
	c := 'c'
    // 格式化 %T 输出变量的数据类型
	fmt.Printf("变量 a 的数据类型为：%T\n", a)
	fmt.Printf("变量 b 的数据类型为：%T\n", b)
	fmt.Printf("变量 c 的数据类型为：%T\n", c)
}
