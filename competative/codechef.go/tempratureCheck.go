package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	t_str, _, _ := rdr.ReadLine()
	t, _ := strconv.Atoi(string(t_str))
	fmt.Println("Test cases are")
	for i := 0; i < t; i++ {
		s, _, _ := rdr.ReadLine()
		val, _ := strconv.Atoi(string(s))
		fmt.Println(checkTemp(val))
	}
}

func checkTemp(a int) string {
	if a > 20 {
		return "HOT"
	} else {
		return "COLD"
	}
}
