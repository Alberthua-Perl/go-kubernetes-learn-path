// iota 用于数据格式转换
package main 

import (
	"fmt"
)

func main() {
	const (
		_ = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
	)

	fmt.Printf("B 转 KB 的进制为：%.0f\n", KB)
	fmt.Printf("B 转 MB 的进制为：%.0f\n", MB)
	fmt.Printf("B 转 GB 的进制为：%.0f\n", GB)
	fmt.Printf("B 转 TB 的进制为：%.0f\n", TB)
}
