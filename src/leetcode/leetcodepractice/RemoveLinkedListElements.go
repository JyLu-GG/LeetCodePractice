package leetcodepractice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveElements(head *ListNode, val int) *ListNode {
	retnode := &ListNode{}
	curnode := retnode
	for head != nil {
		if head.Val != val {
			curnode.Next = &ListNode{Val: head.Val}
			curnode = curnode.Next
		}
		head = head.Next

	}

	return retnode.Next
}

func RemoveElements2(head *ListNode, val int) *ListNode {
	retnode := &ListNode{Next: head}
	curnode := retnode
	for curnode.Next != nil {
		if curnode.Next.Val == val {
			curnode.Next = curnode.Next.Next
		} else {
			curnode = curnode.Next
		}
	}

	return retnode.Next
}
