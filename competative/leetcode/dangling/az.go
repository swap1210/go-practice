package main

import "fmt"

func helper(s string) int{
	count := 0
	total := 0
	for _,c := range s{
		if c == 'A'{
			count++
		}
		if c == 'Z'{
			total += count
		}
	}
	return total
}

func azcounter(s string) int{
	a := helper("A"+s)
	z := helper(s+"Z")

	if(a>z){
		return a
	}else{
		return z
	}
}

func main(){
	fmt.Print(azcounter("ATAAZAAAZZAAAAZAZZZZAZZHZZAA"));
}