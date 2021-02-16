package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

type A struct{}

func main() {
	var buf *bytes.Buffer
	var a *A
	var w io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	fmt.Println(a == nil)
	fmt.Println(buf == nil)
	fmt.Println(w == nil)
	f(buf) // note: subtly wrong
	if debug {
		fmt.Println(buf.String())
	}
}

func f(out io.Writer) {
	// <io.Writer(*bytes.Buffer)>)
	// out: { data:nil <*bytes.Buffer> }
	fmt.Println(out == nil)
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
