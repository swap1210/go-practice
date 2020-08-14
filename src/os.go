package main

import (
	"fmt"
	"os"
)

func main(){
   fp, _ := os.Create("newfile.js")
   fmt.Println("File created", fp)
    
   md := os.Mkdir("newFolder", 0644)
   fmt.Println("Folder created", md)
}