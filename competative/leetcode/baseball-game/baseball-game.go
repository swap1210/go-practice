package main

import (
	"fmt"
	"strconv"
)

//some common leetcode helper func
//global debug flag
var g_df bool

//sum of all integers
func sumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func D(a ...interface{}) {
	if g_df {
		fmt.Print(a)
	}
}

func Dln(a ...interface{}) {
	if g_df {
		fmt.Println(a)
	}
}

//end common leet code helper func

func calPoints(ops []string) int {
	ans := 0
	//sum till now
	sum_tn := 0
	q := []int{}
	for _, v := range ops {
		D("op", v)
		if temp_int, err := strconv.Atoi(v); err == nil {
			sum_tn += temp_int
			q = append(q, temp_int)
		} else {
			//operation
			if v == "+" {
				q = append(q, q[len(q)-1]+q[len(q)-2])
				sum_tn += q[len(q)-1]
			} else if v == "D" {
				//add double value
				q = append(q, q[len(q)-1]+q[len(q)-1])
				sum_tn += q[len(q)-1]
			} else if v == "C" {
				//remove last value
				sum_tn -= q[len(q)-1]
				q = q[:len(q)-1]
			}
		}
		Dln(q, "sum till now", sum_tn)
	}
	// D(q)
	for _, v := range q {
		ans += v
	}
	return ans
}

//copy till here

func main() {
	// g_df = true
	// g_df = false //default
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	Dln("DEBUG MODE")
	fmt.Println(calPoints(ops))
}
