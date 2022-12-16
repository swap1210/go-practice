package main

import (
	"fmt"
)

type Freq struct {
	key  rune
	freq int
}

func main() {
	m := make(map[int]*Freq)
	f := &Freq{546, 2}
	m[2] = f
	m[2].freq = 21
	fmt.Println(m[2])
	//1,2,3,4,5,6,7,8
	//k = 9
	// parcel := []int{2, 3, 6, 10, 11}

	// fmt.Println("v ", getMinParcel_V(parcel, 9))

}

//vaibhav approach
func getMinParcel_V(parcel []int, k int) int {
	k = 9
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
