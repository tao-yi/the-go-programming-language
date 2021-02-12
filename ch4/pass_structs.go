package main

import "fmt"

func main() {
	w := Wheel{size: 1}
	v := Vehicle{wheel: &w, speed: 15}
	fmt.Printf("%v", v)

	ChangeVehicle(v)
	fmt.Printf("%v", v)
}

func ChangeVehicle(v Vehicle) {
	v.wheel.size = 10
	v.speed = 25
}

type Vehicle struct {
	wheel *Wheel
	speed int
}

type Wheel struct {
	size int
}
