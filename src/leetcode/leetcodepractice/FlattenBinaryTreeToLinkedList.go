package leetcodepractice

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil && root.Right != nil {
			t := root.Left
			for t.Right != nil {
				t = t.Right
			}
			t.Right = root.Right
		}

		if root.Left != nil {
			root.Right = root.Left
		}
		root.Left = nil
		root = root.Right
	}
}

func myorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	myorder(root.Left, res)
	myorder(root.Right, res)
}

func Flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	var res []int
	myorder(root, &res)

	for i := 0; i < len(res)-1; i++ {
		root.Val = res[i]
		root.Left = nil
		root.Right = &TreeNode{res[i+1], nil, nil}
		root = root.Right
	}
}

func myorder2(root *TreeNode) {
	if root == nil {
		return
	}
	res2 = append(res2, root)
	myorder2(root.Left)
	myorder2(root.Right)
}

var res2 []*TreeNode

func Flatten3(root *TreeNode) {
	if root == nil {
		return
	}

	myorder2(root)
	
	for i := len(res2) - 1; i > 0; i-- {
		res2[i-1].Left = nil
		res2[i-1].Right = res2[i]
	}

}
