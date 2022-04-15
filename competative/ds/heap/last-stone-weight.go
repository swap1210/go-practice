package main

import (
	"container/heap"
	"fmt"
)

//copy from here
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//needed heap
func lastStoneWeight(s []int) int {
	l := len(s)
	h := &MaxHeap{}
	for i := 0; i < l; i++ {
		heap.Push(h, s[i])
	}

	for h.Len() > 1 {
		// fmt.Println("b", h)
		//remove top2
		diff := Abs(heap.Pop(h).(int) - heap.Pop(h).(int))
		if diff > 0 {
			heap.Push(h, diff)
		}
	}

	if h.Len() == 0 {
		return 0
	} else if h.Len() == 1 {
		return heap.Pop(h).(int)
	} else {
		return Abs(heap.Pop(h).(int) - heap.Pop(h).(int))
	}
}

func main() {
	input := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(input))
}
