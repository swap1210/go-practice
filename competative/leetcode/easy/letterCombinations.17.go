package main

import "fmt"

func main() {
	x := "23"
	fmt.Println(letterCombinations(x))
}

type SS []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[byte]SS)
	m['2'] = []string{"a", "b", "c"}
	m['3'] = []string{"d", "e", "f"}
	m['4'] = []string{"g", "h", "i"}
	m['5'] = []string{"j", "k", "l"}
	m['6'] = []string{"m", "n", "o"}
	m['7'] = []string{"p", "q", "r", "s"}
	m['8'] = []string{"t", "u", "v"}
	m['9'] = []string{"w", "x", "y", "z"}

	res := []string{}

	q := []string{""}

	for len(q) != 0 {
		s := q[len(q)-1]
		q = q[:len(q)-1]
		if len(s) == len(digits) {
			res = append(res, s)
		} else {
			val := m[digits[len(s)]]
			for _, i := range val {
				q = append(q, s+i)
			}
		}
	}
	return res
}
