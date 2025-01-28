package greedy

func canJump(nums []int) bool {
	init_index := len(nums) - 1
	for k := len(nums) - 1; k >= 0; k-- {
		if nums[k] >= init_index-k {
			init_index = k
		}
	}
	return init_index == 0
}
