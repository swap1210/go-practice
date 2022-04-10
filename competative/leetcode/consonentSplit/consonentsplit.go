package main

import (
	"fmt"
)

func main(){
	fmt.Print(countMovesToLeft("hackerrank"))
}

func countMovesToLeft(str string) int{
	count := 0
	v_found, c_found := false, false   
	vowel :=[]rune{'a','e','i','o','u'}
	for _,c := range(str){
		c_found = stringInSlice(c,vowel)
		v_found = stringInSlice(c,vowel)
		// fmt.Print(string(c)," =>",c_found,v_found,"\n")
		if c_found && v_found{
			count++
			v_found, c_found = false, false   
		}
	}
	return count
}
func stringInSlice(a rune, list []rune) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}