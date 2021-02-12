package main

import "fmt"

func main() {
	var two int = 2

	a := two << 1 // 4
	b := two << 2 // 8
	c := two << 3 // 16
	d := two >> 1 // 1
	e := two >> 2 // 0
	f := two >> 3 // 0
	fmt.Println(a, b, c, d, e, f)
}
