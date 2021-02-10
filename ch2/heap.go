package main

import "fmt"

func main() {
	f()
	fmt.Println(*global)
}

var global *int

func f() {
	var x int
	x = 1
	global = &x
}
