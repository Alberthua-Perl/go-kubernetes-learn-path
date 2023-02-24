package main

import "fmt"

const boilingF = 212.0
// The const can be accessed from other files from package.

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)	
}
