package main

import "fmt"

func main(){
 defer   first()
  defer second()
   defer	third()
}

func first(){
fmt.Println("1st")
}

func second(){
	fmt.Println("2nd")
}
func third(){
	fmt.Println("3rd")
}