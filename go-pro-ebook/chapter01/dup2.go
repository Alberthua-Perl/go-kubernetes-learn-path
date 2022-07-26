// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
// NOTE: 
//   Dup2 includes the usage of short variable declaration, slice, map,
//   function, struct type and method.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// make builtin func create empty map[]
	files := os.Args[1:]
	// files declared by os.Args for slice
	// fmt.Println("counts map and files slice:", counts, files)
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// loop return index and file names and ignore index through "_"
			f, err := os.Open(arg)
			// f pointer called *File point to type File struct
			fmt.Println(f.Name())
			// call f.Name() method defined by type File struct to output filename
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// pass f pointer and counts parameters and return
			// counts map including all lines and count nubber
			f.Close()	// call f.Close() method to close file descriptor
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)  // output all lines from files
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	// two parameters: f type *os.file and counts type map[]
	input := bufio.NewScanner(f)
	// previous func return *Scanner which point to type Scanner struct
	for input.Scan() { 
		counts[input.Text()]++
		// input use Scan and Text method of type Scanner struct
	}
	// NOTE: ignoring potential errors from input.Err()
}
