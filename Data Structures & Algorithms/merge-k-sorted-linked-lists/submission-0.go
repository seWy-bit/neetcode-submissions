/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if i + 1 < len(lists) {
				l2 = lists[i+1]
			}
			mergedLists = append(mergedLists, mergeList(l1, l2))
		}
		lists = mergedLists
	}

	return lists[0]
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	node := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	node.Next = list1
	if list1 == nil {
		node.Next = list2
	}

	return dummy.Next
}

