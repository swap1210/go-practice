package main

import "fmt"

func main() {
	// in := []int{1, 2, 2}
	// in := []int{1, 0, 2}
	in := []int{1, 2, 87, 87, 87, 2, 1} //13
	// in := []int{1, 0, 2}
	// in := []int{1, 3, 2, 2, 1} //7
	fmt.Println(candy(in))
}

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	sum := 0
	res := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		res[i] = 1
	}

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(ratings); i++ {
			if i > 0 && ratings[i] > ratings[i-1] && res[i] <= res[i-1] {
				res[i] = res[i-1] + 1
				changed = true
			}
			if i != len(ratings)-1 && ratings[i] > ratings[i+1] && res[i] <= res[i+1] {
				res[i] = res[i+1] + 1
				changed = true
			}
			// fmt.Println(i, res, changed)
		}
	}

	for i := 0; i < len(ratings); i++ {
		sum += res[i]
	}

	return sum
}
