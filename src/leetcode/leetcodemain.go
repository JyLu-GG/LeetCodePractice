package main

import (
	"fmt"

	"example.com/src/leetcode/leetcodepractice"
)

func main() {
	fmt.Println("Leet Code Practice")
	fmt.Println("")
	// fmt.Println("LeetCode Practice TwoSum")
	// input := []int{3, 2, 4}
	// target := 6
	// ans1 := leetcodepractice.TwoSum1(input, target)
	// ans2 := leetcodepractice.TwoSum2(input, target)
	// fmt.Println(ans1)
	// fmt.Println(ans2)
	// fmt.Println("")

	// fmt.Println("LeetCode Practice ContainerWithMostWater")
	// height := []int{2, 3, 4, 5, 18, 17, 6}
	// ans3 := leetcodepractice.MaxArea1(height)
	// ans4 := leetcodepractice.MaxArea2(height)
	// ans5 := leetcodepractice.MaxAreaFail(height)
	// fmt.Println(ans3)
	// fmt.Println(ans4)
	// fmt.Println(ans5)
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Merge Two Sorted Lists")
	// list1 := []int{1, 2, 4}
	// node1 := leetcodepractice.ArrayToListNode(list1)
	// list2 := []int{1, 3, 4}
	// node2 := leetcodepractice.ArrayToListNode(list2)
	// ans6 := leetcodepractice.MergeTwoLists(node1, node2)
	// for ans6 != nil {
	// 	fmt.Printf("Ans6 val: %d\n", ans6.Val)
	// 	ans6 = ans6.Next
	// }
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Add two numbers")
	// list3 := []int{9, 9, 9, 9, 9, 9, 9}
	// node3 := leetcodepractice.ArrayToListNode(list3)
	// list4 := []int{9, 9, 9, 9}
	// node4 := leetcodepractice.ArrayToListNode(list4)
	// ans7 := leetcodepractice.AddTwoNumbers(node3, node4)
	// for ans7 != nil {
	// 	fmt.Printf("Ans7 val: %d\n", ans7.Val)
	// 	ans7 = ans7.Next
	// }
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Valid Parentheses")
	// s := "([)]"
	// // s := "{[]}"
	// // s := "()[]{}"
	// ans8 := leetcodepractice.IsValid(s)
	// fmt.Printf("%t\n", ans8)
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Design Circular Queue")
	// k := 4
	// value := 999
	// obj := leetcodepractice.Constructor(k)
	// fmt.Printf("lens: %d\n", len(obj.Q))
	// fmt.Printf("cap: %d\n", cap(obj.Q))
	// param_1 := obj.EnQueue(value)
	// param_3 := obj.Front()
	// param_4 := obj.Rear()
	// param_5 := obj.IsEmpty()
	// param_6 := obj.IsFull()
	// param_2 := obj.DeQueue()
	// fmt.Printf("en: %t, de: %t, front: %d, rear: %d, empty: %t, full: %t\n", param_1, param_2, param_3, param_4, param_5, param_6)
	// obj.EnQueue(1)
	// obj.EnQueue(2)
	// obj.EnQueue(3)
	// param_1 = obj.EnQueue(4)
	// param_3 = obj.Front()
	// param_4 = obj.Rear()
	// param_5 = obj.IsEmpty()
	// param_6 = obj.IsFull()
	// fmt.Printf("en: %t, de: %t, front: %d, rear: %d, empty: %t, full: %t\n", param_1, param_2, param_3, param_4, param_5, param_6)
	// fmt.Printf("lens: %d\n", len(obj.Q))
	// fmt.Printf("cap: %d\n", cap(obj.Q))
	// param_1 = obj.EnQueue(5)
	// param_3 = obj.Front()
	// param_4 = obj.Rear()
	// param_5 = obj.IsEmpty()
	// param_6 = obj.IsFull()
	// fmt.Printf("en: %t, de: %t, front: %d, rear: %d, empty: %t, full: %t\n", param_1, param_2, param_3, param_4, param_5, param_6)
	// param_2 = obj.DeQueue()
	// param_3 = obj.Front()
	// param_4 = obj.Rear()
	// param_5 = obj.IsEmpty()
	// param_6 = obj.IsFull()
	// fmt.Printf("en: %t, de: %t, front: %d, rear: %d, empty: %t, full: %t\n", param_1, param_2, param_3, param_4, param_5, param_6)
	// obj.DeQueue()
	// obj.DeQueue()
	// param_2 = obj.DeQueue()
	// param_3 = obj.Front()
	// param_4 = obj.Rear()
	// param_5 = obj.IsEmpty()
	// param_6 = obj.IsFull()
	// fmt.Printf("en: %t, de: %t, front: %d, rear: %d, empty: %t, full: %t\n", param_1, param_2, param_3, param_4, param_5, param_6)
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Binary Tree Inorder Traversal")
	// list5 := []int{1, leetcodepractice.NULL, 2, 3}
	// treenode1 := leetcodepractice.IntArrayToTreeNode(list5)
	// ans9 := leetcodepractice.InorderTraversal(treenode1)
	// for _, v := range ans9 {
	// 	fmt.Printf("%d", v)
	// }
	// fmt.Println("")
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Flatten Binary Tree to Linked List")
	// list6 := []int{1, 2, 5, 3, 4, leetcodepractice.NULL, 6}
	// treenode2 := leetcodepractice.IntArrayToTreeNode(list6)
	// leetcodepractice.Flatten(treenode2)
	// tmp := treenode2
	// for tmp.Right != nil {
	// 	fmt.Println(tmp.Right)
	// 	tmp = tmp.Right
	// }
	// treenode2 = leetcodepractice.IntArrayToTreeNode(list6)
	// leetcodepractice.Flatten2(treenode2)
	// tmp = treenode2
	// for tmp.Right != nil {
	// 	fmt.Println(tmp.Right)
	// 	tmp = tmp.Right
	// }
	// treenode2 = leetcodepractice.IntArrayToTreeNode(list6)
	// leetcodepractice.Flatten3(treenode2)
	// tmp = treenode2
	// for tmp.Right != nil {
	// 	fmt.Println(tmp.Right)
	// 	tmp = tmp.Right
	// }
	// fmt.Println("")
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Roman to Integer")
	// var str string = "III"
	// var ans10 = leetcodepractice.RomanToInt(str)
	// fmt.Printf("ans10: %d\n", ans10)
	// str = "LVIII"
	// ans10 = leetcodepractice.RomanToInt(str)
	// fmt.Printf("ans10: %d\n", ans10)
	// str = "MCMXCIV"
	// ans10 = leetcodepractice.RomanToInt(str)
	// fmt.Printf("ans10: %d\n", ans10)

	// str = "III"
	// ans10 = leetcodepractice.RomanToInt2(str)
	// fmt.Printf("ans10: %d\n", ans10)
	// str = "LVIII"
	// ans10 = leetcodepractice.RomanToInt2(str)
	// fmt.Printf("ans10: %d\n", ans10)
	// str = "MCMXCIV"
	// ans10 = leetcodepractice.RomanToInt2(str)
	// fmt.Printf("ans10: %d\n", ans10)
	// fmt.Println("")

	// fmt.Println("LeetCode Practice Longest Substring Without Repeating Characters")
	// var str1 string = "abcabcbb"
	// var ans11 = leetcodepractice.LengthOfLongestSubstring(str1)
	// fmt.Printf("ans11: %d\n", ans11)
	// str1 = "bbbbb"
	// ans11 = leetcodepractice.LengthOfLongestSubstring(str1)
	// fmt.Printf("ans11: %d\n", ans11)
	// str1 = "pwwkew"
	// ans11 = leetcodepractice.LengthOfLongestSubstring(str1)
	// fmt.Printf("ans11: %d\n", ans11)

	// fmt.Println("")

	// fmt.Println("LeetCode Practice Add Binary")
	// var strA string = "11"
	// var strB string = "1"
	// var ans12 = leetcodepractice.AddBinary(strA, strB)
	// fmt.Printf("ans12: %s\n", ans12)

	// strA = "1010"
	// strB = "1011"
	// ans12 = leetcodepractice.AddBinary(strA, strB)
	// fmt.Printf("ans12: %s\n", ans12)

	// fmt.Println("")

	// fmt.Println("LeetCode Practice Divide Two Integer")
	// var ans13 = leetcodepractice.Divide(10, 5)
	// fmt.Printf("ans13: %d\n", ans13)

	// ans13 = leetcodepractice.Divide(-10, 5)
	// fmt.Printf("ans13: %d\n", ans13)

	// ans13 = leetcodepractice.Divide(-1, -1)
	// fmt.Printf("ans13: %d\n", ans13)

	// ans13 = leetcodepractice.Divide(2147483647, 3)
	// fmt.Printf("ans13: %d\n", ans13)

	// ans13 = leetcodepractice.Divide2(2147483647, 3)
	// fmt.Printf("ans13: %d\n", ans13)

	// fmt.Println("")

	// fmt.Println("LeetCode Practice Palindrome Number")
	// var ans14 = leetcodepractice.IsPalindrome(121)
	// fmt.Printf("ans14: %t\n", ans14)

	// ans14 = leetcodepractice.IsPalindrome(-121)
	// fmt.Printf("ans14: %t\n", ans14)

	// ans14 = leetcodepractice.IsPalindrome(10)
	// fmt.Printf("ans14: %t\n", ans14)

	// ans14 = leetcodepractice.IsPalindrome(123)
	// fmt.Printf("ans14: %t\n", ans14)

	// fmt.Println("")

	fmt.Println("LeetCode Practice Gray code")
	var ans15 = leetcodepractice.GrayCode(2)
	fmt.Printf("ans15: %d\n", ans15)

	ans15 = leetcodepractice.GrayCode(1)
	fmt.Printf("ans15: %d\n", ans15)

	fmt.Println("")

	fmt.Println("LeetCode Practice Remove Linked List Elements")
	list5 := []int{1, 2, 6, 3, 4, 5, 6}
	node5 := leetcodepractice.ArrayToListNode(list5)
	val := 6
	var ans16 = leetcodepractice.RemoveElements(node5, val)
	for ans16 != nil {
		fmt.Printf("Ans16 val: %d\n", ans16.Val)
		ans16 = ans16.Next
	}

	fmt.Println("")
}
