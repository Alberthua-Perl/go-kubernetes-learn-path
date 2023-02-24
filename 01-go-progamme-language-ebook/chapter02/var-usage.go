package main

import (
	"fmt"
	"os"
)

func main() {
	declaration()
	pointer()

	fmt.Println("\n--- f() func ---")
	fmt.Println("vaule of p:", p)  // call f() func directly
	fmt.Println(f() == f())  // verify different f() func return, false

	fmt.Println("\n--- incr() func ---")
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

// var and short variable list
func declaration() {
	// declare variable by type
	fmt.Println("\n--- declaration() func ---")
	var s string
	fmt.Printf("Print vaule of s: %s", s)
	
	// declare variable list by type or expression
	var i, j, k int  // int, int, int
	i = 7  // assign vaule
	var b, g, p = true, 2.3, "four"  // bool, float64, string
	fmt.Printf("Print vaule of i, j, k, b, g, p: %d, %d, %d, %t, %g, %s\n", i, j, k, b, g, p)
	
	// Note: how to use os.Open() function?
	f, err := os.ReadFile("./osReadFile-func-test.txt")
	if err == nil {
		os.Stdout.Write(f)
	}	
}

// pointer basic
func pointer() {
	fmt.Println("\n--- pointer() func ---")
	s := 2
	p := &s
	// q := &s
	// use &s assign to p and q pointer
	// var s int = 2
	// var p *int = &s
	// declare p pointer to int by &s
	// two different methods to declare variables
	if p != nil {
		// pointer to int can be compared
		fmt.Println("pointer of s: ", &s)
		fmt.Println("vaule of p: ", p)
		fmt.Println("vaule of p point to s: ", *p)
		*p = 3
		// assign vaule to p pointer to change original vaule
		fmt.Println("\ncurrent vaule of s: ", s)
		fmt.Println("current pointer of s is: ", &s)
		fmt.Println("current vaule of p is: ", p)
	}
}

var p = f()
func f() *int {
	v := 1
	return &v
}
// generate different pointer when call f() func every time

func incr(q *int) int {
	*q++
	return *q
}
// return vaule of q point to variable