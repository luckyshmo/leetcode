package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func twoSum(numbers []int, target int) []int {
	begin, end := 0, len(numbers)-1
	for {
		sum := numbers[begin] + numbers[end]
		if sum < target {
			begin++
		} else if sum > target {
			end--
		} else {
			return []int{begin + 1, end + 1}
		}
	}
}
