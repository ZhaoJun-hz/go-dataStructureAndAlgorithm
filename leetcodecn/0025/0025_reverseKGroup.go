package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{Val: -1}
	result.Next = head
	for p := result; ; {
		q := p
		for i := 0; i < k && q != nil; i++ {
			q = q.Next
		}
		if q == nil {
			break
		}
		a, b := p.Next, p.Next.Next
		for i := 0; i < k-1; i++ {
			c := b.Next
			b.Next = a
			a, b = b, c
		}
		c := p.Next
		p.Next = a
		c.Next = b
		p = c
	}

	return result.Next
}
