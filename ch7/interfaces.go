package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer
	var rwc io.ReadWriteCloser

	rwc = os.Stdout
	rwc.Close()
	rwc.Read()
	rwc.Write()

	w = rwc
	w.Write()
}
