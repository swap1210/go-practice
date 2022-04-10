package main

import "fmt"

type triangle struct{
sides int
edges int
color string
}

func main(){
 
	tri := new(triangle)
	tri.sides = 3
	tri.edges = 3
	tri.color = "blue"
	fmt.Println(tri)
 
	var tri1 = new(triangle)
	tri1.sides = 4
	tri1.edges = 5
	tri1.color = "yellow"
	fmt.Println(tri1)

	tri2 := &triangle{}
	tri2.sides = 3
	tri2.edges = 3
	tri2.color = "green"
	fmt.Println(tri2)

}