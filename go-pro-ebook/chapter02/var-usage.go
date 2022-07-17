package main

import (
	"fmt"
)

func main() {
	declaration()
	pointer()
}

func declaration() {
	// include var and short variable list
	// just declare variable through type
	var s string
	fmt.Printf("Print vaule of s: %s\n", s)
	
	// declare variable list through type or expression
	var i, j, k int
	// int, int, int
	var b, g, p = true, 2.3, "four"
	// bool, float64, string
	fmt.Printf("Print vaule of i, j, k, b, g, p: %d, %d, %d, %t, %g, %s\n", i, j, k, b, g, p)
	
	// f, err := os.Open("./file1")
	// fmt.Print("Output passwd file as following: f, err\n")
}

func pointer() {
	s := 2
	p := &s
	fmt.Println("\npointer of s is: ", &s)
	fmt.Println("vaule of p is: ", p)
	fmt.Println("vaule of p point to variable: ", *p)
	*p = 3
	fmt.Println("\ncurrent vaule of s is: ", s)
	fmt.Println("current pointer of s is: ", &s)
	fmt.Println("current vaule of p is: ", p)
}	