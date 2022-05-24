package main

import "fmt"

func main() {
	fmt.Println(stringRotate("1234567", 3))
}

func stringRotate(s string, count int) string {
	temp := make([]byte, len(s))
	j := 0
	count++
	for i := count; i < len(s); i++ {
		temp[j] = s[i]
		j++
	}
	for k := 0; k < count; k++ {
		temp[j] = s[k]
		j++
	}

	return string(temp)
}
