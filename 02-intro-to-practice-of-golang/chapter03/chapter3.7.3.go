package main

import (
	"fmt"
	"strings"
	"bytes"
)

func main() {
	n := "hello world"
	m := "I am alberthua-perl"

	// '+' 实现字符串连接
	j := n + "," + m
	fmt.Println(j)

	// fmt.Sprintf() 函数实现字符串连接
	k := fmt.Sprintf("%s%s%s", n, ",", m)
	fmt.Println(k)

	// Golang 官方推荐使用的方法
	var builder strings.Builder
	builder.WriteString(n)
	builder.WriteString(",")
	builder.WriteString(m)
	fmt.Println(builder.String())

	// 使用 bytes.Buffer 结构体并调用其方法以实现
	var buffer bytes.Buffer
	buffer.WriteString(n)
	buffer.WriteString(",")
	buffer.WriteString(m)
	fmt.Println(buffer.String())
}
