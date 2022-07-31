package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxInt
	res := 0
	for i := 0; i < len(nums); i++ {

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[i] + nums[r]
			diff := Abs(target - sum)
			if diff == 0 {
				return sum
			}
			if diff < min {
				min = diff
				res = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
func Abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}
