package main

import "fmt"

func main() {
	arr := []int{1, 1, 2}
	kek := removeDuplicates(arr)
	fmt.Println("arr: ", arr, "; kek: ", kek)
}

func removeDuplicates(n []int) int {
	i, j := 0, 1
	for ; i < len(n)-1 && j < len(n); j++ {
		if n[i] != n[j] {
			n[i+1] = n[j]
			i++
		}
	}
	return i + 1
}
