// This example demonstrates an integer heap built using the heap interface.
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

//end common leet code helper func
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}

	for _, val := range nums {
		D("b")
		Dln(*h, val, h.Peek(), h.Len() == k, val > h.Peek())
		if h.Len() == k && val > h.Peek() {
			heap.Pop(h)
			heap.Push(h, val)
		} else if h.Len() < k {
			heap.Push(h, val)
		}
		D("a")
		Dln(*h)
	}
	Dln(h)
	ans := heap.Pop(h).(int)
	return ans
}

//copy till here
//ctrl+A hold shift click here

// This example inserts several ints into an MinHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	g_df = true
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Print(findKthLargest(nums, 2))
	// // initialize the heap data structure
	// h := &MinHeap{}

	// // add all the values to heap, O(n log n)
	// for _, val := range nums { // O(n)
	// 	heap.Push(h, val) // O(log n)
	// }

	// // print all the values from the heap
	// // which should be in ascending order
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Printf("%d,", heap.Pop(h).(int))
	// }
}
