package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

type elem struct {
	value      int
	currentMin int
}

type MinStack struct {
	arr []elem
}

func Constructor() MinStack {
	return MinStack{make([]elem, 0)}
}

func (this *MinStack) Push(val int) {
	var min int
	if len(this.arr) < 1 {
		min = val
	} else if this.arr[len(this.arr)-1].currentMin < val {
		min = this.arr[len(this.arr)-1].currentMin
	} else {
		min = val
	}
	this.arr = append(this.arr, elem{
		value:      val,
		currentMin: min,
	})
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1].value
}

func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1].currentMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
