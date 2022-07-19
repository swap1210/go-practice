package main

import "fmt"

var M int

func main() {
	m := 1
	n := 3
	maxMove := 3
	startRow := 0
	startColumn := 1
	fmt.Println(findPaths(m, n, maxMove, startRow, startColumn))
}

func make3DInt(m, n, p int) [][][]int {
	buf := make([]int, m*n*p)

	x := make([][][]int, m)
	for i := range x {
		x[i] = make([][]int, n)
		for j := range x[i] {
			x[i][j] = buf[:p:p]
			buf = buf[p:]
		}
	}
	return x
}

func findPaths(m, n, maxMove, startRow, startColumn int) int {
	M = 1000000007
	memo := make3DInt(m, n, maxMove+1)

	for _, l := range memo {
		for _, sl := range l {
			for i := 0; i < len(sl); i++ {
				sl[i] = -1
			}
		}
	}

	return helper(m, n, maxMove, startRow, startColumn, memo)
}

func helper(m, n, N, i, j int, memo [][][]int) int {
	if i == m || j == n || i < 0 || j < 0 {
		return 1
	}
	if N == 0 {
		return 0
	}

	if memo[i][j][N] >= 0 {
		return memo[i][j][N]
	}

	memo[i][j][N] = ((helper(m, n, N-1, i-1, j, memo)+helper(m, n, N-1, i+1, j, memo))%M + (helper(m, n, N-1, i, j-1, memo)+helper(m, n, N-1, i, j+1, memo))%M) % M
	return memo[i][j][N]
}
