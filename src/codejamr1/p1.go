package main

import (
	"bytes"
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T) //T = 1 //
	for t := 1; t <= T; t++ {
		var str string
		fmt.Scan(&str) //str = "SWAP" //
		l := len(str)

		var buffer bytes.Buffer
		for i, ch := range str {
			fmt.Println(i)
			buffer.WriteRune(ch)
			if i < l-1 && ch < rune(str[i+1]) {
				buffer.WriteRune(ch)
			}
		}
		fmt.Printf("Case #%d: %s\n", t, buffer.String())
	}
}
