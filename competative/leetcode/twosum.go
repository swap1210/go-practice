package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2, 8, 11, 15}, 9))
}

func twoSum(arr []int, sum int) bool {

	dict := make(map[int]bool)
	for _, v := range arr {
		if _, ok := dict[v]; !ok {
			dict[v] = true
		}
	}

	for _, v := range arr {
		if _, ok := dict[sum-v]; ok {
			return true
		}
	}

	return false
}
