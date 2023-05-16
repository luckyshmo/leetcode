package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(50))
}

var mem []int

func climbStairs(n int) int {
	mem = make([]int, n+1)
	return rec(n)
}

func rec(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	if mem[n] != 0 {
		return mem[n]
	}
	mem[n] = rec(n-1) + rec(n-2)
	return mem[n]
}
