package main

import "fmt"

  type ListNode struct {
      Val int
      Next *ListNode
  }

 func main(){
	x := ListNode{8,&ListNode{7,&ListNode{4,nil}}}
	y := ListNode{9,&ListNode{3,&ListNode{3,nil}}}
	var a *ListNode = addTwoNumbers2(&x,&y)
	
	for cur:=a;cur!=nil;cur = cur.Next{
		fmt.Print(cur.Val)
		if(cur.Next != nil){
			fmt.Print("->")
		}
	}
	// for check:=true;check;{
	// 	fmt.Print(a.Val)
	// 	check = (a.Next != nil)
	// 	if(check){
	// 		fmt.Print("->")
	// 	}
	// 	a = a.Next
	// 	}
 }

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy,sum := new(ListNode),0
	for curr:=dummy; l1 != nil || l2 != nil || sum != 0; curr = curr.Next{
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum%10}
		sum /= 10
	}
    return dummy.Next
}


func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, sum := new(ListNode),0
	for curr:=dummy;(l1 != nil && l2 != nil); curr=curr.Next{
		sum += l1.Val + l2.Val
		curr.Next = &ListNode{Val: sum%10}
		sum /= 10
		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 == nil{
		for curr:=dummy;l1 != nil; curr=curr.Next{
			sum += l1.Val
			curr.Next = &ListNode{Val: sum%10}
			sum /= 10
			l1 = l1.Next
		}
	}else if l1==nil{
		for curr:=dummy;l2 != nil; curr=curr.Next{
			sum += l2.Val
			curr.Next = &ListNode{Val: sum%10}
			sum /= 10
			l2 = l2.Next
		}
	}
	return dummy.Next
}