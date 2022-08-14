package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {

	tillNow := strs[0]
	for i := 1; i < len(strs); i++ {
		ind := -1
		for len(tillNow) > 0 && ind != 0 {
			ind = strings.Index(strs[i], tillNow)
			if ind != 0 {
				if len(tillNow) >= 1 {
					tillNow = tillNow[:len(tillNow)-1]
				} else {
					tillNow = ""
				}
				continue
			}
		}

	}
	return tillNow
}
