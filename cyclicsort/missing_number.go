package cyclicsort

func findMissingNumber(nums []int) int {
	var ite int
	for ite < len(nums) {
		if nums[ite] == ite || nums[ite] >= len(nums) {
			ite++
		} else {
			corrIndex := nums[ite]
			nums[ite], nums[corrIndex] = nums[corrIndex], nums[ite]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
