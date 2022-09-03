package leetcodepractice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {
	retnode := &ListNode{Next: head}
	curnode := retnode
	for curnode.Next != nil && curnode.Next.Next != nil {
		tmp1 := curnode.Next
		tmp2 := curnode.Next.Next
		tmp1.Next = tmp2.Next
		curnode.Next = tmp2
		curnode.Next.Next = tmp1
		curnode = curnode.Next.Next
	}

	return retnode.Next
}
