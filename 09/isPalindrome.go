// isPalindrome returns true if integer is a palindrome in decimal notation
func isPalindrome(x int) bool {
	// negative numbers and those ending with zero
	// are NOT palindromes
	if x < 0 || x % 10 == 0 && x != 0 {
		return false
	}
	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if x == reversed || x == reversed/10 {
		return true
	} else {
		return false
	}
}
