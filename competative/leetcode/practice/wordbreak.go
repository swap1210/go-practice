package main

import "fmt"

func main() {
	s := "swapnil"
	t := ""
	t = s[:2] + s[4:]
	fmt.Println(t)
}
