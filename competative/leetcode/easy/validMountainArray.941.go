package main

import "fmt"

func main() {
	x := []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(x))
}

func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}
	l := 1
	increasingMode := true
	last := arr[0]
	for i := 1; i < len(arr); i++ {
		if increasingMode {
			if arr[i] > last {
				//fmt.Println("inc mode ", last, arr[i])
				l++
			} else if arr[i] < last && l > 1 {
				//fmt.Println("desc started ", last, arr[i])
				l++
				increasingMode = false
			}
		} else {
			//fmt.Println("desc mode ", last, arr[i])
			if arr[i] < last {
				l++
			} else {
				return false
			}
		}
		last = arr[i]
	}
	//fmt.Println("ctr ", l)
	return l == len(arr) && !increasingMode
}
