// 布尔值的使用
package main

import (
	"fmt"
)

func main() {
	var n int = 10
	r := n == 20 // 先比较判断返回布尔值，再进行短变量声明赋值。
	fmt.Printf("10 是否等于 20：%v\n", r)

	if n > 0 {
		fmt.Printf("if 的条件判断为 true\n")
	}
	for n > 0 {
		fmt.Printf("for 的条件判断为 true\n")
		break // 使用 break 关键字终止循环，否则进入无限循环。
	}
}
