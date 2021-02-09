package main

import (
	"fmt"
	"os"
	"strings"
)

// morify the echo program to also print os.Args[0]
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
