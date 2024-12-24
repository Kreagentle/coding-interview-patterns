package twopointers

func isPalindromeii(str string) bool {
	if len(str) < 3 {
		return true
	}
	var left, right int
	temp := false
	right = len(str) - 1
	if len(str) != 3 {
		for left <= right {
			if str[left] != str[right] {
				if temp == false {
					temp = true
					if str[left+1] == str[right] && str[left] != str[right-1] {
						left += 1
					} else {
						right -= 1
					}
				} else {
					return false
				}
			}
			right -= 1
			left += 1
		}
	} else {
		for left <= right {
			if str[left] != str[right] {
				return false
			}
		}
	}
	return true
}
