package main

import "fmt"

func main() {
	ages := make(map[string]int) // create a map through make func, mapping from strings to ints
	fmt.Printf("create a map through make:%v\n", ages)
}
