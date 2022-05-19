package main

import (
	"fmt"
)

func main() {
	//1,2,3,4,5,6,7,8
	//k = 9
	// parcel := []int{2, 3, 6, 10, 11}
	parcel := []int{1, 2, 3, 6}

	fmt.Println("v ", getMinParcel_V(parcel, 5))

	fmt.Println("s ", getMinParcel(parcel, 5))
}
func getMinParcel(parcel []int, k int) int {
	sum := 0
	m := make(map[int]int)
	for _, v := range parcel {
		m[v] = 0
	}
	pending_cnt := k - len(m)

	for i := 1; pending_cnt > 0; i++ {
		if _, ok := m[i]; !ok {
			sum += i
			pending_cnt--
		}
	}

	return sum
}

//vaibhav approach
func getMinParcel_V(parcel []int, k int) int {
	m := make(map[int]int)
	for _, v := range parcel {
		m[v] = 0
	}
	i := 1
	n := len(parcel)
	parcelpending := k - n
	totalCost := 0

	for parcelpending > 0 {
		if _, ok := m[i]; !ok {
			// fmt.Println(i, "not in ", m)
			totalCost += i
			parcelpending -= 1
		}
		i += 1
	}
	return totalCost
}
