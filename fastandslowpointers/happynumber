package fastandslowpointers

func isHappy(num int) bool {
	var slow, fast int
	slow = num
	fast = squareDigits(num)
  for fast != 1 && slow != fast {
    slow = squareDigits(slow)
    fast = squareDigits(squareDigits(fast))
  }
  if fast == 1 {
    return true
  }
	return false
}

func squareDigits(num int) int {
  result := 0
  for num != 0 {
    result += (num % 10) * (num % 10)
    num /= 10
  }
  return result
}
