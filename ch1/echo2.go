package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		// 每次 += 都会创建一个新的字符串
		// 然后赋值给 s，旧的s就不再被使用，所以会被垃圾回收
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
