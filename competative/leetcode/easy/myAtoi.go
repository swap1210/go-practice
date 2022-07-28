package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "words and 987"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	t := ""

	for _, str := range s {
		if str >= '0' && str <= '9' || str == '-' {
			t += string(str)
		}
	}
	v, _ := strconv.Atoi(t)
	return v
}
