/* 
   浮点数类型与精度丢失
   建议：实际开发中使用 float64 浮点数，其精度较 float32 更高。
*/
package main

import (
	"fmt"
	"math"
)

func getMaxFloat() {
	fmt.Println("支持的最大 float32 与 float64 浮点数为：")
	fmt.Println(" ", math.MaxFloat32)
	fmt.Println(" ", math.MaxFloat64)
}

func main() {
	getMaxFloat() // 获取最大的 float32 与 float64 的浮点数

	/* 使用 math 包获取支持的最小与最大整型数
	   其中无符号整形数的最小值为 0
	*/
	fmt.Printf("支持的最小与最大 int8 整型数为：%v, %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("支持的最小与最大 int16 整型数为：%v, %v\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("支持的最小与最大 int32 整型数为：%v, %v\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("支持的最小与最大 int64 整型数为：%v, %v\n", math.MinInt64, math.MaxInt64)

	fmt.Printf("支持的最大 uint8 无符号整型数为：%v\n", math.MaxUint8)
	fmt.Printf("支持的最大 uint16 无符号整型数为：%v\n", math.MaxUint16)
	fmt.Printf("支持的最大 uint32 无符号整型数为：%v\n", math.MaxUint32)

	var umax64 uint64
	umax64 = 1<<64 - 1 // 位运算赋值
	fmt.Printf("支持的最大 uint64 无符号整型数为：%v\n", umax64)
}
