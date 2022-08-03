package main

import (
	"fmt"
)

func main() {
	//:= []int{0, 1,  2, 3, 4}
	a := []int{1, 2, 10, 5, 7}
	// a := []int{105, 924, 32, 968}
	fmt.Println(canBeIncreasing(a))
	// fmt.Println(arrayWithout(a, 3))
}

func canBeIncreasing(nums []int) bool {
	return canBeIncreasingRec(nums, false)
}

func canBeIncreasingRec(nums []int, reduced bool) bool {
	if len(nums) == 1 {
		//true if reduced from 2 items
		return reduced
	}
	fmt.Println("Scanning ", nums)
	p1, p2 := 0, 1
	for p1 < p2 && p2 < len(nums) {
		if nums[p1] >= nums[p2] {
			if !reduced {
				return canBeIncreasingRec(arrayWithout(nums, p1), true) || canBeIncreasingRec(arrayWithout(nums, p2), true)
			} else {
				return false
			}
		}
		p1++
		p2++
	}
	return true
}

func arrayWithout(a []int, skip int) []int {
	res := make([]int, len(a)-1)
	for k, i := 0, 0; k < len(res); k++ {
		if i != skip {
			res[k] = a[i]
		} else {
			//fmt.Println("Skipping ", skip)
			k--
		}
		i++
	}
	//fmt.Println("New array is ", res)
	return res
}
