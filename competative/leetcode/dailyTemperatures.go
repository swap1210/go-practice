package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, t := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if t > temperatures[stack[len(stack)-1]] {
			for len(stack) > 0 &&
				t > temperatures[stack[len(stack)-1]] {
				ans[stack[len(stack)-1]] =
					i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}
	return ans
}
