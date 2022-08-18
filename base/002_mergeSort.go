package main

import "fmt"

var ints = make([]int, 100010)
var temp = make([]int, 100010)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ints[i])
	}

	mergeSort(ints, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", ints[i])
	}
}

func mergeSort(q []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		j++
		k++
	}
	for i, k = l, 0; i <= r; i, k = i+1, k+1 {
		q[i] = temp[k]
	}
}
