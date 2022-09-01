package main

import "fmt"

/*
给定两个有序的数组，找中位数(n + m) / 2，等价于找第k小的元素，k = (n + m) / 2
1、当一共有偶数个数时，找到第total / 2小left和第total / 2 + 1小right，结果是(left + right / 2.0)
2、当一共有奇数个数时，找到第total / 2 + 1小，即为结果


*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 0 {
		left := find(nums1, 0, nums2, 0, total/2)
		right := find(nums1, 0, nums2, 0, total/2+1)
		return (float64(left) + float64(right)) / 2.0
	} else {
		return float64(find(nums1, 0, nums2, 0, total/2+1))
	}
}

func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}
	if k == 1 {
		if len(nums1) == i {
			return nums2[j]
		} else {
			if nums1[i] >= nums2[j] {
				return nums2[j]
			} else {
				return nums1[i]
			}
		}
	}
	if len(nums1) == i {
		return nums2[j+k-1]
	}
	si := min(len(nums1), i+k/2)
	sj := j + k - k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))

	}
}

func min(nums1 int, nums2 int) int {
	if nums1 < nums2 {
		return nums1
	} else {
		return nums2
	}
}

func main() {
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Printf("%.5f\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
