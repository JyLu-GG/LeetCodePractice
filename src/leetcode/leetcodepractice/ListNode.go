package leetcodepractice

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListNode(data []int) *ListNode {
	var HeadNode *ListNode = &ListNode{Val: data[0]}
	tmp := HeadNode
	for i := 1; i < len(data); i++ {
		node := ListNode{Val: data[i]}
		tmp.Next = &node
		tmp = &node

	}
	return HeadNode
}
