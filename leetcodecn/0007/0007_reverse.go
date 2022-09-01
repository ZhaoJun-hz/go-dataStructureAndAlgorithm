package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		if result > 0 && result > (math.MaxInt32-x%10)/10 {
			return 0
		}
		if result < 0 && result < (math.MinInt32-x%10)/10 {
			return 0
		}
		result = result*10 + x%10
		x /= 10
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}
