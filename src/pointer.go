package main

import "fmt"

func main(){
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(*b)
}