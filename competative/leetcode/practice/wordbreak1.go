package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordbreak("programcreek", map[string]int{"programcre": 0, "program": 0, "creek": 0}))
}

func wordbreak(s string, dict map[string]int) bool {
	return wordBreakHelper(s, dict, 0)
}

func wordBreakHelper(s string, dict map[string]int, start int) bool {

	if start == len(s) {
		return true
	}

	for v := range dict {
		len1 := len(v)
		end := start + len1

		//if length is greater continue
		if end > len(s) {
			continue
		}

		fmt.Println("Matching with ", s[start:start+len1], v)
		if v == s[start:start+len1] {
			//continue searching for words after 1st dictionary letter
			if wordBreakHelper(s, dict, start+len1) {
				return true
			}
		}
	}
	return false
}
