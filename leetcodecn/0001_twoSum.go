package main

import "fmt"

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := temp[target-nums[i]]; ok {
			return []int{value, i}
		} else {
			temp[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
