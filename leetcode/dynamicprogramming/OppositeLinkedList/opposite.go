package main

import "fmt"

type Node struct{
	Val int
	Next *Node
}

var res string = ""

func main(){

	one := Node{}
	one.Val = 12
	two := &Node{Val:3,Next: nil}
	one.Next = two
	three := &Node{Val:78,Next: nil}
	two.Next = three
	four := &Node{Val:9,Next: nil}
	three.Next = four

	reverse(one);

	res = res[:len(res)-2]
	
	fmt.Print(res)
}

func reverse(one Node) {
	if(one.Next == nil){
		//fmt.Print(one.Val)
		res += fmt.Sprint(one.Val)
	}else{
		reverse(*one.Next)
		// fmt.Print(one.Val)
		res += fmt.Sprint(one.Val)
	}

	res += "<-"
}