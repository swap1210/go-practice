package main

import (
	"fmt"
	"math"
)

func main() {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 2
	k := 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	price := make([]int, n)
	for i, _ := range price {
		price[i] = math.MaxInt64
	}
	price[src] = 0

	for i := 0; i <= k; i++ {
		tempPrice := make([]int, n)
		for ti, tv := range price {
			tempPrice[ti] = tv
		}
		for ind := range flights {
			s, d, p := flights[ind][0], flights[ind][1], flights[ind][2]
			if price[s] == math.MaxInt64 {
				continue
			}
			if price[s]+p < tempPrice[d] {
				tempPrice[d] = price[s] + p
			}
		}
		for ti, tv := range tempPrice {
			price[ti] = tv
		}
	}

	if price[dst] == math.MaxInt64 {
		return -1
	} else {
		return price[dst]
	}
}
