package hashmaps

import "strings"

func customSortString(order string, s string) string {
	mp := make(map[rune]int)

	for _, letter := range s {
		mp[letter]++
	}

	var output strings.Builder
	for _, value := range order {
		if el, ok := mp[value]; ok {
			output.WriteString(strings.Repeat(string(value), el))
			mp[value] = 0
		}
	}

	for _, value := range s {
		if el, ok := mp[value]; ok && el != 0 {
			output.WriteString(strings.Repeat(string(value), el))
			mp[value] = 0
		}
	}

	return output.String()
}
