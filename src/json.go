package main

import (
	"fmt"
	"encoding/json"
)

type Lang struct{
	Id int
	Name string
}
func main(){
   text := "[{\"id\":100,\"name\":\"Rohit\"}]"
   jsondata := []byte(text)
   var la []Lang
   json.Unmarshal(jsondata, &la)
   fmt.Println(la)
}