package main

import "fmt"

// 1 2 3 4 5
// 1 3 6 10 15
func arrangeCoins(n int) int {
	l, r := 1, n
	for l < r {
		mid := (l + r) >> 1
		if mid*(mid+1)/2 >= n {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l*(l+1)/2 == n {
		return l
	} else {
		return l - 1
	}
}

//func arrangeCoins(n int) int {
//	temp := 0
//	tempCount := 1
//	for n >= tempCount {
//		n = n - tempCount
//		temp++
//		tempCount++
//	}
//	return temp
//}

func main() {
	fmt.Println(arrangeCoins(1))  // 1
	fmt.Println(arrangeCoins(2))  // 1
	fmt.Println(arrangeCoins(3))  // 2
	fmt.Println(arrangeCoins(4))  // 2
	fmt.Println(arrangeCoins(5))  // 2
	fmt.Println(arrangeCoins(6))  // 3
	fmt.Println(arrangeCoins(7))  // 3
	fmt.Println(arrangeCoins(8))  // 3
	fmt.Println(arrangeCoins(9))  // 3
	fmt.Println(arrangeCoins(10)) // 4
}
