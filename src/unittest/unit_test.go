package main

import (
	//"fmt"
	"testing"
)

func TestAdd(t *testing.T){
	ans := Add(5,6)

	if ans!=11 {
		//fmt.println("Incorrect add value",)
		t.Errorf("add was incorrect, got: %d, want: %d", ans, 11)
	}
}