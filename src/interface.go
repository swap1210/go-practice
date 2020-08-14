package main

import "fmt"

type human interface {
	sayHello() string
}

type man struct{
    greeting string
}

type women struct{
    greeting string
}

func (m man) sayHello()string {
   return m.greeting
}
func (w women) sayHello()string {
	return w.greeting
 }

 func printGreeting(h human){
	 fmt.Println(h.sayHello())
 }

 func main(){

Rohit := man{greeting: "HAI DUDE"}
ramya := women{greeting: "yes hai dude"}
printGreeting(Rohit)
printGreeting(ramya)

 }
