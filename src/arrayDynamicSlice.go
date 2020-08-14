package main

import "fmt"

func main(){

	data := make([]int, 5)
	fmt.Println(cap(data))
data[0] = 568
	////data = append(data, 20)
	//data = append(data, 30)
	//data = append(data, 40)
	//data = append(data, 50)
//	data = append(data, 60)
	//fmt.Println(cap(data))
	fmt.Println(len(data))
	fmt.Println(data)
	data = append(data, 60)
	fmt.Println(len(data))
	fmt.Println(cap(data))
	fmt.Println(data)
}