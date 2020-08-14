package main

import "fmt"

func main(){
	age := 20
	if age>=20 {
		fmt.Println("Can Vote")
	} else if age<20 {
		fmt.Println("should not")
	}
}