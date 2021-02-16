package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")()

	// lots of work
	time.Sleep(2 * time.Second)
	fmt.Println("exiting...")
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
