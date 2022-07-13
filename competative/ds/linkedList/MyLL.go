package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type MyLL struct {
	head *Node
	tail *Node
}

func (obj *MyLL) isEmpty() bool {
	return obj.head == nil
}
func (obj *MyLL) addFirst(v int) {
	t := &Node{value: v}
	if obj.isEmpty() {
		obj.head = t
		obj.tail = t
	} else {
		t.next = obj.head
		obj.head = t
	}
}

func (obj *MyLL) addLast(v int) {
	t := &Node{value: v}
	if obj.isEmpty() {
		obj.head = t
	} else {
		obj.tail.next = t
	}
	obj.tail = t
}

func (obj *MyLL) removeFirst() (int, bool) {
	if obj.isEmpty() {
		return -1, false
	}

	if obj.head == obj.tail {
		v := obj.head.value
		obj.head = nil
		obj.tail = nil
		return v, true
	}

	second := obj.head.next
	v := obj.head.value
	obj.head = second

	return v, true
}

func (obj *MyLL) removeLast() (int, bool) {
	if obj.isEmpty() {
		return -1, false
	}

	if obj.head == obj.tail {
		v := obj.head.value
		obj.head = nil
		obj.tail = nil
		return v, true
	}

	if obj.head.next == nil {
		obj.head = nil
		obj.tail = nil
	}

	t := obj.head

	for t.next.next != nil {
		t = t.next
	}

	v := t.value
	obj.tail = t
	t.next = nil
	return v, true
}

func (obj *MyLL) contains(v int) bool {
	return obj.indexOf(v) != -1
}

func (obj *MyLL) indexOf(toFind int) int {
	if obj.isEmpty() {
		return -1
	}
	turtle := obj.head
	ind := 0
	for turtle != nil {
		if turtle.value == toFind {
			return ind
		}
		turtle = turtle.next
		ind++
	}
	return -1
}

func (obj *MyLL) print() {
	if obj.isEmpty() {
		fmt.Println("EMPTY LinkedList")
		return
	}
	turtle := obj.head
	for turtle.next != nil {
		fmt.Print(turtle.value, " -> ")
		turtle = turtle.next
	}
	if turtle != nil {
		fmt.Println(turtle.value)
	}
}

//excercise
func (obj *MyLL) mid() (int, int) {

	if obj.isEmpty() {
		return -1, -1
	}

	t1, t2 := obj.head, obj.head

	for t2 != obj.tail && t2.next != obj.tail {
		t1 = t1.next
		t2 = t2.next.next
	}

	if t2 == obj.tail {
		return t1.value, -1
	} else {
		return t1.value, t1.next.value
	}
}

func (obj *MyLL) hasLoop() bool {

	if obj.isEmpty() {
		return false
	}

	t1, t2 := obj.head, obj.head

	for t2 != obj.tail && t1 != obj.tail && (t1 == obj.head || t1 != t2) {
		// fmt.Println("t1", "t2", t1.value, t2.value)
		if t1.next != nil {
			t1 = t1.next
		}
		if t2.next == nil || t2.next.next == nil {
			break
		}
		t2 = t2.next.next
	}

	if t1 == t2 {
		return true
	} else {
		return false
	}
}

func (obj *MyLL) makeLoop() {
	obj.tail.next = obj.head.next.next.next
}

func (obj *MyLL) reverse() {
	if obj.isEmpty() {
		return
	}
	pre, cur, temp := obj.head, obj.head.next, &Node{}

	for cur != nil {
		temp = cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}

	obj.tail = obj.head
	obj.tail.next = nil
	obj.head = pre
}
func (obj *MyLL) kthNode(k int) int { //fin kth node from the end

	// 12 -> 45 -> 76 -> 90 -> 180
	// |           |
	// p1         p2

	if obj.isEmpty() {
		return -1
	}

	p1, p2 := obj.head, obj.head
	i := 0
	for ; i < k-1 && p2 != obj.tail; i++ {
		p2 = p2.next
	}
	if i < k-1 {
		return -1
	}
	for p2 != obj.tail {
		fmt.Println("p1=", p1.value, " p2=", p2.value)
		p1 = p1.next
		p2 = p2.next
	}

	return p1.value
}

func main() {
	test_obj := MyLL{}

	// test_obj.print()
	test_obj.addLast(10)
	test_obj.addFirst(30)
	test_obj.addFirst(300)
	test_obj.addFirst(3000)
	test_obj.addLast(20)
	test_obj.addLast(29)
	test_obj.addLast(58)
	test_obj.addLast(19)
	test_obj.addLast(47)
	test_obj.addLast(32)
	test_obj.addLast(19)
	test_obj.print()
	// fmt.Println(test_obj.contains(10))

	// test_obj.print()
	// test_obj.removeLast()
	// test_obj.print()
	// test_obj.removeLast()
	// test_obj.print()
	// test_obj.removeLast()
	// test_obj.print()
	// test_obj.removeLast()
	// test_obj.print()
	// test_obj.removeLast()
	// test_obj.removeLast()
	// test_obj.removeLast()
	// test_obj.removeLast()
	// test_obj.makeLoop()
	// test_obj.print()
	// x, y := test_obj.mid()
	// fmt.Println("Mid", x, y)
	// test_obj.reverse()
	// test_obj.print()
	// fmt.Println("Has loop ", test_obj.hasLoop())
	k := 5
	fmt.Println(k, "th value is ", test_obj.kthNode(k))
}
