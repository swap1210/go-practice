package main

import "fmt"

type vehicle interface {
	sayBudget() int
}

type tata struct{
    budget int
}

type tesla struct{
    budget int
}

func (t tata) sayBudget()int {
   return t.budget
}
func (te tesla) sayBudget()int {
	return te.budget
 }

 func printBudget(v vehicle){
	 fmt.Println(v.sayBudget(), "crores")
 }

 func main(){

tata1 := tata{budget: 50}
tesla1 := tesla{budget: 600}
printBudget(tata1)
printBudget(tesla1)

 }
