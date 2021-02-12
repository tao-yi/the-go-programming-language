package main

import "fmt"

func main() {
	str := "hello!"
	fmt.Printf("%s, %[1]s, %[1]s\n", str) // hello!, hello!, hello!

	// unsigned integer sufficient to hold all the bits of a pointer value
	var ptr uintptr
	fmt.Println(ptr)
}
