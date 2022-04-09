package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		var size int
		fmt.Scan(&size)
		num_arr := make([]int, size)
		for i := 0; i < size; i++ {
			fmt.Scan(&num_arr[i])
		}
		sort.Ints(num_arr)
		l := 0
		for _, num := range num_arr {
			if num > l {
				l++
			}
		}
		fmt.Printf("Case #%d: %d\n", t, l)
	}
}
