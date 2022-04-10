package main

import (
	"fmt"
	"encoding/xml"
)

var myxml = []byte(
    `<emp>
    <fname>Zeus</fname>
    <age>22</age>
    <location>BDC6E</location>
    </emp>`)

type student struct{
	Name string `xml:"fname"`
	Age int `xml:"age"`
	Location string `xml:"location"`
}

func main(){
   var stu student
   xml.Unmarshal(myxml, &stu)
   fmt.Println(stu)
}