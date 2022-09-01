package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{-1, nil}
	result.Next = head
	for temp := result; temp.Next != nil && temp.Next.Next != nil; {
		a, b := temp.Next, temp.Next.Next
		temp.Next = b
		a.Next = b.Next
		b.Next = a
		temp = a
	}

	return result.Next
}

func main() {
	swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
}
