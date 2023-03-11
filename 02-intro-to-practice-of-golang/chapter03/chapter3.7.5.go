// 遍历字符串
package main

import (
	"fmt"
)

func main() {
	m := "你好, Go语言" // Go 语言中默认 UTF-8 编码
	n := []rune(m) // 短变量声明字符串切片

	// 遍历方式 1：使用 []rune() 切片实现遍历
	for i :=0; i < len(n); i++ {
		fmt.Printf("%c - %d\n", n[i], n[i])
	}
	fmt.Println()
	// 遍历方式 2：使用 range 方法直接遍历字符串
	for _, s := range m {
		fmt.Printf("%c - %d\n", s, s)
	}
}
