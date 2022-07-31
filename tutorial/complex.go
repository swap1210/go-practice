package main

import (
	"fmt"
	"strconv"
)

type ComplexNumber struct {
	real    int
	complex int
}

func (xx *ComplexNumber) DoubleNum() {
	xx.real = xx.real * 2
	xx.complex = xx.complex * 4
}

func main() {
	s := "3.14159"
	t, _ := strconv.Atoi(s)
	fmt.Println(s, t)
}
