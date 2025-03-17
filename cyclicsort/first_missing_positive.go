package cyclicsort

func smallestMissingPositiveInteger(nums []int) int {
	ind := 0
	for ind < len(nums) {
		newInd := nums[ind] - 1
		if newInd == ind || newInd >= len(nums) || newInd < 0 || nums[ind] == nums[newInd] {
			ind++
		} else {
			nums[ind], nums[newInd] = nums[newInd], nums[ind]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums)
}
