package main

import "fmt"

// 双指针
// 每次都计算,更新下当前最大值
func maxArea(height []int) int {
	maxArea := -1
	for i, j := 0, len(height)-1; i < j; {
		minNumber := min(height[i], height[j])
		tempArea := minNumber * (j - i)
		if tempArea > maxArea {
			maxArea = tempArea
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
