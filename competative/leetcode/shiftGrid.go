package main

import (
	"fmt"
	"math"
)

//max heap start
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int {
	if h.Len() > 0 {
		return h[0]
	} else {
		return math.MinInt
	}
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int {
	if h.Len() > 0 {
		return h[0]
	} else {
		return math.MinInt
	}
}

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//some common leetcode helper func
//global debug flag
var g_df bool

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//sum of all integers
func sumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func D(a ...interface{}) {
	if g_df {
		fmt.Print(a)
	}
}

func Dln(a ...interface{}) {
	if g_df {
		fmt.Println(a)
	}
}

func mapKeysInts(mymap map[int]int) []int {
	keys := make([]int, len(mymap))

	i := 0
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func mapValsInts(mymap map[int]int) []int {
	keys := make([]int, len(mymap))

	i := 0
	for _, v := range mymap {
		keys[i] = v
		i++
	}
	return keys
}

//end common leet code helper func

//name variable correctly
func shiftGrid(grid [][]int, k int) [][]int {
	R, C := len(grid), len(grid[0])
	n := make([][]int, R)
	for i, _ := range n {
		n[i] = make([]int, C)
	}
	for i := 1; i <= k; i++ {
		r, c := 0, 0

		if C > 1 {
			r, c = 0, 1
		} else if R > 1 {
			r, c = 1, 0
		}
		for r+c != 0 {
			new_r, new_c := resolveIndex(r, c, R, C)
			// Dln(r, c, new_r, new_c)
			n[new_r][new_c] = grid[r][c]
			r, c = resolveIndex(r, c, R, C)
		}

		if C > 1 {
			n[0][1] = grid[0][0]
		} else if R > 1 {
			n[1][0] = grid[0][0]
		}
		Dln(n)
		//copy n to grid
		//grid = n
		for x := 0; x < R; x++ {
			for y := 0; y < C; y++ {
				grid[x][y] = n[x][y]
			}
		}
	}
	return n
}

func resolveIndex(i, j, R, C int) (int, int) {
	//last to one
	if i == (R-1) && j == (C-1) {
		return 0, 0
	} else {
		if (j + 1) <= (C - 1) {
			j++
		} else if (j + 1) == C {
			i++
			j = 0
		}
	}
	return i, j
}

//copy till here
//ctrl+A hold shift

func main() {
	//comment this flag set before submissionor ignore for function base exams
	// g_df = true
	g_df = false //default
	Dln("DEBUG MODE\n")
	//end debug ops
	// input := [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}
	input := [][]int{{3}, {19}, {4}, {12}}
	fmt.Println(shiftGrid(input, 2))
	Dln(resolveIndex(0, 0, 4, 1))
}
