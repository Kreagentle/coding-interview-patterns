package fastandslowpointers

func findMaxAverage(nums []int, k int) float64 {
	var currentsum, maxsum int
	for i := 0; i < k; i++ {
		currentsum += nums[i]
	}
	maxsum = currentsum
	for i, j := 0, k; j < len(nums); i, j = i+1, j+1 {
		currentsum = currentsum - nums[i] + nums[j]
		if currentsum > maxsum {
			maxsum = currentsum
		}
	}
	return float64(maxsum) / float64(k)
}
