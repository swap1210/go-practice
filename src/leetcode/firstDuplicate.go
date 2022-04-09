package main

import "fmt"

func main(){
	x := findDuplicate([]int{3,1,3,4,2})
	fmt.Print(x)
}

func findDuplicate(nums []int) int {
    m := make(map[int]int,0)
    
    for _,val := range nums{
        if _,ok := m[val]; ok{
            return val
        }else{
            m[val]=1
        }
    }
    return 0
}