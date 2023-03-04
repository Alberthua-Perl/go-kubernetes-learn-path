package main

import (
	"fmt"
	"./exportedIdentifierDemo"
	// 该自定义包也可位于包默认存储目录 /usr/local/go/src/ 中
	// 设置 GO111MODULE 环境变量为 off 可直接使用相对路径指定的包
)

func main() {
	fmt.Printf("调用自定义包的 A_string: %s\n", exportedIdentiferDemo.A_string)
}
