package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) (profit int) {
	lowest := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] > lowest {
			profit += prices[i] - lowest
		}
		lowest = prices[i]
	}
	return
}
