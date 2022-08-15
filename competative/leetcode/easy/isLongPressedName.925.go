package main

import (
	"fmt"
)

func main() {
	// x, y := "alex", "aaleex"
	// x := "leelee"
	// y := "lleeelee"
	// x := "saeed"
	// y := "ssaaedd"
	x, y := "a", "b"
	fmt.Println(isLongPressedName(x, y))
}

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}
	nameP, typedP := 0, 0
	for typedP < len(typed) { //loop through longer string
		if nameP < len(name) && name[nameP] == typed[typedP] {
			nameP++
		} else if typedP == 0 || typed[typedP] != typed[typedP-1] {
			return false
		}
		typedP++
	}
	return nameP == len(name)
}
