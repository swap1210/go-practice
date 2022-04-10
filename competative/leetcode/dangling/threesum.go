package main

import (
	"fmt"
	"sort"
)
func threeSum(nums []int) [][]int {
	fin_arr := [][]int{}
	fin_arr = make([][]int,0)

    sort.Ints(nums)
	for i :=0 ; i < len(nums)-2; i++{
		if(i ==0 || (i > 0 && nums[i] != nums[i-1])){
			var low,high,sum int=i+1,len(nums)-1,0-nums[i]
			for low < high{
				if nums[low] + nums[high] == sum{
						fin_arr = append(fin_arr,[]int{nums[i],nums[low],nums[high]})
						
						for low < high && nums[low] == nums[low+1]{ low++}
						for low < high && nums[high] == nums[high-1]{ high--}
						low++
						high--
				}else if (nums[low] + nums[high] > sum){
					high--
				}else{
					low++
				}
			}
		}
	}

	return fin_arr;
}
func main(){
	fmt.Println("sol ",threeSum([]int{-1,0,1,2,-1,-4}))
}