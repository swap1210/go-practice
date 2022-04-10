package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 3, 8}
	fmt.Print(lastStoneWeight(a))
}

func lastStoneWeight(s []int) int {
	sort.Ints(s)
	ans := 0
	for i := (len(s) - 2); i > 0; i-- {
		fmt.Print(i-1, i, s, ans)
		if ans == s[i] {
			continue
		}
		if ans > s[i] {
			ans -= s[i]
			continue
		}
		if s[i-1] == s[i] {
			continue
		}
		if s[i-1] < s[i] {
			ans = s[i-1] - s[i]
		}
		fmt.Println(" got", ans)
	}
	fmt.Println(s, ans)
	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
