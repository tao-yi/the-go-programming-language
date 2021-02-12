package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := map[string]int{}
	m2 := make(map[string]int)

	fmt.Printf("%v, %v\n", m1, m2)
	fmt.Printf("%v, %v\n", len(m1), len(m2))

	m1["alice"] = 32
	fmt.Printf("%v\n", m1)
	delete(m1, "alice") // remove element m1["alice"]
	fmt.Printf("%v\n", m1)

	m1["aaa"] = 1
	m1["abc"] = 2
	m1["bbb"] = 3
	m1["bcd"] = 4
	for k, v := range m1 {
		fmt.Printf("key: %s, val: %d\n", k, v)
	}

	var names []string
	for name := range m1 {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, m1[name])
	}

	v, ok := m1["bob"]
	if !ok {
		fmt.Println("bob is not a key in this map, age === 0")
	}
	fmt.Println("value", v)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
