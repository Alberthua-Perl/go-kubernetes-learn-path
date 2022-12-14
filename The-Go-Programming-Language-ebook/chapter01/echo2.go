package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// range return index and vaule of slice
		// use "_" to omit assign of index
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// output final s variable including all cli arguements
}