package main

import (
	"fmt"
	"strconv"
)

func main() {
	// yields a new string consisting of the bytes of the original string starting at index i
	s := "Hello, world!"
	s1 := s[1:5]
	fmt.Println(s1)

	// create a copy of string s
	s2 := s[:]
	fmt.Println(s == s2)
	fmt.Println(&s == &s2)

	str := "0123456"
	str1 := str[:5]
	fmt.Printf("%v, %v\n", &str, &str1)

	numStr := 12341
	num := strconv.Itoa(numStr)
	strN, _ := strconv.Atoi(num)
	fmt.Printf("%s, %d", num, strN)
}
