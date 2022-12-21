package main

func main() {
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms1 := [][]int{{1}, {2}, {3}, {}}
	println(canVisitAllRooms(rooms))
	println("\nSecond case:\n")
	println(canVisitAllRooms(rooms1))

}

func canVisitAllRooms(rooms [][]int) bool {
	ans := make([]bool, len(rooms))
	ans[0] = true
	openRoom(0, rooms, ans)
	for _, r := range ans {
		if !r {
			return false
		}
	}
	return true
}

func openRoom(key int, rooms [][]int, ans []bool) {
	for _, key := range rooms[key] {
		if ans[key] {
			continue
		}
		ans[key] = true
		openRoom(key, rooms, ans)
	}
}
