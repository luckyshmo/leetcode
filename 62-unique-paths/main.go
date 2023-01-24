package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
	b := time.Now()
	fmt.Println(uniquePaths(51, 9))
	fmt.Println(time.Since(b))

}

func uniquePaths(m int, n int) (paths int) {
	if m == 1 || n == 1 {
		return 1
	}

	mm := make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		mm[i][0] = 1
	}

	for i := 0; i < n; i++ {
		mm[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mm[i][j] = mm[i-1][j] + mm[i][j-1]
		}
	}

	return mm[m-1][n-1]
}

func uniquePathsReq(m int, n int) (paths int) {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(n-1, m)
}
