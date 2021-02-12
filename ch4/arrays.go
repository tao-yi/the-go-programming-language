package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	for _, i := range a {
		fmt.Println(i)
	}

	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%T, %T\n", arr1, arr2)

	symbol := [...]string{USD: "$", RMB: "$"}
	fmt.Printf("%v", symbol)
}
