/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	l, r := head, head

	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
		if l == r {
			return true
		}
	}

	return false
}
