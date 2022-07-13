package main

import "fmt"

func main() {
	// s1, s2, s3 := "aabcc", "dbbca", "aadbbbaccc"
	s1, s2, s3 := "aabcc", "dbbca", "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}

	l1--
	l2--
	l3--

	for l3 >= 0 {
		if l2 < 0 || (l1 >= 0 && s1[l1] == s3[l3]) {
			l1--
		} else if l1 < 0 || (l2 >= 0 && s2[l2] == s3[l3]) {
			l2--
		}
		l3--
	}

	if l3 == -1 {
		return true
	} else {
		return false
	}
}
