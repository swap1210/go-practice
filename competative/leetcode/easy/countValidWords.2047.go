package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countValidWords("!this  1-s b8d!"))
}

func countValidWords(sentence string) int {
	cnt := 0
	for _, word := range strings.Split(sentence, " ") {

		if check(word) {
			cnt++
		}
	}
	return cnt
}

func check(s string) bool {

	if len(s) == 0 {
		return false
	}
	hiCount, pCount := 0, 0
	for i, c := range s {
		if c == '-' {
			hiCount++
			if hiCount > 1 {
				return false
			}
			if i == 0 || i == len(s)-1 {
				return false
			} else {
				if !isAlpha(s[i-1]) || !isAlpha(s[i+1]) {
					return false
				}
			}
		} else if isAlpha(s[i]) {
			continue
		} else if isPunch(s[i]) {
			pCount++
			if pCount > 1 {
				return false
			}
			if i != len(s)-1 {
				return false
			}

		} else if isNum(s[i]) {
			return false
		} else {
			return false
		}
	}

	return true
}

func isAlpha(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	} else {
		return false
	}
}

func isNum(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
func isPunch(c byte) bool {
	if c == '!' || c == '.' || c == ',' {
		return true
	} else {
		return false
	}
}
