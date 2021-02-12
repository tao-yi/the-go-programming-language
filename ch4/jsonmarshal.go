package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Actors: []string{"a", "b"}},
		{Title: "Casablanca1", Year: 1943, Color: false, Actors: []string{"c", "d"}},
		{Title: "Casablanca2", Year: 1944, Color: false, Actors: []string{"e", "f"}},
	}

	data, _ := json.Marshal(movies)
	fmt.Printf("%s\n", data)

	data, _ = json.MarshalIndent(movies, "", "  ")
	fmt.Printf("%s\n", data)
}
