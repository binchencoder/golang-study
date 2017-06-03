package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("1", "a")
	addEdge("1", "b")
	addEdge("2", "c")
	addEdge("2", "d")
	addEdge("3", "e")

	from := "1"
	to := "c"
	fmt.Printf("hasEdge from: %s, to: %s \t %t\n", from, to, hasEdge(from, to))
}
