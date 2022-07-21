package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "Sw"
	// fmt.Println(s[3:])
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	// fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))

}

func numMatchingSubseq(s string, words []string) int {
	cnt := 0
	for _, v := range words {
		if matchSequence(s, v) {
			fmt.Println(v, " is true")
			cnt++
		} else {
			fmt.Println(v, " is false")
		}
	}
	return cnt
}

func matchSequence(s, v string) bool {
	// fmt.Println("finding ", v, " in ", s)
	if len(v) == 0 {
		return true
	}

	f := strings.Index(s, v[:1])
	if f == -1 {
		return false
	} else {
		return matchSequence(s[f+1:], v[1:])
	}

}
