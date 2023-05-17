package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	l := len(prices)
	buy, sell, rest := make([]int, l), make([]int, l), make([]int, l)
	buy[0] = -prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		buy[i] = max(rest[i-1]-price, buy[i-1])
		sell[i] = max(buy[i-1]+price, sell[i-1])
		rest[i] = max(max(sell[i-1], buy[i-1]), rest[i-1])
	}
	return max(max(sell[l-1], buy[l-1]), rest[l-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
