package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g`C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

type MyType struct {
	X int
}

func main() {
	var x = 1
	var num = 'A'
	fmt.Printf("%s, %T\n", string(x), string(x))
	fmt.Printf("%v, %T\n", int(num), int(num))
	// fmt.Printf("%v, %T\n", MyType(num), MyType(num))

	fmt.Println(AbsoluteZeroC)
}
