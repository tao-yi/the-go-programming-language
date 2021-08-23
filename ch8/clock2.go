package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	portPtr := flag.Int("port", 8000, "port number")
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *portPtr)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		timezone := os.Getenv("TZ")
		loc, err := time.LoadLocation(timezone)
		now := time.Now()
		if err == nil {
			now = now.In(loc)
		}
		_, err = io.WriteString(c, now.Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
