package twopointers

func isPalindrome(inputString string) bool {
	var a, b byte
	for i := 0; i < len(inputString)/2; i++ {
		a = inputString[i]
		b = inputString[len(inputString)-1-i]
		if a != b {
			return false
		}
	}
	return true
}
