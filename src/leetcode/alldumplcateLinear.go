package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
    var res []int
	
	for i :=0;i<len(nums);i++ {
		var indexToBe int = int(math.Abs(float64(nums[i]))) - 1;
		if nums[indexToBe] < 0 {
			res = append(res,indexToBe+1)
		}
		nums[indexToBe] = -1 * nums[indexToBe]
		fmt.Println(nums)
	}
	return res
}

func main(){
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}