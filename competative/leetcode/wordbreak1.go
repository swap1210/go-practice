package main

import (
	"fmt"
)

func main() {
	// fmt.Println(wordbreak("programcreek", map[string]int{"programcre": 0, "program": 0, "creek": 0}))
	fmt.Println(wordbreak("prgcreek", map[string]int{"prgcre": 0, "prg": 0, "creek": 0}))
}

func wordbreak(s string, dict map[string]int) bool {
	t := make([]bool, len(s)+1)
	t[0] = true
	for i := 0; i < len(s); i++ {

		if !t[i] {
			continue
		}

		for v := range dict {

			len1 := len(v)
			end := i + len1

			if end > len(s) {
				continue
			}

			fmt.Println(s[i:end], v)
			if s[i:end] == v {
				t[end] = true
			}
		}
	}
	fmt.Println(t)
	return t[len(s)]
}
