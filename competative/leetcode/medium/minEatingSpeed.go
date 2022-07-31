package main

import (
	"fmt"
	"math"
)

func main() {

	a := []int{3, 6, 7, 11}
	b := 8
	fmt.Println(minEatingSpeed(a, b))
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, max(piles)
	res := r

	for l <= r {
		k := (l + r) / 2
		temp_hours := 0
		fmt.Print(l, r, " for speed ", k)
		for _, v := range piles {
			temp_hours += int(math.Ceil(float64(v) / float64(k)))
		}
		fmt.Println(" got ", temp_hours)
		if temp_hours <= h {
			res = min(res, k)
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}

func max(a []int) int {
	max := a[0]
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
