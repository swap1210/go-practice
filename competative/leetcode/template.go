package main

import "fmt"

//some common leetcode helper func
//global debug flag
var g_df bool

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//sum of all integers
func sumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func D(a interface{}) {
	if g_df {
		fmt.Print(a)
	}
}

func Dln(a interface{}) {
	if g_df {
		fmt.Println(a)
	}
}

//end common leet code helper func

func fun_name(input int) int {
	ans := 0

	return ans
}

//copy till here
//ctrl+A hold shift click here

func main() {
	//comment this flag set before submissionor ignore for function base exams
	g_df = true
	// g_df = false //default
	Dln("DEBUG MODE\n")
	//end debug ops
	input := 0
	fmt.Println(fun_name(input))
}
