package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}
	fmt.Printf("%#v\n", movies)

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data1, err1 := json.MarshalIndent(movies, "", "  ")
	if err1 != nil {
		log.Fatal("JSON marshaling failed: %s", err1)
	}
	fmt.Printf("%s\n", data1)
}
