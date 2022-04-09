package main

import "fmt"


func searchMatrix(matrix [][]int, target int) bool {
    i,j := len(matrix)-1,0
    
    for ;i>=0 && j<len(matrix[0]);{
        if matrix[i][j] == target{
            return true
        }
        
        if matrix[i][j] > target{
            i--
        }else{
            j++
        }
    }
    return false
    
}

func main(){
	fmt.Print(searchMatrix([][]int{{12,14},{46,54}},46))
}