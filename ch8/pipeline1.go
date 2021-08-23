package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			time.Sleep(1 * time.Second)
			naturals <- x
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for {
		fmt.Println(<-squares)
	}
}
