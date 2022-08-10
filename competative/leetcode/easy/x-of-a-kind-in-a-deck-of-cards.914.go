package main

import "fmt"

func main() {
	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(hasGroupsSizeX(deck))
}

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for _, c := range deck {
		count[c]++
	}
	g := -1
	for _, v := range count {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	} else {
		return gcd(y%x, x)
	}
}
