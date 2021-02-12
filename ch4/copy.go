package main

import "fmt"

func main() {
	str1 := "str1"
	str2 := "str2"
	a := [...]*string{&str1, &str2}
	fmt.Printf("%v, %v\n", *a[0], *a[1])

	inc(a)
	fmt.Printf("%v, %v\n", *a[0], *a[1])

	a1 := []int{1, 2, 3, 4, 5}
	var a2 []int
	copy(a2, a1)

	fmt.Printf("%v, %d, %d\n", a2, a2, a1)
}

func inc(arr [2]*string) {
	for _, i := range arr {
		*i += "hello"
	}
}
