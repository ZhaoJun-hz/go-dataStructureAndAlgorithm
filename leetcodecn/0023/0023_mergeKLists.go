package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Item struct {
	value    interface{}
	priority int
}

type ItemList []*Item

func (lst ItemList) Len() int {
	return len(lst)
}

func (lst ItemList) Less(i, j int) bool {
	return lst[i].priority < lst[j].priority
}

func (lst ItemList) Swap(i, j int) {
	lst[i], lst[j] = lst[j], lst[i]
}

func (lst *ItemList) Push(val interface{}) {
	item := val.(*Item)
	*lst = append(*lst, item)
}

func (lst *ItemList) Pop() interface{} {
	old := *lst
	n := len(old)
	item := old[n-1]
	*lst = old[0 : n-1]
	return item
}

type PQueue struct {
	pq ItemList
}

func NewPQueue() *PQueue {
	queue := new(PQueue)
	queue.pq = make(ItemList, 0)
	heap.Init(&queue.pq)
	return queue
}

func (queue *PQueue) Init() {
	queue.pq = make(ItemList, 0)
	heap.Init(&queue.pq)
}

func (queue *PQueue) Add(value interface{}, priority int) {
	heap.Push(&queue.pq, &Item{value: value, priority: priority})
}

func (queue *PQueue) Remove() interface{} {
	return heap.Pop(&queue.pq).(*Item).value
}

func (queue *PQueue) Len() int {
	return queue.pq.Len()
}

func (queue *PQueue) IsEmpty() bool {
	return queue.pq.Len() == 0
}

func mergeKLists(lists []*ListNode) *ListNode {
	queue := NewPQueue()
	result := &ListNode{Val: -1}
	tail := result
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			queue.Add(lists[i], lists[i].Val)
		}
	}
	for queue.IsEmpty() == false {
		node := queue.Remove().(*ListNode)
		tail.Next = node
		tail = tail.Next
		if node.Next != nil {
			queue.Add(node.Next, node.Next.Val)
		}
	}

	return result.Next
}

func main() {
	queue := NewPQueue()
	queue.Add("banana", 2)
	queue.Add("apple", 1)
	queue.Add("orange", 4)
	queue.Add("mango", 3)
	fmt.Println(queue.Len())
	for queue.IsEmpty() == false {
		fmt.Println(queue.Remove())
	}
}
