package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	number := x
	for number > 0 {
		result = result*10 + number%10
		number /= 10
	}
	if result == x {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
