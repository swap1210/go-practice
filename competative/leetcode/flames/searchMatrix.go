package main

import (
	"fmt"
)

var ctr = 6

func main() {
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 12, 91}}
	// matrix := [][]int{{-1, 3}}
	matrix := [][]int{{1, 4}, {}}
	p2d(matrix)
	fmt.Println(searchMatrix(matrix, 3))
	// fmt.Println(indexTranslate(matrix, 2))
}

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)*len(matrix[0])-1
	return search(matrix, target, l, r)
}

func search(matrix [][]int, target, l, r int) bool {
	if ctr < 0 {
		return false
	}
	if l > r {
		return false
	}
	row, col := indexTranslate(matrix, ((r-l)/2)+l)
	fmt.Println("meadian is ", row, col, l, r)
	diff := target - matrix[row][col]

	if diff == 0 {
		return true
	} else {
		if diff < 0 {
			// ctr--
			return search(matrix, target, l, ((r-l)/2)-1+l)
		} else {
			// ctr--
			return search(matrix, target, ((r-l)/2)+1+l, r)
		}
	}
}

func indexTranslate(m [][]int, index int) (int, int) {
	// return int(math.Floor(float64(index / len(m[0])))), (index % len(m))
	return index / len(m[0]), (index % len(m[0]))
}

func p2d(m [][]int) {
	for _, v := range m {
		for _, b := range v {
			fmt.Print(b, " ")
		}
		fmt.Println()
	}
}
