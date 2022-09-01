package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, len(nums)-1; k < l; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for l > k+1 && nums[i]+nums[j]+nums[k]+nums[l-1] >= target {
					l--
				}
				if nums[i]+nums[j]+nums[k]+nums[l] == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
