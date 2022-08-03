package leetcodepractice

// The treenode traversal have some rule,
// can see picture
// All_DFS_Traversals(Preorder_Inorder_Postorder)_TreeTraversal.png
// or https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/283746/All-DFS-traversals-(preorder-inorder-postorder)-in-Python-in-1-line

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InorderTraversal(root *TreeNode) []int {
	var ret = make([]int, 0)

	if root == nil {
		return ret
	}

	ret = append(ret, InorderTraversal(root.Left)...)

	ret = append(ret, root.Val)

	ret = append(ret, InorderTraversal(root.Right)...)

	return ret
}

func InorderTraversal2(root *TreeNode) []int {
	// This is interesting program, write function in function
	ans := []int{}

	var inorder func(*TreeNode)

	//// here
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)

	}
	inorder(root)

	return ans
}

func InorderTraversal3(root *TreeNode) []int {
	// This will use stack logic to do
	s := []*TreeNode{}
	ans := []int{}

	node := root
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		node = s[len(s)-1]
		s = s[:len(s)-1]
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans

}
