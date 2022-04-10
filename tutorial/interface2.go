package main

import "fmt"

type vehicle interface{
	engine()
}

type ducati struct{
	model string
	cost int
}

func (d ducati) engine(){
       fmt.Println("You have purchased",d.model,"worth of",d.cost,"lakhs")
}

func main(){

	a := ducati{"veyron",50}
	a.engine()
}