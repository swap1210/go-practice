package main

import "fmt"

//google interview May 2022
func main() {
	fmt.Println(solution("aabcdea"))
}

func solution(str string) int {
	//println("looking for " + str)
	if len(str) == 1 {
		return 1
	}
	char_map := make(map[rune]int)
	for _, v := range str {
		if _, ok := char_map[v]; ok {
			char_map[v]++
		} else {
			char_map[v] = 1
		}
	}

	last_val := 0
	no_break := true
	//last_c := 'a'
	for _, v := range char_map {
		if last_val != v && last_val != 0 {
			//fmt.Printf("%s %s %d != %s %d\n", str, string(last_c), last_val, string(k), v)
			no_break = false
			break
		}
		last_val = v
		//last_c = k
	}
	if no_break {
		return len(str)
	} else {
		return max2(solution(str[1:]), solution(str[:len(str)-1]))
	}
}

func max2(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
