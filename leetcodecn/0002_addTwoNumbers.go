package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	current := &result
	temp := 0
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: temp % 10}
		temp = temp / 10
		current = current.Next
	}
	return result.Next
}

func main() {
	l1 := ListNode{9, &ListNode{9, nil}}
	l2 := ListNode{9, nil}
	addTwoNumbers(&l1, &l2)
}
