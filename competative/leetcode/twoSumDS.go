package main

import "fmt"

type TwoSum map[int]bool

func main() {

	var samp TwoSum

	samp.add(1)
	samp.add(3)
	samp.add(5)
	fmt.Println(samp.find(4))
	fmt.Println(samp.find(7))
}

func (ts *TwoSum) add(val int) {
	if (*ts) == nil {
		(*ts) = make(TwoSum)
	}
	if _, ok := (*ts)[val]; !ok {
		(*ts)[val] = false
	}
}

func (ts *TwoSum) find(val int) bool {

	for v := range *ts {
		if _, ok := (*ts)[val-v]; ok {
			return true
		}
	}
	return false
}
