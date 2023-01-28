package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}

	for nonEmpty(lists) { //listLoop

	}
	return res
}

func nonEmpty(lists []*ListNode) bool {
	for _, l := range lists {
		if l != nil {
			return true
		}
	}

	return false
}
