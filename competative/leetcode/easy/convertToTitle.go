package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(701))
}

func convertToTitle(a int) string {
	ans := ""
	for a > 0 {
		a--
		ans = string('A'+a%26) + ans
		a /= 26
	}
	return ans
}
