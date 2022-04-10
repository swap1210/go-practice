package main

import (
	"fmt"
	"strings"
)

func main(){
	info := "hello"
	fmt.Println(strings.Compare(info, "hello")) 
	fmt.Println(strings.Contains(info, "hel")) 
	fmt.Println(strings.Count(info, "l"))
	fmt.Println(len(info)) 
	fmt.Println(strings.Index(info, " ")) 
	fmt.Println(strings.Split(info, "l")) 
	fmt.Println(strings.ToUpper(info)) 
	fmt.Println(strings.Replace(info, "l", "L", 1)) 
	fmt.Println(strings.Contains(info, "hel")) 
	fmt.Println(info[0:2]) 
	
	
}

