package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordbreak("catsanddog", map[string]int{"cat": 0, "cats": 0, "sand": 0, "and": 0, "dog": 0}))
}

func wordbreak(s string, dict map[string]int) []string {
	dp := make([][]string, len(s)+1)

	dp[0] = []string{}
	for i := 0; i < len(s); i++ {

		if dp[i] == nil {
			continue
		}

		for word := range dict {
			len_word := len(word)
			end := len_word + i
			if end > len(s) {
				continue
			}

			if s[i:end] == word {
				dp[end] = append(dp[end], word)
			}
		}
	}

	res := []string{}
	dfs(dp, len(s), &res, []string{})
	return res
}

func dfs(dp [][]string, end int, output *[]string, temp []string) {
	if end <= 0 {
		path := temp[len(temp)-1]
		for i := len(temp) - 2; i >= 0; i-- {
			path += " " + temp[i]
		}
		*output = append(*output, path)
		return
	}

	for _, str := range dp[end] {
		temp = append(temp, str)
		dfs(dp, end-len(str), output, temp)
		temp = temp[:len(temp)-1]
	}
}
