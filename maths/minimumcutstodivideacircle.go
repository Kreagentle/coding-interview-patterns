package maths

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 != 0 {
		return n
	}
	return n / 2
}
