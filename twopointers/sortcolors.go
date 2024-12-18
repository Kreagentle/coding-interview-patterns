package twopointers

func sortColors(colors []int) []int {
	var start, end, current int
	end = len(colors) - 1
	for current <= end {
		if colors[current] == 0 {
			temp := colors[start]
			colors[start] = colors[current]
			colors[current] = temp
			start++
			current++
		} else if colors[current] == 2 {
			temp := colors[end]
			colors[end] = colors[current]
			colors[current] = temp
			// current++
			end--
		} else {
			current++
		}
	}
	return colors
}
