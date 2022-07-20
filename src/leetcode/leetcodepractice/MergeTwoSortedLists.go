package leetcodepractice

/**
 * 10
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}
	var tmp = &ListNode{}
	var ret = tmp
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
			tmp = tmp.Next
		}
	}
	if list1 != nil {
		tmp.Next = list1
	} else if list2 != nil {
		tmp.Next = list2
	}
	return ret.Next
}
