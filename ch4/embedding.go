package main

import "fmt"

func main() {
	var w Wheel
	w.x = 8 // equivalent to w.Circle.Point.x = 8
	fmt.Printf("%#v\n", w)
	w.Circle.Point.x = 10
	fmt.Printf("%#v\n", w)

}

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
