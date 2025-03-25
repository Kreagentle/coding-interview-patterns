package hashmaps

func uniqueOccurrences(arr []int) bool {
	mp := make(map[int]int)
	result := make(map[int]int)

	for _, el := range arr {
		mp[el]++
	}

	for _, value := range mp {
		result[value]++
	}

	for _, value := range result {
		if value != 1 {
			return false
		}
	}

	return true
}
