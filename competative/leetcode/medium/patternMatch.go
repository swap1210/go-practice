package main

import "fmt"

func main() {
	w := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	p := "abb"
	fmt.Println(findAndReplacePattern(w, p))
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	for _, word := range words {
		if matchs(word, pattern) {
			res = append(res, word)
		}
	}
	return res
}

func matchs(word, pattern string) bool {
	pattern2Word := make(map[byte]byte)
	word2Pattern := make(map[byte]byte)
	for i := range word {
		p, w := pattern[i], word[i]
		if _, ok := pattern2Word[p]; !ok {
			pattern2Word[p] = w
		}
		if _, ok := word2Pattern[w]; !ok {
			word2Pattern[w] = p
		}
		if pattern2Word[p] != w || word2Pattern[w] != p {
			//fmt.Println(word, " is ", false, p)
			return false
		}
	}
	fmt.Println(word, pattern, pattern2Word, word2Pattern)
	return true
}
