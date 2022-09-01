package main

import "fmt"

func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}
	x := q[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func intersect(nums1 []int, nums2 []int) []int {
	quickSort(nums1, 0, len(nums1)-1)
	quickSort(nums2, 0, len(nums2)-1)
	result := make([]int, 0, len(nums2))
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
