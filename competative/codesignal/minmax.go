package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minmaxpercent([]int{1, 4, 5, 11}))
}

func minmaxpercent(arr []int) float64 {
	min, max := math.MaxInt64, math.MinInt64

	for _, v := range arr {
		if min > v {
			min = v
		}

		if max < v {
			max = v
		}
	}

	if min == 0 {
		return -1
	} else {
		return math.Round((float64(max-min) * 100.0) / float64(max))
	}
}
