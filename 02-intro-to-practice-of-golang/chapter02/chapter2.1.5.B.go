package main

import (
	"fmt"
	"exportedIdentifierDemo"
	// 该自定义包位于包默认存储目录 /usr/local/go/src/ 中
)

func main() {
	fmt.Printf("调用自定义包的 A_string: %s\n", exportedIdentiferDemo.A_string)
}
