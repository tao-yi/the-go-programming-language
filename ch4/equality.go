package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [...]int{2, 1}
	fmt.Println(a == b)
	fmt.Println(a == c)
}
