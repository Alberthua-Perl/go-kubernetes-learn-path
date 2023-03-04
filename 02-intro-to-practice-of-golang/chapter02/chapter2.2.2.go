// 常量声明的多种方式与 iota 特殊常量
package main

import (
	"fmt"
)

func main() {
	// 常量定义的多种方式
	const a int = 10
	const b = 10
	// 定义并赋值整型常量
	const (
		c int = 10
		d string = "golang"
	)
	// 定义并赋值整型与字符串常量
	//const e, f int string = 20, "rock"
	const e, f = 20, "rock"

	fmt.Printf("单个常量的定义方式一：%d\n", a)
	fmt.Printf("单个常量的定义方式二：%d\n", b)
	fmt.Printf("多个常量的定义方式一：%d, %s\n", c, d)
	fmt.Printf("多个常量的定义方式二：%d, %s\n", e, f)

	// iota 特殊变量的使用示例
	const (
		u = iota  // iota 重置为 0
		v         // iota 累加为 1，常量的值为之前常量定义的值。
		w = 10    // iota 累加为 2
		x         // iota 累加为 3，常量的值为之前常量定义的值。
		y = iota  // iota 累加为 4
	)
	fmt.Printf("u 的值为：%d\n", u)
	fmt.Printf("v 的值为：%d\n", v)
	fmt.Printf("w 的值为：%d\n", w)
	fmt.Printf("x 的值为：%d\n", x)
	fmt.Printf("y 的值为：%d\n", y)
}
