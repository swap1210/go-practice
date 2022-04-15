package main

import (
	"container/heap"
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

func D(a interface{}) {
	if g_df {
		fmt.Print(a)
	}
}

func Dln(a interface{}) {
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

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}

	for _, val := range nums {
		// D("b")
		// Dln(*h, val, h.Peek(), h.Len() == k, val > h.Peek())
		if h.Len() == k && val > h.Peek() {
			heap.Pop(h)
			heap.Push(h, val)
		} else if h.Len() < k {
			heap.Push(h, val)
		}
		// D("a")
		// Dln(*h)
	}
	// Dln(h)
	ans := heap.Pop(h).(int)
	return ans
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	kthLargest := findKthLargest(mapValsInts(m), k)
	ans_arr := []int{}
	for k, f := range m {
		if f >= kthLargest {
			ans_arr = append(ans_arr, k)
		}
	}
	return ans_arr
}

//copy till here
//ctrl+A hold shift

func main() {
	//comment this flag set before submissionor ignore for function base exams
	g_df = true
	// g_df = false //default
	Dln("DEBUG MODE\n")
	//end debug ops
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
