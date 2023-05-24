package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
	head := buildGraph([][]int{{1, 2}, {2, 1}, {1, 3}, {3, 1}, {2, 3}, {3, 2}})
	nh := cloneGraph(head)
	fmt.Println(nh.Val)
}

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited = map[*Node]*Node{}
	return rec(node)
}

func rec(curr *Node) *Node {
	neighbors := make([]*Node, 0)
	clone := &Node{Val: curr.Val}
	visited[curr] = clone
	for _, n := range curr.Neighbors {
		if v, ok := visited[n]; ok {
			neighbors = append(neighbors, v)
		} else {
			neighbors = append(neighbors, rec(n))
		}
	}
	clone.Neighbors = neighbors
	return clone
}
