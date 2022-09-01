package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l >= len(nums) {
		return -1
	}
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func main() {
	//fmt.Println(search([]int{-1,0,3,5,9,12}, 9)) // 4
	//fmt.Println(search([]int{-1,0,3,5,9,12}, 2)) // -1
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 13)) // -1
}
