package main

import "fmt"

func main() {
	x := []int{4, 7, 1, 2, 8}
	r := 3

	res2 := [][]int{}
	HeapPermutation(x, r, func(a []int) {
		c := make([]int, r)
		copy(c, a)
		res2 = append(res2, c)
	})
	for i, rec := range res2 {
		fmt.Println(i, rec)
	}
	fmt.Println(res2)

	// res := []string{}
	// Perm([]rune("1234"), func(a []rune) {
	// 	res = append(res, string(a))
	// })
	// for i, rec := range res {
	// 	fmt.Println(i, rec)
	// }

}

func HeapPermutation(a []int, size int, onComplete func([]int)) {
	if size == 1 {
		onComplete(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1, onComplete)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
