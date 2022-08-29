package main

import "fmt"

func shuffle(nums []int, n int) []int {
	result := make([]int, 0, 2*n)
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		result = append(result, nums[i])
		result = append(result, nums[j])
	}
	return result
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}
