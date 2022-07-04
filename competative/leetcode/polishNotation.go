package main

import (
	"fmt"
	"strconv"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) peek() string {
	if s.isEmpty() {
		return ""
	} else {
		return (*s)[len(*s)-1]
	}
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		ele := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return ele, true
	}
}

func (s *Stack) push(ele string) {
	*s = append(*s, ele)
}

func main() {
	arr := []string{"2", "1", "+", "3", "*"}
	fmt.Println("polishNotation ", polishNotation(arr))
}

func polishNotation(arr []string) int {
	var st Stack
	for _, s := range arr {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			e1, _ := st.pop()
			e2, _ := st.pop()

			v1, _ := strconv.Atoi(e1)
			v2, _ := strconv.Atoi(e2)

			if s == "+" {
				st.push(strconv.Itoa(v1 + v2))
			} else if s == "-" {
				st.push(strconv.Itoa(v1 - v2))
			} else if s == "*" {
				st.push(strconv.Itoa(v1 * v2))
			} else if s == "/" {
				st.push(strconv.Itoa(v1 / v2))
			}
		} else {
			st.push(s)
		}
	}

	res, _ := st.pop()
	res_i, _ := strconv.Atoi(res)
	return res_i
}
