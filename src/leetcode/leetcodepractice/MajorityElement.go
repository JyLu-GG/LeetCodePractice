package leetcodepractice

import "sort"

func MajorityElement(nums []int) int {
	elementmap := make(map[int]int)
	max, ret := 0, 0
	for _, val := range nums {
		elementmap[val]++
		if elementmap[val] > max {
			max = elementmap[val]
			ret = val
		}
	}
	return ret
}

func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func MajorityElement3(nums []int) int {
	ans := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			ans = nums[i]
		}
		if nums[i] == ans {
			count++
		}
		if nums[i] != ans {
			count--
		}

	}
	return ans
}
