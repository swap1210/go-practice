package main

import (
	"fmt"
	"math"
)

func main() {
	v := []int{-2, 0, 10, -19, 4, 6, -8}
	fmt.Println("=", checkIfExist(v))
}
func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
		if v == 0 {
			if m[0] == 2 {
				return true
			} else {
				continue
			}
		}
		_, ok1 := m[v*2]
		temp := float64(float64(v) / 2.0)
		ok2 := false
		if temp == math.Trunc(temp) { // check if division is whole
			_, ok2 = m[int(temp)]
		}
		if ok1 || ok2 {
			fmt.Println("*", v*2, "/", float64(float64(v)/2.0))
			return true
		}
	}

	return false
}
