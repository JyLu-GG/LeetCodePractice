package leetcodepractice

import "fmt"

func TwoSum1(nums []int, target int) []int {
	index := 0
	otherIndex := -1
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			} else {
				if nums[j] == value {
					otherIndex = j
					break
				}
			}
		}
		if otherIndex == -1 {
			index += 1
		} else {
			break
		}
	}
	ans := []int{index, otherIndex}
	if otherIndex == -1 {
		return nil
	} else {
		return ans
	}
}

func TwoSum2(nums []int, target int) []int {
	var dic = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := dic[target-nums[i]]; ok {
			fmt.Println(ok)
			fmt.Println(v)
			fmt.Println(dic[v])
			fmt.Println(dic[target-nums[i]])
			return []int{v, i}
		} else {
			dic[nums[i]] = i
			fmt.Println("check fail")
			fmt.Printf("%d, %d\n", nums[i], dic[nums[i]])
		}
	}
	return nil
}
