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

//end common leet code helper func

type MinStack struct {
	arr     []int
	minList *MinHeap
}

func Constructor() MinStack {
	return MinStack{[]int{}, &MinHeap{}}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	heap.Push(this.minList, val)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
	tempMinList := &MinHeap{}
	for _, num := range this.arr {
		heap.Push(tempMinList, num)
	}
	this.minList = tempMinList
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minList.Peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//copy till here
//ctrl+A hold shift

func main() {
	//comment this flag set before submissionor ignore for function base exams
	g_df = true
	// g_df = false //default
	Dln("DEBUG MODE\n")
	//end debug ops
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(14)
	obj.Push(12)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3, param_4)
}
