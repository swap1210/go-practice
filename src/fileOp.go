package main 

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
f, _ := os.Create("content2.txt")
fmt.Println(os.Stat("content2.txt"))
bdata := []byte("Happy Learning is always required")
f.Write(bdata)
fmt.Println("data written")
f.Close()
//fmt.Println(os.Stat("content2.txt"))
data, _ := ioutil.ReadFile("content2.txt")
fmt.Println(string(data))
}