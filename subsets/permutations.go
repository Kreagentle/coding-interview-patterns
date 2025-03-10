package subsets

func permuteWord(word string) []string {
	var result []string

	var backtracking func(gen, newstring string)
	backtracking = func(gen, newstring string) {
		if len(gen) == len(word) {
			result = append(result, gen)
			return
		}
		for counter, i := range newstring {
			backtracking(gen+string(i), newstring[:counter]+newstring[counter+1:])
		}
	}
	backtracking("", word)
	return result
}
