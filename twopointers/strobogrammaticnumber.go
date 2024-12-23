package twopointers

func isStrobogrammatic(num string) bool {
	var start, end int
	end = len(num) - 1
	for start <= end {
		if (num[start] == 54 && num[end] == 57) || (num[start] == 57 && num[end] == 54) ||
			(num[start] == 56 && num[end] == 56) || (num[start] == 49 && num[end] == 49) ||
			(num[start] == 48 && num[end] == 48) {
			start += 1
			end -= 1
		} else {
			return false
		}
	}
	return true
}
