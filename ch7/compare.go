package main

import "fmt"

type Comparable interface {
	Compare() bool
}

type A struct{}
type B struct{}

func (a A) Compare() bool {
	return true
}

func (a B) Compare() bool {
	return true
}

func main() {
	var a Comparable
	var b Comparable
	a = A{}
	b = B{}
	// interface values are comparable
	fmt.Println(a == b) // false
}
