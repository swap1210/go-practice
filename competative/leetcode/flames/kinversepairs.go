package main

import (
	"fmt"
)

func main() {
	fmt.Println(kInversePairs(5, 4))
}

var memo [][]int

func kInversePairs(n int, k int) int {
	memo = make([][]int, 1001)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, 1001)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	return helper(n, k)
}

func helper(n, k int) int {
	if n == 0 {
		return 0
	}
	if k == 0 {
		return 1
	}
	if memo[n][k] != -1 {
		return memo[n][k]
	}
	inv := 0
	for i := 0; i <= min(k, n-1); i++ {
		inv = (inv + helper(n-1, k-i)) % 1000000007
	}
	memo[n][k] = inv
	return inv
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
