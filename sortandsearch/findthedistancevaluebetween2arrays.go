package sortandsearch

import (
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	totalCount := 0
	sort.Ints(arr2)

	for _, num1 := range arr1 {
		low, high := 0, len(arr2)-1
		valid := true

		for low <= high {
			mid := low + (high-low)/2

			if arr2[mid] == num1 {
				valid = false
				break
			} else if arr2[mid] < num1 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if low < len(arr2) && int(math.Abs(float64(arr2[low]-num1))) <= d {
			valid = false
		}

		if high >= 0 && int(math.Abs(float64(arr2[high]-num1))) <= d {
			valid = false
		}

		if valid {
			totalCount++
		}
	}
	return totalCount
}
