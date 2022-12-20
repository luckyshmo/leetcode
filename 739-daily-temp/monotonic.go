package main

// One famous monotonic stack interview question is the sunset view problem.
// In this problem, you’re given an array of buildings,
// where each building is represented by its height.
// The buildings are in line with each other,
// meaning that if one building is taller than the other,
// they will be blocked from seeing the sunset.
// The sun sets to the right side.
// Your goal is to find all the buildings that can see the sunset.
func MonotonicExample() {
	buildings := []int{18, 14, 13, 16, 12}

	st := NewStack[int]()
	for i := 0; i < len(buildings); i++ {
		currentHeight := buildings[i]
		for !st.IsEmpty() && currentHeight >= buildings[st.Peek()] {
			st.Pop()
		}
		st.Push(i)
	}
	st.Print()
}
