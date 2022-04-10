package main

import (
	"fmt"
)

func convert(s string, numRows int) string{
	temp := ""
	interval := 2*numRows-2
	size := len(s)
	if(size==1 || size <= numRows || numRows ==1){
		return s
	}
	for r:=0;r<numRows;r++{
		for p:=r;p<size;p+=interval{
			temp += (string)(s[p])
			if(r != 0 && r != (numRows-1) && (p+interval-2*r) < size){
				temp+= (string)(s[p+interval-2*r])
			}
		}
	}
	return temp
}

func main(){
	fmt.Print(convert("ABCDEF",2))
}