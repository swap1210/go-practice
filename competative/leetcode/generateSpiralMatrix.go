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

//merge 2 arrays ignoring duplicates
func mergeSortedArray(a, b []int) []int {
	var temp []int
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			temp = append(temp, a[i])
			i++
		} else if a[i] > b[j] {
			temp = append(temp, b[j])
			j++
		} else {
			temp = append(temp, a[i])
			i++
			j++
		}
		k++
	}

	for i < len(a) {
		temp = append(temp, a[i])
		i++
	}

	for j < len(b) {
		temp = append(temp, b[j])
		j++
	}

	return temp
}

//merge 2 arrays preserving duplicate
func mergeSortedArrayPresDup(a, b []int) []int {
	var temp []int
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			temp = append(temp, a[i])
			i++
		} else {
			temp = append(temp, b[j])
			j++
		}
		k++
	}

	for i < len(a) {
		temp = append(temp, a[i])
		i++
	}

	for j < len(b) {
		temp = append(temp, b[j])
		j++
	}

	return temp
}

//end common leet code helper func

func generateMatrix(n int) [][]int {

	if n == 1 {
		return [][]int{{1}}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	ctr := 1
	top, bottom, left, right := 0, n, 0, n
	for top < bottom && left < right {

		//left to right
		// Dln("a", top, bottom, left, right)
		for i := left; i < right; i++ {
			// Dln("a", top, i)
			ans[top][i] = ctr
			ctr++
		}
		top++

		// Dln("b", top, bottom, left, right)
		for i := top; i < bottom; i++ {
			// Dln("b", i, right-1)
			ans[i][right-1] = ctr
			ctr++
		}
		right--

		if !(top < bottom && left < right) {
			break
		}

		// Dln("c", top, bottom, left, right)
		for i := right - 1; i > left-1; i-- {
			//Dln("c", bottom-1, i)
			ans[bottom-1][i] = ctr
			ctr++
		}
		bottom--

		// Dln("d init", top, bottom, left, right)
		for i := bottom - 1; i >= top; i-- {
			// Dln("d", i, left)
			ans[i][left] = ctr
			ctr++
		}
		left++
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
	in := 3
	// fmt.Println(fmt.Sprint(in, out))
	res := generateMatrix(in)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			fmt.Print(res[i][j])
		}
		fmt.Println()
	}

}