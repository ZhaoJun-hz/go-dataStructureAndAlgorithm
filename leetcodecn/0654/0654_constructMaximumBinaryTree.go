package main

// 假设数组中0 到 i - 1 已经构建成一棵满足结果的树,
// 需要将第i个数插入, 使该数的中序遍历是原序列, 那么这个数,在结果树的右链最下面的那个点
// 这里需要分情况讨论, 情况1, nums[i]是最大数, 这个数当成根节点就可以了, 之前的树成为左子树就可以了
// 情况2, nums[i]小于最大值, 顺着右子树找, 找到最后一个比nums[i]大的数, 在这个节点上插入nums[i], 之前的树成为新节点的左节点
func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := []*TreeNode{}
	for _, num := range nums {
		treeNode := TreeNode{Val: num}
		for len(stack) > 0 && stack[len(stack)-1].Val < num {
			treeNode.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			stack[len(stack)-1].Right = &treeNode
		}
		stack = append(stack, &treeNode)
	}
	for len(stack) > 1 {
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
