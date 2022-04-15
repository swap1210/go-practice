package main

import (
	"fmt"
	"math"
	"strconv"
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

// function, which takes a string as
// argument and return the reverse of string.
func reverseStr(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

//end common leet code helper func

func reverse(input int) int {
	str := fmt.Sprint(input)
	negative := str[0] == '-'
	if negative {
		str = str[1:]
	}
	ans, _ := strconv.Atoi(reverseStr(str))
	if negative {
		ans *= -1
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return ans
}

//copy till here
//ctrl+A hold shift

func main() {
	//comment this flag set before submissionor ignore for function base exams
	g_df = true
	// g_df = false //default
	Dln("DEBUG MODE")
	//end debug ops
	fmt.Println(reverse(-123))
}
