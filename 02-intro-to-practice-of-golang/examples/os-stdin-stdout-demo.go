// Go 语言处理标准输入的多种方式
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 方式 1：使用 os.Stdin.Read() 函数
	var buffer [512]byte               // 声明字节类型的数组变量
	n, err := os.Stdin.Read(buffer[:]) // 返回标准输入字节的数量
	if err != nil {
		fmt.Println("read error: ", err)
	}
	// string() 方法返回标准输入
	fmt.Print("标准输入字节数：", n, ", 标准输入内容：", string(buffer[:]))
	fmt.Println()

	// 方式 2：使用 bufio.NewReader() 方法创建结构体
	var reader *bufio.Reader // 声明 *bufio.Reader 类型的变量
	fmt.Print("标准输入内容：")
	reader = bufio.NewReader(os.Stdin)     // 创建 reader 结构体对象
	result, err := reader.ReadString('\n') // 以 '\n' 作为一行的分隔符
	if err != nil {
		fmt.Println("read error: ", err)
	}
	//fmt.Print("result: ", result)
	fmt.Print("确认标准输入内容：", result)
	fmt.Println()

	// 方式 3：使用 fmt.Scan() 函数与 fmt.Scanln() 函数
	var name, lang, platform string // 若声明为 rune 数据类型将导致输入失败，因此使用 string 数据类型。
	var age int
	fmt.Print("请输入您的姓名：")
	fmt.Scan(&name)
	/*
	   fmt.Scan() 函数等待用户的标准输入，否则始终处于阻塞状态。
	   fmt.Scan() 函数可接收多个参数，用户输入参数默认使用空格或者回车换行符分割输入设备传入的参数，直到接收所有的参数为止。
	*/
	fmt.Print("请输入您的年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入喜欢的语言与系统：")
	// fmt.Scanln() 函数也可以接收多个参数，用户输入参数默认使用空格分割输入设备传入的参数，遇到回车换行符结束接收参数。
	fmt.Scanln(&lang, &platform)
	fmt.Print("您的姓名：", name, " 您的年龄：", age, "\n")
	fmt.Print("喜欢的语言与系统：", lang, ", ", platform, "\n")
	// fmt.Println() 函数将在参数的两边添加空格，而 fmt.Print() 函数不添加。
}
