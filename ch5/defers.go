package main

import "fmt"

func main() {
	defer say("one")
	defer say("two")
	defer say("three")
	fmt.Println("main end")
}

func say(word string) {
	fmt.Printf("%s\n", word)
}
