package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(10, []int{0, 4, 2}, []int{2, 1, 3}))
	fmt.Println(carFleet(10, []int{8, 3, 7, 4, 6, 5}, []int{4, 4, 4, 4, 4, 4}))
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{6, 8}, []int{3, 2}))
}

func carFleet(target int, position []int, speed []int) int {
	cars := Cars{
		positions: position,
		speed:     speed,
	}
	sort.Sort(&cars)

	t := time(speed[0], position[0], target)
	var st Stack = []float64{t}
	for i := 1; i < len(position); i++ {
		t := time(speed[i], position[i], target)
		st.Push(t)
	}

	return len(st)
}

func time(speed, position, target int) float64 {
	return float64(target-position) / float64(speed)
}

type Stack []float64

func (s *Stack) Push(time float64) {
	k := *s
	if k[len(k)-1] < time {
		*s = append(*s, time)
	}
}

type Cars struct {
	positions []int
	speed     []int
}

func (c *Cars) Len() int { return len(c.positions) }

func (c *Cars) Less(i, j int) bool { return c.positions[i] > c.positions[j] }

func (c *Cars) Swap(i, j int) {
	c.positions[i], c.positions[j] = c.positions[j], c.positions[i]
	c.speed[i], c.speed[j] = c.speed[j], c.speed[i]
}
