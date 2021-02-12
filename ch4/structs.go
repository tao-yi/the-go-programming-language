package main

import "time"

func main() {

}

type Employee struct {
	ID            string
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}
