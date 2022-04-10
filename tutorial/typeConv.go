package main

import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 42
	fmt.Println("%v", "%T\n", i, i)
	var j string
	j = strconv.Itoa(i)
	fmt.Println("%v", "%T\n", j, j)
}