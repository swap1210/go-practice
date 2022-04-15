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
func (h MaxHeap) Peek(i int) int {
	if h.Len() > 0 {
		return h[i]
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

// maxheap end

// This example inserts several ints into an MaxHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	nums := []int{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}
	maxSize := 3
	// initialize the heap data structure
	h := &MaxHeap{}

	// add all the values to heap, O(n log n)
	fmt.Println("maxSize", "h.Len()", "val", "root")
	for _, val := range nums { // O(n)
		root := h.Peek(0)
		if maxSize == h.Len() && val < root {
			heap.Pop(h)
			heap.Push(h, val) // O(log n)
		} else if h.Len() < maxSize {
			heap.Push(h, val) // O(log n)
		}
	}

	// print all the values from the heap
	// which should be in ascending order
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Printf("%d,", heap.Pop(h).(int))
	// }

	fmt.Println(h)
}
