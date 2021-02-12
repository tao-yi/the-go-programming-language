package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
	s = append(s, 10)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
	reverse(s)
	fmt.Printf("%v", s)
}

func reverse(input []int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
