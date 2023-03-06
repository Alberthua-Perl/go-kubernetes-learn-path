// 练习：编写程序用于生成个人 BMI 指数，判断身体健康状况。
package main

import (
	"fmt"
	//"os"
	//"bufio"  从 os.Stdin 标准输入或创建的文件句柄读取文件内容，也可输出内容。
	"strconv" // 数值与字符串的相互转换
)

func main() {
	var weight, height, result, bmi float64 // 批量声明 float64 数据类型的变量

	fmt.Printf("请输入您的体重(KG): ")
	fmt.Scan(&weight)
	/* fmt.Scan() 函数直到标准输入中传递参数前，始终处于阻塞状态。
	   fmt.Scan() 函数可以接受数值或字符串的指针变量
	*/
	//fmt.Printf("weight 的值：%v\n", weight)

	fmt.Printf("请输入您的身高(M): ")
	fmt.Scan(&height)
	//fmt.Printf("heigh 的值：%v\n", heigh)

	result = weight / (height * height) // BMI 指数的计算公式
	bmi, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	/* fmt.Sprintf() 函数返回字符串，因此使用 strconv.ParseFloat() 函数转换字符串为
	   float64 浮点数，并使用空白标识符省略 err 返回值。
	*/
	n, _ := fmt.Printf("您的 BMI 指数为：%v\n", bmi) // 省略 err 输出
	fmt.Printf("前一条标准输出的字节数为：%v bytes\n", n)

	// if 条件判断语句
	if bmi < 18.4 {
		fmt.Println("您的体重过轻")
	} else if bmi < 24 && bmi >= 18.5 {
		fmt.Println("您的体重正常")
	} else if bmi < 27 && bmi >= 24 {
		fmt.Println("您的体重过重")
	} else if bmi < 30 && bmi >= 27 {
		fmt.Println("您属于轻度肥胖")
	} else if bmi < 35 && bmi >= 30 {
		fmt.Println("您属于中度肥胖")
	} else if bmi >= 35 {
		fmt.Println("您属于重度肥胖")
	}
}

/*
func getStdin(f *os.File) string {
	stdin := bufio.NewReader(f)
	content, err := stdin.ReadString('\n')
	// ReadString 方法返回字符串的内容，而非数值，因此在数值计算中会由于数据类型问题而报错。
	if err != nil {
		fmt.Println("read from stdin: err!")
	}
	return content
}
*/

/* 请区分以下 fmt 包中函数的使用方法：
   fmt.Print()：打印标准输出
   fmt.Println()：打印标准输出并追加换行符
   fmt.Printf()：打印格式化标准输出
   fmt.Sprintf()：打印字符串的格式化标准输出并返回字符串
   fmt.Fprint()：打印格式化标准输出并返回字符串

   以上函数均可返回实际输出的 bytes 字节数。
*/
