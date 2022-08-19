package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	quickSort(nums1, 0, len(nums1)-1)
	quickSort(nums2, 0, len(nums2)-1)
	result := make([]int, 0, len(nums2))
	prev := 0
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			if len(result) == 0 {
				result = append(result, nums1[i])
				prev = nums1[i]
			} else {
				if prev != nums1[i] {
					result = append(result, nums1[i])
					prev = nums1[i]
				}
			}
			i++
			j++
		}
	}
	return result
}

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

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
