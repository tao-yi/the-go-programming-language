package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)      // success, f == os.Stdout
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Println(f, b, ok)
}
