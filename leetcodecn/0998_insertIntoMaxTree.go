package main

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	result := root
	node := TreeNode{val, nil, nil}
	if root.Val < val {
		node.Left = root
		return &node
	}
	for root.Right != nil && root.Right.Val > val {
		root = root.Right
	}
	if root.Right == nil {
		root.Right = &node
	} else {
		node.Left = root.Right
		root.Right = &node
	}
	return result
}
