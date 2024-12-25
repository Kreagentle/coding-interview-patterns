package twopointers

func isPalindromeii(str string) bool {
	var left, right int
	right = len(str) - 1
	for left <= right {
		if str[left] != str[right] {
			return isPal(str[left+1:right+1]) || isPal(str[left:right])
		}
		right -= 1
		left += 1
	}
	return true
}

func isPal(str string) bool {
	var i, j int
	j = len(str) - 1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		j -= 1
		i += 1
	}
	return true
}
