package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}
	flag := 1
	if s[i] == '-' {
		flag = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	result := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		x := int(s[i] - '0')
		if flag > 0 && result > (math.MaxInt32-x)/10 {
			return math.MaxInt32
		}
		if flag < 0 && -result < (math.MinInt32+x)/10 {
			return math.MinInt32
		}
		if -result*10-x == math.MinInt32 {
			return math.MinInt32
		}
		result = result*10 + x
		i++
	}
	result *= flag
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}

func main() {
	//fmt.Println(myAtoi("42"))
	//fmt.Println(myAtoi("   -42"))
	//fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("21474836460"))
}
