package main

import "fmt"

type address struct {
   pin string
   city string
}

type student struct {
	id string
	name string
	add address
}

func main(){
a := address{"560043", "bengaluru"}
b := student{"21113", "zeus", a}
fmt.Println(b)
c := student{"423","syed", address{"560043", "bengaluru"}}
fmt.Println(c)
fmt.Println(c.add.city)
}