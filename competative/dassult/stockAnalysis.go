package main

import (
	"fmt"
	"strings"
)

func main() {
	x := [][]string{
		{"P1:a", "P3:b", "P5:x"},
		{"P1:b", "P2:q", "P5:x"},
	}
	fmt.Println(sorter(x))
}

func sorter(num [][]string) []string {
	s := []string{}
	m := make(map[string]string)
	for _, r := range num {
		for _, v := range r {
			ss := strings.Split(v, ":")
			a := ss[0]
			b := ss[1]
			m[a] = b
		}
	}

	for _, v := range m {
		s = append(s, v)
	}
	return s
}
