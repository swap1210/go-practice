package main

import "fmt"

func main() {
	x := []int{4, 7, 1, 2, 8}
	r := 3
	printCombination(x, r)
}

var results [][]int

func printCombination(arr []int, r int) {
	data := make([]int, r)
	combinationUtil(arr, data, 0, 0, r)
	for _, rec := range results {
		fmt.Println(rec)
	}
}

func combinationUtil(arr, data []int, start, index, r int) {
	if index == r {
		tempArray := make([]int, r)
		for j := 0; j < r; j++ {
			tempArray[j] = data[j]
		}
		results = append(results, tempArray)
		return
	}
	end := len(arr) - 1
	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = arr[i]
		combinationUtil(arr, data, i+1, index+1, r)
	}
}
