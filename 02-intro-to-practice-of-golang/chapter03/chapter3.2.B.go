// 整型数与浮点数的类型转换
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f32 float32 = 1.1
	var f64 float64 = 2.2

	r := float64(f32) + f64 
	/* 使用 float64() 内置函数转换为 64 位浮点数
	   十进制的浮点数转换为二进制数后将导致精度丢失，可引入第三方包 decimal 解决。
	*/
	fmt.Printf("浮点数运算结果为：%v\n", r)
	n, err := strconv.ParseFloat(fmt.Sprintf("%.1f", r), 64)
	if err != nil {
		fmt.Println("字符串转换浮点数失败！")
	}
	fmt.Printf("浮点数精度处理结果为：%v\n", n)

	var i int = 10
	rd := float32(i) + f32 // 使用 float32() 内置函数转换为 32 位浮点数
	fmt.Printf("浮点数运算结果为：%v\n", rd)
}
