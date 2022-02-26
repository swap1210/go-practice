package main

import (
	"fmt"
	"math"
)

func main(){
	//mat := [row][col]int{{9,8,7,6},{2,3,2,5},{7,8,9,7},{5,4,3,2},{6,7,8,9}}
	mat := [][]int{{9,8,7,6,4,4,2,1},{2,3,2,5,2,2,2,2,2},{7,8,9,7,7,7,7,7,7},{5,4,3,2,4,56,7,3,3}}
	//fmt.Print(palin(mat,0,0,4))
	brutMat(mat)
}

func brutMat(mat [][]int){
	flag := false
	min := int(math.Min(float64(len(mat)),float64(len(mat[0]))))
	row,col := 0,0
	for ;;{
		fmt.Println("Scanning ",row," ",col," size is ",min)
		flag = palin(mat,row,col,min)
		if flag {
			break;
		}
		if row+min < len(mat) && col+min < len(mat[0]) {
			fmt.Println("Adjusting row col")
			if (row+min < len(mat)){
				row++
			}else if (col+min < len(mat[0])){
				col++
			}
		}else{
			fmt.Println("Adjusting size")
			min--
		}
	}
	fmt.Print("Min is ",min)
}

func palin(mat [][]int,startR, startC, size int) bool{
	// row:= len(mat)
	// col:= len(mat[0])
	if size == 1 {
		return true
	}
	flag := true
	for i:=0;i<size;i++{
		if !flag{
			break
		}
		for j:=0;j<size;j++{
			flag = (mat[i][j] == mat[size-1-i][size-1-j])

			if !flag{
				break
			}
		}
		//fmt.Print("\n")
	}
	//fmt.Print(flag)

	return flag
}