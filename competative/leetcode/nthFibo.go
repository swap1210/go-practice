package main

import "fmt"

func main() {
	fmt.Println(3)
}

func fib(n int) int {
	cnt := 0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b, c := 0, 1, -1
	for cnt <= n-2 {
		c = a + b
		fmt.Println(c, "=", a, "+", b)
		a, b = b, c
		cnt++
	}
	return c
}
