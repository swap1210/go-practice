package main

import (
	//"fmt"
	"os"
)

func main(){

	fp, _ := os.OpenFile("content3.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fp.WriteString("\n updated content")
}

