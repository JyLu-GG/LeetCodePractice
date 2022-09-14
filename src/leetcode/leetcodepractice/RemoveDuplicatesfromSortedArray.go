package leetcodepractice

import (
	"sort"
)

func RemoveDuplicates(nums []int) int {
	k := len(nums)
	var h int = nums[0]
	for i := 1; i < len(nums); i++ {
		if (i + 1) < len(nums) {
			if h == nums[i] {
				nums[i] = 99999
				k--
			} else {
				h = nums[i]
			}
		} else {
			if h == nums[i] {
				nums[i] = 99999
				k--
			}
			break
		}
	}
	sort.Ints(nums)

	return k
}

func RemoveDuplicates2(nums []int) int {
	// i := 0

	// for j := 1; j < len(nums); j++ {
	// 	if nums[i] != nums[j] {
	// 		i++
	// 		nums[i] = nums[j]
	// 	}
	// }

	// return i + 1
	x := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[x] = nums[i]
			x++
		}
	}
	return x
}
