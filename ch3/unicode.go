package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "123vfa`1545`1XZVP; [ 215;1"

	for _, r := range str {
		fmt.Printf("%s, IsDigit %v\n", string(r), unicode.IsDigit(r))
		fmt.Printf("%s, IsLzetter %v\n", string(r), unicode.IsLetter(r))
		fmt.Printf("%s, IsLozwer %v\n", string(r), unicode.IsLower(r))
		fmt.Printf("%s, IsNuzmber %v\n", string(r), unicode.IsNumber(r))
		fmt.Printf("%s, IsSpzace %v\n", string(r), unicode.IsSpace(r))
		fmt.Printf("%s, IsSyzmbol %v\n", string(r), unicode.IsSymbol(r))
	}
}
