package main

import "fmt"

type Stack struct {
	arr []rune
}

func (obj *Stack) isEmpty() bool {
	return len(obj.arr) == 0
}

func (obj *Stack) push(n rune) {
	obj.arr = append([]rune{n}, obj.arr...)
}

func (obj *Stack) pop() rune {
	if obj.isEmpty() {
		return -1
	}
	t := obj.arr[0]
	obj.arr = obj.arr[1:]
	return t
}

func (obj *Stack) peek() rune {
	if obj.isEmpty() {
		return -1
	}
	return obj.arr[0]
}

func main() {
	fmt.Println(reverse(""))
	// t.push(34)
	// t.push(78)
	// fmt.Println(t.peek())
	// fmt.Println(t.pop())
	// fmt.Println(t.pop())
	// fmt.Println(t.pop())
	fmt.Println(balanceOperation("(((65+))"))
}

func reverse(s string) string {
	t := &Stack{}

	for _, v := range s {
		t.push(v)
	}
	res := ""
	for !t.isEmpty() {
		res += string(t.pop())
	}
	return res
}

func balanceOperation(s string) bool {
	t := &Stack{}

	for _, v := range s {
		if v == '(' {
			t.push(v)
		}
		if v == ')' {
			if t.isEmpty() {
				return false
			}
			t.pop()
		}
	}
	return t.isEmpty()
}
