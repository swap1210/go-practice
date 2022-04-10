package main

import (
	"fmt"
	"reflect"
)

func main(){

	tryingB([]byte("Swapnil"))
	tryingB([]byte{43,255})
	tryingS("Swapnil")
	tryingS(string([]byte{43,255}))
	x := 5435

	fmt.Print(reflect.TypeOf(x))
}

func tryingB(arr_b []byte){
	fmt.Print(arr_b)
}
func tryingS(p_s string){
	fmt.Print(p_s)
}