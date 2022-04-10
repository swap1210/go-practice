package main

import "fmt"

func main(){
	fmt.Println(div(12, 0))
	fmt.Println(div(10, 2))
	
	dempanic()
}

func div( a int, b int)int{

	defer func(){
       fmt.Println(recover())
	}() 
	res := a/b
	return res
}

func dempanic(){

	defer func(){
		fmt.Println(recover())
	}()
	panic("Panic")

}