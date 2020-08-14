package main

import "fmt"

type employee struct{
	firstname string
	lastname string
	age int
}

func main(){
	x := employee{firstname:"Syed", lastname:"zeus", age:23}
fmt.Println(x);
fmt.Println(x.firstname)

y := employee{}
y.firstname = "don"
y.lastname = "hardy"
y.age = 23
fmt.Println(y)
y.empdetails()
y.empdetailsbyreference(&y)
fmt.Println(y);
}

func (y employee) empdetails(){
	y.lastname = "mass"
    fmt.Println("Emp details are", y.lastname," ",y.age)	
}

func (y employee) empdetailsbyreference(ptr *employee){
	ptr.lastname = "really "
	fmt.Println("Emp details are", ptr)
}