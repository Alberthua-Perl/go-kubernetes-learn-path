// Note: pay attention to function return
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n *bool = flag.Bool("n", false, "omit trailing newline")
// flag.Bool: func Bool(name string, vaule bool, usage string) *bool
var sep *string = flag.String("s", " ", "separator")
// flag.String: func String(name string, vaule string, usage string) *string
// variables for package level
// access https://pkg.go.dev/flag@go1.16.6 to learn how to use flag package

func main() {
	// n := flag.Bool("n", false, "omit trailing newline")
	// sep := flag.String("s", " ", "separator")
	// local variables in function
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	// no newline output
	// fmt.Println(strings.Join(flag.Args(), *sep))
	// flag.Args: func Args() []string
	if !*n {
		fmt.Println()
		// output newline("\n") directly
	}
}
