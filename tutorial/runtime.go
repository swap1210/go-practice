package main

import (
	"fmt"
	"runtime"
	"math/rand"
)
func hello(){
	fmt.Println("hello")
}
func main(){
	go hello();
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Compiler)
//random number
	fmt.Println(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100), ",")
}

