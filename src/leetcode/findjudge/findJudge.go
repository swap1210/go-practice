package main

import (
	"fmt"
)

func main(){
	var x = [][]int{{1,2}}
	fmt.Print(findJudge(2,x))
}

func findJudge(n int, trust [][]int) int {
    var count = make([]int,n+1)
	for _,x := range(trust){
		count[x[0]]--
		count[x[1]]++
	}

	for i,x := range(count){
		if x == n-1{
			return i
		}
	}
	return -1
}