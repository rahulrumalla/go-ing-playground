package palindrome

// HANNAH, NITIN
func IsPalindrome(input string) bool {
	strLength := len(input)

	var midPoint int
	midPoint = strLength / 2 // should work for odd & even

	for i := 0; i < midPoint; i++ {
		if input[i] != input[strLength-i-1] {
			return false
		}
	}
	return true
}
