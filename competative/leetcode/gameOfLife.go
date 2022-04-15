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

//copy 2d array by value
func copy2D(original, copy [][]int) {
	R, C := len(original), len(original[0])
	for x := 0; x < R; x++ {
		copy[x] = make([]int, C)
		for y := 0; y < C; y++ {
			copy[x][y] = original[x][y]
		}
	}
}

//end common leet code helper func

func gameOfLife(board [][]int) {
	R, C := len(board), len(board[0])
	// Dln("before", board)
	temp := make([][]int, len(board))
	for i := 0; i < R; i++ {
		temp[i] = make([]int, len(board[0]))
		for j := 0; j < C; j++ {
			neighbour := 0
			for n_i := i - 1; n_i < (i + 2); n_i++ {
				for n_j := j - 1; n_j < (j + 2); n_j++ {
					if n_i < 0 || n_j < 0 || n_i == R || n_j == C || (n_i == i && n_j == j) {
						continue
					} else {
						// Dln(i, j, "neighbour", n_i, n_j)
						neighbour += board[n_i][n_j]
					}
				}
			}

			// Dln(i, j, "=>", neighbour)
			if neighbour < 2 || neighbour > 3 {
				temp[i][j] = 0
			} else if neighbour == 3 {
				temp[i][j] = 1
			} else {
				temp[i][j] = board[i][j]
			}
		}
	}
	copy2D(temp, board)
}

//copy till here
//ctrl+A hold shift

func main() {
	//comment this flag set before submissionor ignore for function base exams
	g_df = true
	// g_df = false //default
	Dln("DEBUG MODE")
	//end debug ops
	input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(input)
	// [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
	fmt.Println(input)
}
