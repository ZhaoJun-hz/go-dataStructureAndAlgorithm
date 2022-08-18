package main

import "fmt"

// 1  2  3
func missingNumber(nums []int) int {
	length := len(nums)
	sum := (1 + length) * length / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
