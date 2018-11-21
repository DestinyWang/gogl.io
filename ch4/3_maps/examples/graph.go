package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {     // 类似 Map<String, Map<String, Boolean>>
	edgs := graph[from]             // edgs 是 Map<String, Boolean>
	if edgs == nil {
		edgs = make(map[string]bool)
		graph[from] = edgs
	}
	edgs[to] = true
}

func hasEdgs(from, to string) bool {
	return graph[from][to]
}

func main() {
	fmt.Print("\nlen\tcount\n")
}
