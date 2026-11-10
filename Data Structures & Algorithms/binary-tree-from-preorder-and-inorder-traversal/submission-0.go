/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    idx := make(map[int]int)
	for i, val := range inorder {
		idx[val] = i
	}

	preIdx := 0

	var dfs func(int, int ) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: rootVal}
		mid := idx[rootVal]

		root.Left = dfs(left, mid - 1)
		root.Right = dfs(mid + 1, right)

		return root
	}

	return dfs(0, len(inorder) - 1)
}
