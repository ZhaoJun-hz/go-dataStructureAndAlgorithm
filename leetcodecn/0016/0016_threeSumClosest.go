package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt
	resultTotal := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k > j+1 && nums[i]+nums[j]+nums[k-1] >= target {
				k--
			}
			total := nums[i] + nums[j] + nums[k]
			if int(math.Abs(float64(total-target))) < result {
				result = int(math.Abs(float64(total - target)))
				resultTotal = total
			}
			if k > j+1 {
				total = nums[i] + nums[j] + nums[k-1]
				if int(math.Abs(float64(total-target))) < result {
					result = int(math.Abs(float64(total - target)))
					resultTotal = total
				}
			}
		}
	}
	return resultTotal
}

func main() {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	//fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}
