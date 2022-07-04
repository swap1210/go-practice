package main

import "fmt"

func main() {

	fmt.Println(twoSumSortedArr([]int{4, 5, 6, 9, 11, 15, 20}, 20))
}

func twoSumSortedArr(arr []int, sum int) []int {

	left, right := 0, len(arr)-1
	for left < right {
		x := arr[left] + arr[right]
		if x > sum {
			right--
		} else if x < sum {
			left++
		} else {
			return []int{arr[left], arr[right]}
		}
	}

	return nil
}
