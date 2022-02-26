package main

import "fmt"

func findDuplicates(nums []int) []int {
    var res []int
	trackerList := make(map[int]int,0)
	for _,num := range nums {
		if _, ok := trackerList[num];ok{
			res = append(res,num)
			trackerList[num]++
		}else{
			trackerList[num] = 1
		}
	}
	return res
}

func main(){
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}