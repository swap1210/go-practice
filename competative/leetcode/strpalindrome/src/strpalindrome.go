package main

import (
	"fmt"
)

func main(){
	//x := "racecar"
	//fmt.Print(x[1:7]);
	fmt.Print(longStr("22racecar21212"))
	//fmt.Print(palindrone("racecar"))
}
func longStr(s string) string{
	if(s=="" ||  len(s) < 1){ 
		return ""
	}

	start :=0
	end := 0

	for i:=0; i<len(s); i++{
		len1 := expandMiddle(s,i,i)
		len2 := expandMiddle(s,i,i+1)
		len:=0
		if(len2>len1){
			len = len2;
		}else{
			len = len1;
		}

		if(len > end - start){
			start = i - ((len-1)/2)
			end = i + (len/2)
		}
	}

	return s[start+1:(end)]
}

func expandMiddle(s string, left int, right int) int{
	if(s=="" || left>right){
		 return 0
		}
	for ;(left >=0 && right < len(s) && s[left]==s[right]); {
		left--
		right++
	}

	return right-left + 1
}