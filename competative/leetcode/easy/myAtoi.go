package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// s := "words and 987"
	// s := "4193 with words"
	s := " "
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	t := ""
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if len(s[i:]) == 0 {
		return 0
	}
	sign := '+'
	if s[i] == '-' {
		sign = '-'
		i++
	} else if s[i] == '+' {
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		t += string(s[i])
		i++
	}
	if len(t) == 0 {
		return 0
	}
	v, _ := strconv.Atoi(t)
	if sign != '+' {
		v *= -1
	}
	if v > math.MaxInt32 {
		v = math.MaxInt32
	} else if v < math.MinInt32 {
		v = math.MinInt32
	}
	return v
}
