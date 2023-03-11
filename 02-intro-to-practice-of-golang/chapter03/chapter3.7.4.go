// 字符串的长度
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	n := "Hello, golang"
	m := "你好, Go语言"
	// 获取字符串的真实长度
	fmt.Println("字符串 n 的真实长度：", len(n))
	fmt.Println("字符串 m 的真实长度：", len(m))
	// Go 语言中字母、空格与符号占一个 1 个字节，中文字符占 3 个字节。

	// 获取字符串的实际长度
	u := utf8.RuneCountInString(m)
	v := []rune(m)
	fmt.Println("字符串 m 的实际长度：", u)
	fmt.Println("字符串 m 的实际长度：", len(v))
}
