package main

import (
	"fmt"
	"time"
)

func main(){
   today := time.Now()
   fmt.Println(today.Day())
   fmt.Println(today.Local())
   fmt.Println(today.Month())  
   fmt.Println(today.Clock())
}