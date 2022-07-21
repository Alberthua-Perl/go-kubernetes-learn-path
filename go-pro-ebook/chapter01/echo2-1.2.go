package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		// range return index and vaule of slice
		fmt.Println(index, arg)
		// output index and cli arguements each line
	}
}