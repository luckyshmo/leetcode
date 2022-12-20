package main

import "fmt"

func main() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(t))
}

func dailyTemperatures(t []int) []int {
	ans := make([]int, len(t))
	st := NewStack[int]()
	for i := 0; i < len(t); i++ {
		for !st.IsEmpty() && t[i] > t[st.Peek()] {
			idx := st.Pop()
			ans[idx] = i - idx
		}
		st.Push(i)
	}
	return ans
}
