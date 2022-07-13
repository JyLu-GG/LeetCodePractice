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
}
