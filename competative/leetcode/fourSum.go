package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(num []int, target int) [][]int {
	sort.Ints(num)
	res := [][]int{}
	hash := make(map[[4]int]struct{})
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			temp := [4]int{num[i], num[j], 0, 0}
			k, l := j+1, len(num)-1
			for k < l {
				sum := num[i] + num[j] + num[k] + num[l]
				fmt.Println(i, j, k, l, num[i], num[j], num[k], num[l], num, sum)
				if sum > target {
					l--
				} else if sum < target {
					k++
				} else {
					temp[2], temp[3] = num[k], num[l]
					if _, ok := hash[temp]; !ok {
						hash[[4]int{num[i], num[j], num[k], num[l]}] = struct{}{}
						res = append(res, []int{temp[0], temp[1], temp[2], temp[3]})
					}
					k++
					for k < l && num[k] == num[k-1] {
						k++
					}
				}
			}
		}
	}
	return res
}
