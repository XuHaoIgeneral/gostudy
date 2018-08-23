package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(node *ListNode) *ListNode {
	if node.Next==nil{
		return node
	}

	preNode:=node.Next
	nextNode:=node.Next.Next
	if nextNode!=nil{
		preNode.Next=node
		node.Next=nextNode
	}else {
		preNode.Next=node
		node.Next=nil
	}

	return reverseList(preNode)
}
