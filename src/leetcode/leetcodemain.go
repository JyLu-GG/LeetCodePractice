package main

import (
	"fmt"

	"example.com/src/leetcode/leetcodepractice"
)

func main() {
	fmt.Println("Leet Code Practice")
	fmt.Println("")
	fmt.Println("LeetCode Practice TwoSum")
	input := []int{3, 2, 4}
	target := 6
	ans1 := leetcodepractice.TwoSum1(input, target)
	ans2 := leetcodepractice.TwoSum2(input, target)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println("")

	fmt.Println("LeetCode Practice ContainerWithMostWater")
	height := []int{2, 3, 4, 5, 18, 17, 6}
	ans3 := leetcodepractice.MaxArea1(height)
	ans4 := leetcodepractice.MaxArea2(height)
	ans5 := leetcodepractice.MaxAreaFail(height)
	fmt.Println(ans3)
	fmt.Println(ans4)
	fmt.Println(ans5)
	fmt.Println("")

	fmt.Println("LeetCode Practice Merge Two Sorted Lists")
	list1 := []int{1, 2, 4}
	node1 := leetcodepractice.ArrayToListNode(list1)
	list2 := []int{1, 3, 4}
	node2 := leetcodepractice.ArrayToListNode(list2)
	ans6 := leetcodepractice.MergeTwoLists(node1, node2)
	for ans6 != nil {
		fmt.Printf("Ans6 val: %d\n", ans6.Val)
		ans6 = ans6.Next
	}
	fmt.Println("")

	fmt.Println("LeetCode Practice Add two numbers")
	list3 := []int{9, 9, 9, 9, 9, 9, 9}
	node3 := leetcodepractice.ArrayToListNode(list3)
	list4 := []int{9, 9, 9, 9}
	node4 := leetcodepractice.ArrayToListNode(list4)
	ans7 := leetcodepractice.AddTwoNumbers(node3, node4)
	for ans7 != nil {
		fmt.Printf("Ans7 val: %d\n", ans7.Val)
		ans7 = ans7.Next
	}
	fmt.Println("")

}
