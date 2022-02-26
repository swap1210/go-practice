package main

import (
	"fmt"
)

func main(){
	//var testSample = "abcabcbb"
	//var testSample = "bbbbb"
	 var testSample = "pwwkew"
	fmt.Print(lengthOfLongestSubstring(testSample))
}

func lengthOfLongestSubstring(s string) int {
	a_p,b_p,max := 0,0,0
	m := make(map[rune]int)
	for ;b_p<len(s);{
		if(m[s[b_p]]==0){
			m[s[b_p]]=1
			b_p++
			if(max < len(m)){
				max = len(m)
			}
		}else{
			delete(m,s[a_p])
			a_p++
		}
	}
    return max
}