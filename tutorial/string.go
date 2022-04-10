package main

import "fmt"

func main(){
/*	arr := []int{2, 4, 5, 6, 7}

	for _, data := range arr{
		fmt.Println(" ", data)
	}

	arr1 := []string{"helo","this","is"}
  for _, data1 := range arr1{
	  fmt.Println(" ", data1)
  } */
  /* var intarr [3]int
   intarr[0] = 12
   intarr[1] = 23
   intarr[2] = 75
   fmt.Println(intarr)
}*/

var intarr1 [4]int
 //var j int = 5
for data2 := range intarr1{
	fmt.Scanln(intarr1[data2]) 
}

for data3 := range intarr1{
	fmt.Println(intarr1[data3])
}
}