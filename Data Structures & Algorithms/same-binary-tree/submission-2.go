/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	q1 := []*TreeNode{p}
	q2 := []*TreeNode{q}

	for len(q1) > 0 && len(q2) > 0 {
		for i := len(q1); i > 0; i-- {
			nodeP := q1[0]
			q1 = q1[1:]
			nodeQ := q2[0]
			q2 = q2[1:]

			if nodeP == nil && nodeQ == nil {
				continue
			}
			if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
				return false
			}

			q1 = append(q1, nodeP.Left, nodeP.Right)
			q2 = append(q2, nodeQ.Left, nodeQ.Right)
		}
	}

	return len(q1) == 0 && len(q2) == 0
}
