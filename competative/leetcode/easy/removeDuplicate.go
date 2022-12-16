package main

import "fmt"

func main() {
	x := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Println(x)
	l := removeDuplicates(x)
	fmt.Println("Final ", l)
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	for ind, v := range nums {
		fmt.Println(nums)
		if _, valExists := m[v]; valExists {
			if ind < len(nums)-1 {
				fmt.Println("fixing ", v, " at ", ind)
				nums = append(nums[:ind], nums[ind+1:]...)
			} else {
				nums = nums[:ind-1]
			}
		} else {
			fmt.Println("adding ", v)
			m[v] = v
		}
	}

	fmt.Println(nums)
	return len(nums)
}
