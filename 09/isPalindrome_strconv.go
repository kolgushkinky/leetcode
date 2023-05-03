import "strconv"

// isPalindrome returns true if integer is a palindrome in decimal notation.
// This function converts int to a slice of bytes.
func isPalindrome(x int) bool {
    a := []byte(strconv.Itoa(x))
    l := len(a)
    for i := 0; i <= l / 2; i++ {
        if a[i] != a[l - 1 - i] {
            return false
        }
    }
    return true
}
