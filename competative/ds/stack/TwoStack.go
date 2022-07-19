package main

import "fmt"

type TwoStack struct {
	arr  []int
	top1 int
	top2 int
	size int
}

func (o *TwoStack) push1(v int) {
	if o.top1 < o.top2-1 || o.top1 == -1 {
		o.arr[o.top1+1] = v
		o.top1++
	}
}
func (o *TwoStack) push2(v int) {
	if o.top1 < o.top2-1 || o.top2 == o.size {
		o.arr[o.top2-1] = v
		o.top2--
	}
}

func (o *TwoStack) pop1() int {
	if o.isEmpty1() {
		return -1
	}
	t := o.arr[o.top1]
	o.top1--
	return t
}

func (o *TwoStack) pop2() int {
	if o.isEmpty2() {
		return -1
	}
	t := o.arr[o.top2]
	o.top2--
	return t
}

func (o *TwoStack) peek1() int {
	if o.isEmpty1() {
		return -1
	}
	return o.arr[o.top1]
}

func (o *TwoStack) peek2() int {
	if o.isEmpty2() {
		return -1
	}
	return o.arr[o.top2]
}

func (o *TwoStack) isEmpty1() bool {
	return o.top1 == -1
}

func (o *TwoStack) isEmpty2() bool {
	return o.top2 == -1
}

func (o *TwoStack) isFull1() bool {
	return o.top1 == o.top2 && o.top1 != -1
}

func (o *TwoStack) isFull2() bool {
	return o.top1 == o.top2 && o.top2 != -1
}

func (o *TwoStack) newTwoStack(s int) {
	o.size = s
	o.arr = make([]int, s)
	o.top1, o.top2 = -1, s
}

func main() {
	t := &TwoStack{}
	t.newTwoStack(6)

	fmt.Println(t.isEmpty1(), t.isEmpty2(), t.isFull1(), t.isFull2())
	fmt.Println(t.pop1())
	t.push1(1)
	t.push1(5)
	t.push1(34)
	t.push1(6)
	t.push2(0)
	t.push2(56)
	fmt.Println(t)
	t.pop2()
	t.pop2()
	fmt.Println(t.peek1())
	fmt.Println(t.peek2())
	// fmt.Println(t)
	// fmt.Println(t)
	fmt.Println(t.isEmpty1(), t.isEmpty2(), t.isFull1(), t.isFull2())
}
