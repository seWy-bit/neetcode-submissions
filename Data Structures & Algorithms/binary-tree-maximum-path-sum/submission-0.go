/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := []int{root.Val}

	var dfs func(node *TreeNode) int 
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		res[0] = max(res[0], node.Val + rightMax + leftMax)

		return node.Val + max(leftMax, rightMax)
	}

	dfs(root)
	return res[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

