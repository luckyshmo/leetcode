package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	ans := make([]string, numRows)
	for i, row := 0, 0; i < len(s); {
		ans[row] = ans[row] + string(s[i])
		row++
		i++
		if row < numRows {
			continue
		}
		// fill diagonal
		for row = numRows - 2; i < len(s) && row > 0; i, row = i+1, row-1 {
			ans[row] = ans[row] + string(s[i])
		}
		row = 0
	}
	Ans := ""
	for _, str := range ans {
		Ans += str
	}
	return Ans
}
