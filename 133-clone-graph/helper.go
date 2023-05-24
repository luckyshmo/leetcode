package main

import (
	"fmt"
	"strings"
)

func buildGraph(adjList [][]int) *Node {
	nodeMap := make(map[int]*Node)

	// Create nodes and populate nodeMap
	for _, neighbors := range adjList {
		val := neighbors[0]
		if _, ok := nodeMap[val]; !ok {
			nodeMap[val] = &Node{Val: val}
		}

		for _, neighbor := range neighbors[1:] {
			if _, ok := nodeMap[neighbor]; !ok {
				nodeMap[neighbor] = &Node{Val: neighbor}
			}
		}
	}

	// Connect neighbors
	for _, neighbors := range adjList {
		val := neighbors[0]
		node := nodeMap[val]

		for _, neighbor := range neighbors[1:] {
			neighborNode := nodeMap[neighbor]
			node.Neighbors = append(node.Neighbors, neighborNode)
		}
	}

	return nodeMap[adjList[0][0]]
}

func printGraph(node *Node) {
	visited := make(map[*Node]bool)
	printGraphHelper(node, visited, 0)
}

func printGraphHelper(node *Node, visited map[*Node]bool, level int) {
	if visited[node] {
		return
	}

	visited[node] = true

	fmt.Println(strings.Repeat("  ", level), node.Val)

	for _, neighbor := range node.Neighbors {
		fmt.Println(strings.Repeat("  ", level+1), "|")
		fmt.Println(strings.Repeat("  ", level+1), "└─", neighbor.Val)
		printGraphHelper(neighbor, visited, level+1)
	}
}
