package twopointers

func minMovesToMakePalindrome(s string) int {
	runes := []rune(s)
	var moves int
	for i, j := 0, len(runes)-1; i < j; i++ {
		tmp := j
		for tmp > i {
			if runes[tmp] == runes[i] {
				for tmp < j {
					runes[tmp], runes[tmp+1] = runes[tmp+1], runes[tmp]
					moves += 1
					tmp += 1
				}
				j -= 1
				break
			}
			tmp -= 1
		}
		if tmp == i {
			moves += len(runes)/2 - i
		}
	}
	// Replace this placeholder return statement with your code
	return moves
}
