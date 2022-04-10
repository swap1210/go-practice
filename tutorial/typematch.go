package main

import "fmt"

func main(){
	var a int
	var b int
  var c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c = a/b
	fmt.Println("sum =", a + b)
	fmt.Println("sub =", a - b)
	fmt.Println("Div =", c)
}