package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read error: ", err)
	}
	fmt.Print("result: ", result)
}
