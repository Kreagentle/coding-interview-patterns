package bitwise_manipulation

func extraCharacterIndex(str1 string, str2 string) int {
	var temp int
	for _, i := range str1 {
		temp ^= int(i)
	}
	for _, j := range str2 {
		temp ^= int(j)
	}
	if len(str1) > len(str2) {
		for counter, i := range str1 {
			if temp == int(i) {
				return counter
			}
		}
	} else {
		for counter, i := range str2 {
			if temp == int(i) {
				return counter
			}
		}
	}
	return 0
}
