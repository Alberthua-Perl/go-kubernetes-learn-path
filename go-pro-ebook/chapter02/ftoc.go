package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	// declare two const at the same time
	fmt.Printf("%g F = %g C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToc(boilingF))

	// just declare variable through type
	var s string
	fmt.Printf("Print vaule of s: \n", s)

	// declare variable list through type or expression
	var i, j, k int  // int, int, int
	var b, g, p = true, 2.3, "four"  // bool, float64, string
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
	// use return keyword
}