package main

import "fmt"

func main(){
add(1, 2)
add(1,2,3)
nums := []int{1, 2, 3, 4}
add(nums...)
}

func add(a ...int){

	total := 0
	for _, data := range a{
		total += data
	}
	fmt.Println(total)
}

