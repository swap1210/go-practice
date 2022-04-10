// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"fmt"
)

func main() {
	x := []int{6, 4, 12, 1}
	x = append(x, 3)
	fmt.Print(x)
}

// func lastStoneWeight(s []int) int {
// 	myMH := newMaxHeap(len(s))
// 	for i := 0; i < len(s); i++ {
// 		myMH.insert(s[i])
// 	}
// 	fmt.Println("init", myMH.heapArray)
// 	for myMH.size > 1 {
// 		temp := myMH.heapArray[0] - myMH.heapArray[1]
// 		myMH.remove()
// 		myMH.remove()
// 		if temp > 0 {
// 			myMH.insert(temp)
// 		}
// 		fmt.Println(">", (myMH.heapArray))
// 	}
// 	if myMH.size > 0 {
// 		return myMH.heapArray[0]
// 	} else {
// 		return 0
// 	}
// }
