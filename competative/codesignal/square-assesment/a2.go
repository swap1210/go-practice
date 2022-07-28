package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]string{
		{".", ".", ".", "."},
		{".", "F", ".", "."},
		{".", "F", "F", "F"},
		{".", ".", ".", "."},
		{".", ".", "#", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	fmt.Println("before")
	p2d(matrix)
	fmt.Println("after")
	p2d(freefall(matrix))
}

func freefall(m [][]string) [][]string {
	ans := make([][]string, len(m))
	min := math.MaxInt64
	for r_i, r := range m {
		ans[r_i] = make([]string, len(m[r_i]))
		for c_i, c := range r {
			ans[r_i][c_i] = m[r_i][c_i]
			//fmt.Println("copying", c, " at ", r_i, c_i, ans[r_i][c_i])
			if c == string('F') {
				t := dFromFloor(m, r_i, c_i)
				if t != -1 && t < min {
					//fmt.Println(t, " for ", r_i, c_i)
					min = t
				}
				if min == 0 {
					break
				}
			}
		}
	}

	if min == 0 {
		return m
	} else {

		for r_i, r := range m {
			for c_i, c := range r {
				if c == string('F') {
					if r_i-min >= 0 {
						ans[r_i][c_i] = m[r_i-1][c_i]
					} else {
						ans[r_i][c_i] = string('.')
					}
					ans[r_i+min][c_i] = string('F')
				}
			}
		}
	}
	// fmt.Println("final min dis is", min)
	return ans
}

func dFromFloor(m [][]string, p_i, p_j int) int {
	dist := 0
	for i := p_i + 1; i < len(m); i++ {
		//fmt.Println("seaching ", i, p_j)
		if m[i][p_j] == string('#') {
			return dist
		} else if m[i][p_j] == string('F') {
			// fmt.Println("breking at")
			return -1
		}
		dist++
	}
	return dist
}

func p2d(m [][]string) {
	for _, v := range m {
		for _, b := range v {
			fmt.Print(b, " ")
		}
		fmt.Println()
	}
}
