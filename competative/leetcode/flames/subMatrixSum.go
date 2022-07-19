package main

import "fmt"

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	//fmt.Println(matrix[0][0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// fmt.Println("i,j", i+1, j+1)
			dp[i+1][j+1] = matrix[i][j]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] += dp[i-1][j] + dp[i-1][j] + dp[i][j-1] + dp[i-1][j-1]
		}
	}

	cnt := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			low := 1
			hi := min(n-i, m-i) + 1

			found := false

			for low <= hi {

				mid := (low + hi) / 2

				ni := i + mid - 1
				nj := j + mid - 1
				fmt.Println("ni,nj", ni, nj)
				sum := dp[ni][nj] - dp[ni-1][nj] - dp[ni][nj-1] - dp[ni-1][nj-1]

				if sum >= target {

					if sum == target {
						found = true
					}
					hi = mid - 1
				} else {
					low = mid + 1
				}
			}

			if found {
				cnt++
			}
		}
	}

	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
