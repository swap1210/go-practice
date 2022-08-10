package main

import (
	"container/heap"
	"fmt"
)

func main() {
	a := []int{3, 2, 1, 4}
	fmt.Println(thirdMax(a))
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func thirdMax(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		m[v] = false
	}
	l := &MaxHeap{}
	for v := range m {
		heap.Push(l, v)
	}

	if len(m) < 3 {
		return heap.Pop(l).(int)
	}

	for i := 0; i < 2; i++ {
		heap.Pop(l)
	}
	return heap.Pop(l).(int)
}
