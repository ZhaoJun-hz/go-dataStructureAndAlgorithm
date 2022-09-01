package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	reps := []string{"M",
		"CM", "D", "CD", "C",
		"XC", "L", "XL", "X",
		"IX", "V", "IV", "I"}
	values := []int{1000,
		900, 500, 400, 100,
		90, 50, 40, 10,
		9, 5, 4, 1}
	builder := strings.Builder{}
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			builder.WriteString(reps[i])
			num -= values[i]
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
