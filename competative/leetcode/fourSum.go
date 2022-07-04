package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(num []int, target int) [][]int {
	sort.Ints(num)
	res := [][]int{}
	hash := make(map[string]int)
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			k := j + 1
			l := len(num) - 1
			for k < l {
				sum := num[i] + num[j] + num[k] + num[l]
				fmt.Println(i, j, k, l, num[i], num[j], num[k], num[l], num, sum)
				if sum > target {
					l--
				} else if sum < target {
					k++
				} else {
					temp := []int{num[i], num[j], num[k], num[l]}
					fmt.Println("checking", num[i], num[j], num[k], num[l])
					sort.Ints(temp)
					str := ""
					for x := 0; x < 4; x++ {
						str += strconv.Itoa(temp[x]) + " "
					}
					if _, ok := hash[str]; !ok {
						hash[str] = 0
						res = append(res, temp)
					}
				}
				k++
				l--
			}
		}
	}
	return res
}
