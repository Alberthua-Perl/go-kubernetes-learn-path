package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// os.Args: slice type variable
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	// output final s variable including all cli arguements
}