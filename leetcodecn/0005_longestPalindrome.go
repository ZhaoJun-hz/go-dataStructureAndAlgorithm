package main

import "fmt"

func longestPalindrome(s string) string {
	result := ""
	maxLength := 0
	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		tempLength := r - l + 1
		if tempLength > maxLength {
			maxLength = tempLength
			result = s[l+1 : r]
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		tempLength = r - l + 1
		if tempLength > maxLength {
			maxLength = tempLength
			result = s[l+1 : r]
		}
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
