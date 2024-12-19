package twopointers

func reverseWords(sentence string) string {
	var result string
	sentence = reverse(sentence)
	var begin, end int
	for end < len(sentence) {
		if sentence[end] == ' ' {
			end++
			begin++
		} else {
			for end <= len(sentence)-1 && sentence[end] != ' ' {
				end++
			}
			temp := reverse(sentence[begin:end])
			result += temp
			end++
			begin = end
			result += " "
		}
	}
	return result[:len(result)-1]
}

func reverse(sentence string) string {
	var result []byte
	end := len(sentence) - 1
	for end >= 0 {
		result = append(result, sentence[end])
		end--
	}
	return string(result)
}
