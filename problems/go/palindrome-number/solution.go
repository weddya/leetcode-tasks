package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	origX := x
	var revX, revDigit int
	for x >= 1 {
		revDigit = x % 10
		revX = revX*10 + revDigit
		x = (x - revDigit) / 10
	}

	return origX == revX
}
