package main

import "fmt"

func main() {
	// x, n := []int{1, 0, 0, 0, 1}, 1
	// x, n := []int{1, 0, 0, 0, 1}, 2
	x, n := []int{0, 0, 0, 0, 0, 1, 0, 0}, 0
	fmt.Println(canPlaceFlowers(x, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	count := 0
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			emptyLeft := (i == 0) || flowerbed[i-1] == 0
			emptyRight := (i == len(flowerbed)-1) || flowerbed[i+1] == 0

			if emptyLeft && emptyRight && flowerbed[i] == 0 {
				//fmt.Println("Space at ", i)
				flowerbed[i] = 1
				count++
			}
			//fmt.Println("count is ", count)
			if count == n {
				return true
			}
		}
	}
	return count == n
}
