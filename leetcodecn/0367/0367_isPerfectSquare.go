package main

import "fmt"

func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l < r {
		mid := (l + r) >> 1
		if mid*mid >= num {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l*l == num {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(25))
	fmt.Println(isPerfectSquare(24))
	fmt.Println(isPerfectSquare(36))
}
