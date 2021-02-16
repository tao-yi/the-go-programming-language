package main

import (
	"fmt"
)

type A struct {
	X, Y float64
}

func (a *A) Print() {
	fmt.Printf("%#v", *a)
}

type B struct {
	M, N float64
}

func (b *B) Print() {
	fmt.Printf("%#v", *b)
}

type AB struct {
	A
	B
}

func main() {
	var a AB
	var b AB
	a.X = 1
	a.Y = 2
	b.M = 3
	b.N = 4
	a.B.Print()
	a.A.Print()
}
