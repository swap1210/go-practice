package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x, k := 1111122222, 5
	fmt.Println(followAlgo(x, k))
}

func followAlgo(input, k int) int {
	ans := input
	for len(strconv.Itoa(ans)) > k {
		temp := strconv.Itoa(ans)
		newNum := 0
		for len(temp) > k {
			x, _ := strconv.Atoi(temp[:k])
			fmt.Println(temp, x)
			newNum += digiSum(x)
			if len(temp) < k {
				break
			}
			temp = temp[k:]
		}
		b, _ := strconv.Atoi(temp)
		ans = newNum + b
	}
	return ans
}

func digiSum(a int) int {
	sum := 0
	for a > 0 {
		sum += int(math.Mod(float64(a), float64(10)))
		a /= 10
	}
	return sum
}
