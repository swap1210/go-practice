package main

import "fmt"

func main() {
	a := []int{3, 56, 67, 111, 219}
	b := []int{13, 78, 112}
	fmt.Println(mergeSortedArray(a, len(a), b, len(b)))
}

func mergeSortedArray(a []int, m int, b []int, n int) []int {
	res := make([]int, m+n)
	k := m + n - 1
	m--
	n--
	for k >= 0 {
		if n < 0 || (m >= 0 && a[m] > b[n]) {
			res[k] = a[m]
			m--
		} else {
			res[k] = b[n]
			n--
		}
		k--
	}

	return res
}
