package main

import (
	"fmt"
	"math"
	"strconv"
)
	
 func calcFraction(num,denom int) (s string){

	if(num == 0){
		return "0"}

	if(denom == 0){
		return ""}
	
	s = ""
	if(num < 0 || denom <0){
		s = "-"
		num = int(math.Abs(float64(num)))
		denom = int(math.Abs(float64(denom)))
	}

	q := num/denom
	var r int = int(math.Mod(float64(num),float64(denom)) * 10)

	s +=strconv.Itoa(q)
	if r==0{ return }

	s += "."

	m := make(map[int]int)

	for ;r != 0; {
		if val, ok := m[int(r)]; ok {
			s1 := s[:val]
			s2  := "("+s[val:len(s)]+")"
			return s1+s2
		}
		m[int(r)]=len(s)
		q = r/denom
		s += strconv.Itoa(q)
		r = (int(r)%denom)*10
	}
	return s
}

func main(){
	fmt.Print(calcFraction(50,22));
}