package main

import (
	"fmt"
	"math"
)

func main() {
	input := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(input))
}

var g_seen [][]bool
var g_grid [][]int

func twoDP(a [][]bool) {
	for _, r := range a {
		for _, c := range r {
			if c {
				fmt.Print("✅")
			} else {
				fmt.Print("❌")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
func area(r, c int) int {
	// fmt.Println(r, " ", c)
	// twoDP(g_seen)
	if r < 0 || r >= len(g_grid) || c < 0 || c >= len(g_grid[0]) || g_seen[r][c] || g_grid[r][c] == 0 {
		return 0
	}

	g_seen[r][c] = true
	return 1 + area(r+1, c) + area(r-1, c) + area(r, c-1) + area(r, c+1)
}

func maxAreaOfIsland(grid [][]int) int {
	g_grid = grid
	g_seen = make([][]bool, len(g_grid))

	for i := range g_grid {
		g_seen[i] = make([]bool, len(g_grid[i]))
	}
	ans := math.MinInt
	for r := 0; r < len(g_grid); r++ {
		for c := 0; c < len(g_grid[r]); c++ {
			new_area := area(r, c)
			if new_area > ans {
				ans = new_area
			}
		}
	}

	return ans
}
