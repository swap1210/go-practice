package main

import "fmt"

func main(){

	s := "SwaSpnilPS"
	t := ""

	set := make (map[rune]int,len(s))

	for _,r := range(s){
		if _,ok := set[r]; !ok{
			set[rune(r)] = 1
			t += string(r)
		}
	}

	// for i := len(s)-1;i>-1;i--{
	// 	if _,ok := set[rune(s[i])]; !ok{
	// 		set[rune(s[i])] = 1
	// 	}
	// }

	fmt.Printf(string(t))

}