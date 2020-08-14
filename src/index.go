package main

import (
	    "fmt"
		"reflect"
	)

var a = "Zeus says die"
var b int32 = 10
var d float32 = 30.02
func main(){
//b = 10
c := 20
fmt.Println(c);
	fmt.Println("Hello earth!!")
	fmt.Println(d);
	fmt.Print(reflect.TypeOf(d));
	fmt.Println(a,b,"times");
	fmt.Printf("%T", c);
	fmt.Printf("\n%T", a);
}