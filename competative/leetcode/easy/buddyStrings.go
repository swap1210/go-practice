package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ab"))

}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return duplicates(s)
	}

	indexes := []int{}
	for i := range s {
		if len(indexes) > 2 {
			return false
		}

		if s[i] != goal[i] {
			indexes = append(indexes, i)
		}
	}

	if s[indexes[0]] == goal[indexes[1]] && s[indexes[1]] == goal[indexes[0]] {
		return true
	} else {
		return false
	}
}

func duplicates(s string) bool {
	m := make(map[rune]int)
	for _, ch := range s {
		if _, ok := m[ch]; ok {
			return true
		} else {
			m[ch] = 1
		}
	}
	return false
}
