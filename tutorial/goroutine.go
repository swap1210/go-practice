package main

import (
	"fmt"
	"time"
	"sync"
)

/*func hello(){
	fmt.Println("Hello goroutin")
}
func main(){
   go hello()
   time.Sleep(1*time.Second)
	fmt.Println("mai func")
}*/

/*func hello(s string){
	for i:=0 ; i < 5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}
func main(){
	go hello("world")
	hello("hi")
}*/
var wg = sync.WaitGroup{}
func hello(){
	for i:=0 ; i < 5;i++{
		//time.Sleep(100*time.Millisecond)
		fmt.Println("hello")
		time.Sleep(1000*time.Millisecond)
	}
	wg.Done()
}
func world(){
	for i:=0 ; i < 5;i++{
		//time.Sleep(100*time.Millisecond)
		fmt.Println("world")
		time.Sleep(1000*time.Millisecond)
	}
	wg.Done()
}
func main(){
	wg.Add(2)
	go hello()
	//time.Sleep(1*time.Second)
	go world()
	fmt.Println("in main line")
	wg.Wait()
	fmt.Println("in last line")
}