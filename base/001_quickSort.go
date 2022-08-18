package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ints[i])
	}

	quickSort(ints, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", ints[i])
	}

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
