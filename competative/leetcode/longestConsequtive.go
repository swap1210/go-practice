package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(in))
}

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)
	longest, current := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				current++
			} else {
				longest = max(longest, current)
				current = 1
			}
		}
	}

	return max(longest, current)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
