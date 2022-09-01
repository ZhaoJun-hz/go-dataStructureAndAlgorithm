package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) >> 1
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
}
