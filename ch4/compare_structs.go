package main

import "fmt"

func main() {
	p := Point{x: 1, y: 2}
	q := Point{x: 1, y: 2}
	fmt.Println(p == q)

	// c1 := Complex{}
	// c2 := Complex{}
	// fmt.Println(c1 == c2) // compile error, (operator == not defined for Complex)
}

type Point struct {
	x, y int
}

type Complex struct {
	num []int // array is not comparable
	x   int
	y   int
}
