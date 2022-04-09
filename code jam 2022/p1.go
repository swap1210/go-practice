package main

import (
	"fmt"
)

func main(){
	var T int
	 _,_ = fmt.Scan(&T)

	for k:=0;k<T;k++{
		R,C := 0,0
		fmt.Scan(&R, &C)
		m := make([][]rune,R+R+1)
		odd := true
		for r:=0;r<len(m);r++ {
			m[r] = make([]rune,C+C+1)
			f_s := true
			for c:=0;c<len(m[r]);c++{
				//first
				if r <= 1 && c <= 1 {
					m[r][c] = '.'
					continue
				}
				if(odd){
					if(f_s){
						m[r][c] = '+'
					}else{
						m[r][c] = '-'
					}
					f_s = !f_s
				}else{
					if(f_s){
						m[r][c] = '|'
					}else{
						m[r][c] = '.'
					}
					f_s = !f_s
				}
			}
			odd = !odd
		}
		fmt.Printf("Case #%d:\n",(k+1))
		for i := range m {
			for j := range m[i]{
				fmt.Print(string(m[i][j]))
			}
			fmt.Println()
		}
	}
}