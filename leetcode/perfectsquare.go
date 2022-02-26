package main

import (
	"fmt"
)

func main(){
	test := [2]int{12,13}

	for _,v := range test{
		fmt.Print(v,"-",numSquaresB(v),"\n")
	}
}

func numSquaresB(n int) int {
	var ns int = 0
	slc := []int{1}

	for i, j:=2, 4; n%j==0 || j<=n; i,j =i+1, (i+1)*(i+1) {
		// fmt.Printf("%d %d\n",i,j)
		slc = append([]int{j},slc...)
	}

	fmt.Print(slc,"\n")

	// for _, v := range slc {
	// 	fmt.Print(v," ")
	// }
	var mod,div = 0,n
	for j := 0; j < len(slc); j++ {
		fmt.Print(div," = ");
		mod = div%slc[j]
		div = div/slc[j]
		fmt.Print("slc[j] ",slc[j]," div ",div," mod ",mod," ns ",ns)
		if mod==0{
			ns += div
		}else if mod < slc[j]{
			fmt.Print(" .= ",ns,"\n")
			continue
		}else if mod > slc[j]{
			
		}
		fmt.Print(" = ",ns,"\n")
	}
    return ns;
}
