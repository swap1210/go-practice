package main

import "fmt"

func main(){
	x := [][]int{{1,5,2},{7,4,7},{5,0,9}}
	print(x)
}

func print(mat [][]int){
	for i:=0;i<len(mat[0]);i++{
		for j:=0;j<len(mat[0]);j++{
			fmt.Print(mat[i][j]," ")
		}
		fmt.Println()
	}
}