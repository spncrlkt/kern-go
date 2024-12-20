package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// go run dup3.go a b
func main() {
	counts := make(map[string]int)
	for _, fname := range os.Args[1:] {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			line = fname + " " + line
			counts[line]++
		}
	}

	fmt.Printf("#\tfile\tline\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
