package twopointers

import (
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			if i != left && i != right {
				result := nums[i] + nums[left] + nums[right]
				if result == target {
					return true
				} else if result > target {
					right -= 1
				} else {
					left += 1
				}
			}
		}
	}
	return false
}
