package twopointers

func validWordAbbreviation(word string, abbr string) bool {
	if len(word) < len(abbr) {
		return false
	}
	var first, second int
	for first < len(word) && second < len(abbr) {
		if isDigit(abbr[second]) {
			if abbr[second] == '0' {
				return false
			}
			var temp int
			for second < len(abbr) && isDigit(abbr[second]) {
				temp = temp*10 + int(abbr[second]-'0')
				second += 1
			}
			first += temp
		} else if abbr[second] != word[first] {
			return false
		} else {
			first += 1
			second += 1
		}
	}
	return first == len(word) && second == len(abbr)
}

func isDigit(el byte) bool {
	if el >= '0' && el <= '9' {
		return true
	}
	return false
}
