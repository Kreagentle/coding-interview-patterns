package bitwise_manipulation

func findBitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	counter, mask := num, 1
	for counter != 0 {
		num ^= mask
		mask = mask << 1
		counter = counter >> 1

	}
	return num
}
