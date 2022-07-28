package main

import "fmt"

func main() {
	s := "Swapnil"
	// for i := 1; i < len(s); i++ {
	// 	s = s[:i-1] + s[i:]
	// }
	// red(s)
	i := 1
	s = s[:i] //+ s[i+1:]
	fmt.Println(s)
}

func red(s string) {
	// if len(s) == 0 {
	// 	return
	// }
	// fmt.Println("Reduce ", s)
	// red(s[1:])
}
