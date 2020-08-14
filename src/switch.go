package main

import "fmt"

func main(){
	ch := "i"
	switch ch{
	case "a": fmt.Println("letter = ",ch)
	case "i": fmt.Println("letter = ", ch)
	default: fmt.Println("No letter")
	}
}