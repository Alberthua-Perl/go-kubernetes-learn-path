package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(os.Args[0])  // slice output
	fmt.Println(strings.Join(os.Args[1:], " "))
	// effective method to output cli arguements
	// use fmt, strings and os packages
}