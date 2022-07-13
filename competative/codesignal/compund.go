package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(1, 100, 64))
}

func solution(deposit int, rate int, threshold int) int {
	term := float64(deposit)
	ctr := 0
	for term < float64(threshold) {
		ctr++
		term = float64(deposit) * math.Pow(float64(1)+float64(rate)/100, float64(ctr))
		// fmt.Println(term)
	}
	return ctr
}
