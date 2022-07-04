package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(threeSum([]int{-4, -4, 4, 0, 5, -1}, 0))
}

func threeSum(arr []int, sum int) [][]int {
	sort.Ints(arr)
	var res [][]int
	for i := range arr {

		if i == len(arr)-1 {
			break
		}

		if i == 0 || arr[i] > arr[i-1] {
			start, end := i+1, len(arr)-1
			if i == end {
				break
			}

			for start < end {
				// fmt.Println(i, start, end, arr[i], arr[start], arr[end])
				if sum > arr[i]+arr[start]+arr[end] {
					start++
				} else if sum < arr[i]+arr[start]+arr[end] {
					end--
				} else {
					res = append(res, []int{arr[i], arr[start], arr[end]})
					fmt.Println("Index ", i, start, end)
					start++
					end--
				}
			}
		}

	}

	return res
}
