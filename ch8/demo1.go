package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	start := time.Now()
	concurrency := 10
	res := make(chan string)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < concurrency; i++ {
		go SendRequest(i, res)
	}

	for i := 0; i < concurrency; i++ {
		go ProcessResponse(res)
	}

	log.Println("All Done!")
	logTime(start)
	SigHandler(c)
}

func logTime(start time.Time) {
	duration := time.Since(start)
	log.Println(duration)
}

// could be any thing
func SendRequest(taskID int, res chan<- string) {
	// time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	time.Sleep(2 * time.Second)
	res <- fmt.Sprintf("task: %d Done!", taskID)
}

func ProcessResponse(res <-chan string) {
	time.Sleep(3 * time.Second)
	log.Println(<-res)
}

func SigHandler(c <-chan os.Signal) {
	sig := <-c
	log.Println("receive signal", sig)
}
