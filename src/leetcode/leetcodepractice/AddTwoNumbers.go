package leetcodepractice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remainder int
	var quotient int
	headnode := &ListNode{}
	tmp := headnode
	for {
		var n1, n2 int
		if l1 == nil && l2 == nil {
			if quotient > 0 {
				node := ListNode{Val: quotient}
				tmp.Next = &node
				tmp = &node
			}
			break
		}
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		remainder = (n1 + n2 + quotient) % 10
		quotient = (n1 + n2 + quotient) / 10
		node := ListNode{Val: remainder}
		tmp.Next = &node
		tmp = &node

	}

	return headnode.Next
}
