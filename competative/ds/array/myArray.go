package main

import (
	"fmt"
	"math"
)

type MyArray struct{ d []int }

func (ar *MyArray) insert(a int) {
	ar.d = append(ar.d, a)
}

func (ar *MyArray) delete(i int) {
	if i >= len(ar.d) || i < 0 {
		return
	}
	ar.d = append(ar.d[:i], ar.d[i+1:]...)
}

func (ar *MyArray) indexAt(i int) int {
	return ar.d[i]
}

func (ar *MyArray) indexOf(i int) int {
	for ind, v := range ar.d {
		if v == i {
			return ind
		}
	}
	return -1
}

//excersize

func (ar *MyArray) max() int {
	max := math.MinInt
	for _, v := range ar.d {
		if max < v {
			max = v
		}
	}
	return max
}

func (ar *MyArray) intersect(new_arr *MyArray) MyArray {
	temp := []int{}

	l1, l2 := len(ar.d), len(new_arr.d)

	if l1 < l2 {
		for _, v := range ar.d {
			if new_arr.indexOf(v) > -1 {
				temp = append(temp, v)
			}
		}
	} else {
		for _, v := range new_arr.d {
			if ar.indexOf(v) > -1 {
				temp = append(temp, v)
			}
		}
	}
	return MyArray{d: temp}
}

func (ar *MyArray) insertAt(i, v int) bool {
	if i > len(ar.d) || i < 0 {
		return false
	}
	ar.d = append(ar.d[:i], append([]int{v}, ar.d[i:]...)...)
	return true
}

func (ar *MyArray) reverse() MyArray {
	temp := []int{}

	for i := len(ar.d) - 1; i > -1; i-- {
		temp = append(temp, ar.d[i])
	}
	return MyArray{d: temp}
}

func main() {
	obj := MyArray{d: []int{}}
	obj.insert(10)
	obj.insert(40)
	obj.insert(80)
	// obj.delete(1)
	obj.indexOf(40)
	// fmt.Println(
	// 	obj.indexOf(40))
	// fmt.Println(
	// 	obj.intersect(&MyArray{d: []int{1, 3, 40, 7}}))
	obj.insertAt(3, 110)
	fmt.Println(obj)
}
