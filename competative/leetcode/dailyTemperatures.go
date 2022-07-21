package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

type MyStack struct {
	arr []int
}

func (o *MyStack) pop() int {
	if o.isEmpty() {
		return -1
	}
	t := o.arr[0]
	o.arr = o.arr[1:]
	return t
}

func (o *MyStack) peek() int {
	if o.isEmpty() {
		return -1
	}
	return o.arr[0]
}

func (o *MyStack) push(i int) {
	o.arr = append([]int{i}, o.arr...)
}

func (o *MyStack) isEmpty() bool {
	return len(o.arr) == 0
}

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	if l < 2 {
		return []int{0}
	}

	s := &MyStack{arr: []int{l - 1}}
	ans := make([]int, l)
	for i := l - 2; i > -1; i-- {

		for !s.isEmpty() && temperatures[s.peek()] <= temperatures[i] {
			s.pop()
		}
		if !s.isEmpty() && temperatures[i] < temperatures[s.peek()] {
			ans[i] = s.peek() - i
		}
		s.push(i)
	}
	return ans
}
