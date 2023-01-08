package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(detectCapitalUse("Hello"))
}

func detectCapitalUse(word string) bool {
	if word == strings.ToLower(word) || word == strings.ToUpper(word) || word == strings.ToUpper(word[:1])+strings.ToLower(word[1:]) {
		return true
	} else {
		return false
	}
}
