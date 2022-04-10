package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m1 := make(map[int]int, len(nums))
    
    for i, num := range nums {
        if idx, ok := m1[target - num]; ok {
            return []int{idx, i}            
        }
        m1[num] = i
    }
    return []int{0, 0}
}
func main() {
	fmt.Println(twoSum([]int{1, 2, 3},6))
}