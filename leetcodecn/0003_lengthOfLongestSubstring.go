package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	temp := make(map[byte]int)
	maxLength := 0
	for i, j := 0, 0; i < len(s); i++ {
		if value, exist := temp[s[i]]; exist {
			temp[s[i]] = value + 1
		} else {
			temp[s[i]] = 1
		}

		for temp[s[i]] > 1 {
			if s[j] == s[i] {
				temp[s[i]] = temp[s[i]] - 1
			} else {
				delete(temp, s[j])
			}
			j++
		}
		length := i - j + 1
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
