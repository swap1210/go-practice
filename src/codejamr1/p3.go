package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T) //T = 1 //
	for t := 1; t <= T; t++ {
		fmt.Printf("Case #%d: %d\n", t, sol())
	}
}

func sol() int {
	val := 0
	E, W := 0, 0
	fmt.Scan(&E) //E = 3 //
	fmt.Scan(&W) //W = 3 //
	// m := [][]int{{3, 1, 1}, {3, 3, 3}, {2, 3, 3}}
	sum := make([]int, E)
	min := make([]int, W)
	m := make([][]int, E)
	for r := 0; r < E; r++ {
		m[r] = make([]int, W)
		for c := 0; c < W; c++ {
			fmt.Scan(&m[r][c])
			sum[r] += m[r][c]
			// fmt.Print(r, c, ", ")
			if r == 0 || min[c] > m[r][c] {
				min[c] = m[r][c]
			}
		}
		// fmt.Println()
	}

	fmt.Println(sum)
	fmt.Println(min)
	fmt.Println(m)
	for _, v := range sum {
		val += v
	}
	return val
}
