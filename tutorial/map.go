package main

import "fmt"

func main(){

	mymap := map[int]string{10:"helo",20:"this",30:"is"}
	fmt.Println(mymap)
	mymap[20]="I"
	fmt.Println(mymap)
	delete(mymap, 20)
	fmt.Println(mymap)
	mymap[40]="new"
	fmt.Println(mymap)
}