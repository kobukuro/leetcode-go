package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	}
	if p2 != nil {
		curr.Next = p2
	}
	return dummy.Next
}
func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(list1, list2))
}
