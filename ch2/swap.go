package main

import "fmt"

func main() {
	var x, y = 1, 2
	x, y = y, x
	fmt.Println(x, y)

	arr := [2]int{1, 2}
	fmt.Printf("%v\n", arr)
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Printf("%v\n", arr)
}
