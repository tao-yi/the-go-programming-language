package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	var err error
	err = errors.New("hello world!")
	fmt.Println("%v", err)
	err = syscall.Errno(1)
	fmt.Println("%v", err)
}
