package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// the key may be of any type whose values can be compared with `==`
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// each call to input.Scan() reads the next line
	// and removes the newline character from the end
	// returns false when there is no more input
	for input.Scan() {
		counts[input.Text()]++
	}

	// the order of map iteration is not specified
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
