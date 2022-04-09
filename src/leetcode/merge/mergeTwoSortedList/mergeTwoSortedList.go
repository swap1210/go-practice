package main

func main(){

}

type ListNode struct {
	    Val int
	    Next *ListNode
	 }
   func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	   dummy := new(ListNode)
	   
	   for curr:=dummy;(list1 != nil && list2 != nil);{
		   if(list1.Val < list2.Val){
			curr.Next = &ListNode{Val:list1.Val}
			   list1 = list1.Next
		   }else{
			curr.Next = &ListNode{Val:list2.Val}
			   list2 = list2.Next
		   }
	   }
		   if(list1 != nil){
			   dummy.Next = list1
		   }
		   
		   if(list2 != nil){
			   dummy.Next = list2
		   }
	   return dummy
   }