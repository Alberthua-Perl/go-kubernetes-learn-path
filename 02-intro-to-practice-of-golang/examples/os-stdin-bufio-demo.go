// 使用带缓冲的 bufio 标准库处理标准输入与输出
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputReader *bufio.Reader // 声明 inputReader 指针指向 bufio.Reader 结构体
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	// 对应的方法为 func (*bufio.Reader).ReadString(delim byte) (string, error)，该示例中使用 S 表示分隔符（终止符），
	// 也可以用其它字符，如 '\n'。
	input, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("error report: %s\n", err)
	} else {
		fmt.Printf("The input was: %s\n", input)
	}
}

/*
   上例中，inputReader 是个指针，它指向一个 bufio 类的 Reader，然后在 main() 函数中通过 bufio.NewReader(os.Stdin) 创建了一个
   buffer reader, 并且 inputReader 指针指向该 buffer reader。
   bufio.NewReader() 构造器的原型是这样的 func NewReader(rd io.Reader) *Reader

   任何符合 io.Reader 接口的对象（即实现了 Read() 方法对象）都可以作为 bufio.NewReader() 里的参数，并返回一个新的带缓冲的 io.Reader,
   os.Stdin 符合这个条件。 这个带缓冲的 reader 有一个方法 ReadString(delim byte), 这个方法会一直读数据，直到遇到了指定的终止符，
   终止符将成为输入的一部分，一起放到 buffer 里去。ReadString 方法返回的是读到的字符串及 nil，当读到文件的末端时，将返回把读到的字符串及
   io.EOF，如果在读到结束时没有发现所指定的结束符（delim byte），将返回一个 err !=nil 。

   在上面的例子中，我们从键盘输入直到键入 "S"，屏幕是标准输出 os.Stdout，错误信息被写到 os.Stderr。
   大多情况下，os.Stderr 等同 os.Stdout。
   一般情况下，在 GO 的代码里，省略变量声明，而直接使用 ":=" 声明，如：inputReader := bufio.NewReader(os.Stdin)
   input ,err :=inputReader.ReadString('\n')
*/
