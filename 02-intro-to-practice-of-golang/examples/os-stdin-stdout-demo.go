// Go 语言处理标准输入的多种方式
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
    // 方式 1：使用 os.Stdin.Read() 函数
	var buffer [512]byte // 声明数组变量
	n, err := os.Stdin.Read(buffer[:]) // 返回标准输入字节的数量
	if err != nil {
		fmt.Println("read error: ", err)
	}
	// string() 方法返回标准输入
	fmt.Print("标准输入字节数：", n, ", 标准输入内容：", string(buffer[:]))
	fmt.Println()

	// 方式 2：使用 bufio.NewReader() 方法创建结构体
	var reader *bufio.Reader
	fmt.Print("请输入您的邮箱：")
	reader = bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read error: ", err)
	}
	//fmt.Print("result: ", result)
	fmt.Print("您的邮箱：", result)
}
