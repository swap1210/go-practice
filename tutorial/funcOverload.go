package main

import "fmt"

func main(){
 add(4, 5)

 multi, div := multi(12, 6)
 fmt.Println("mul = ",  multi)
 fmt.Println("Div = ", div)

}

func add(a int, b int){
fmt.Println("sum = ", a + b)
}

func multi(a int, b int)(int, int){
return a * b, a / b
}