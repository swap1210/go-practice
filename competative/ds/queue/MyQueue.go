package main

import "fmt"

type MyQueue struct{ arr []int }

func (queue *MyQueue) enqueue(element int) {
	queue.arr = append(queue.arr, element)
}

func (queue *MyQueue) dequeue() int {
	if queue.isEmpty() {
		return -1
	}
	element := queue.arr[0]
	queue.arr = queue.arr[1:]
	return element
}

func (queue *MyQueue) isEmpty() bool {
	return len(queue.arr) == 0
}

func (queue *MyQueue) peek() int {
	if queue.isEmpty() {
		return -1
	}
	l := len(queue.arr)
	return queue.arr[l-1]
}

func (o *MyQueue) reverse() {
	st := &MyStack{}
	for !o.isEmpty() {
		st.push(o.dequeue())
	}

	for !st.isEmpty() {
		o.enqueue(st.pop())
	}
}

type MyStack struct {
	arr []int
}

func (o *MyStack) push(a int) {
	o.arr = append([]int{a}, o.arr...)
}

func (o *MyStack) pop() int {
	if o.isEmpty() {
		return -1
	} else {
		t := o.arr[0]
		o.arr = o.arr[1:]
		return t
	}
}

func (o *MyStack) isEmpty() bool {
	return len(o.arr) == 0
}

func main() {
	queue := &MyQueue{}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)

	// fmt.Println("Queue:", queue)

	// queue.dequeue()
	// fmt.Println("Queue:", queue)

	queue.enqueue(40)
	fmt.Println("Queue:", queue)
	// fmt.Println("Queue:", queue.peek())
	// queue.dequeue()
	fmt.Println("Queue:", queue, queue.peek())
	queue.reverse()
	fmt.Println(queue, queue.peek())
}
