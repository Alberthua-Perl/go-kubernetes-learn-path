/*
   字符串拼接的 5 种方式：
   1. 操作符：+
   2. 标准库函数：fmt.Sprintf()
   3. 标准库结构体：strings.Builder、bytes.Buffer
   其中 Golang 官方推荐 strings.Builder 标准库结构体
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	n := "hello world"
	m := "I am alberthua-perl"

	/* '+' 操作符：
	   "," 必须使用使用双引号圈引以表示为字符串，若使用单引号则表示为字符，
	   将导致 "mismatched types" 报错。
	*/
	j := n + "," + m
	fmt.Println(j)

	// fmt.Sprintf() 标准库函数
	k := fmt.Sprintf("%s%s%s\n", n, ",", m)
	fmt.Println(k)

	// strings.Builder 标准库结构体
	var builder strings.Builder
	builder.WriteString(n)
	builder.WriteString(",")
	builder.WriteString(m)
	//fmt.Printf("builder 变量的数据类型：%T\n", builder)
	fmt.Println(builder.String())

	// bytes.Buffer 标准库结构体
	var buffer bytes.Buffer
	buffer.WriteString(n)
	buffer.WriteString(",")
	buffer.WriteString(m)
	fmt.Println(buffer.String())
}
