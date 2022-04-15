package main

import "fmt"

type A struct {
	attribute1 int
}

func f1() string {
	return "message from f1() 🫠"
}

func (a *A) myMethod() string {
	a.attribute1 = 45
	if a.attribute1 > 0 {
		return "positive"
	} else {
		return "negative"
	}
}

func main() {
	println(f1())
	a_obj := &A{attribute1: 67}
	fmt.Printf("%s", a_obj.myMethod())

	fmt.Printf("%v", a_obj)

}
