package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	s := mirroredQuery()
	log.Println(s)
}

func request(hostname string) (response string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return fmt.Sprintf("response for: %s", hostname)
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("asia.gopl.io")
	}()
	go func() {
		responses <- request("europe.gopl.io")
	}()
	go func() {
		responses <- request("americas.gopl.io")
	}()

	return <-responses
}
