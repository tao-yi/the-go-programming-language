package main

import "fmt"

func main() {
	x := 1
	addr := &x

	var nullptr *int
	fmt.Printf("%v\n", addr)    // the address of variable x: 0xc0000140a0
	fmt.Printf("%v\n", *addr)   // the value addr points to: 1
	fmt.Printf("%v\n", nullptr) // the zero value for a pointer of any type is <nil>

	var nullptr1 *int
	fmt.Println(nullptr1 == nullptr)

	var intPtr = new(int)
	var strPtr = new(string)
	fmt.Printf("%v->value(%v), %v->value(%v)\n", intPtr, *intPtr, strPtr, *strPtr)
}
