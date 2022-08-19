package main

func guessNumber(n int) int {
	l, r := 1, n
	for l < r {
		mid := (l + r) >> 1
		if guess(mid) == -1 || guess(mid) == 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
