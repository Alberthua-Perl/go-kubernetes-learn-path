// 空白标识符 "_" 的三个使用场景
package main

import (
	"fmt"
	_ "net/http/pprof"
	// 场景1：_ 空白标识符只执行包的初始化函数 init()，不调用包的任何变量或函数。
)

func myfunc() (int, string) {
	a := 10
	b := "golang"
	return a, b
}

type Foo interface {
	Say()
}

type Dog struct {
	name string
}
func (d Dog) Say() {
	fmt.Println(d.name + " say hi")
}

func main() {
	a, _ := myfunc()
	// 场景2：省略 myfunc() 函数第二个字符串变量的返回，只返回第一个整型变量的值，空白标识符作为匿名占位符。
	fmt.Printf("只获得函数 myfunc 的第一个返回值: %d\n", a)

	var _ Foo = Dog{"black dog"}
	// 场景3：类型断言，即判断 Dog 结构体是否对 Foo 接口的调用。
}
