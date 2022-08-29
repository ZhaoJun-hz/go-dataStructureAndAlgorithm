package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	builder := strings.Builder{}
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += 2 * (numRows - 1) {
				builder.WriteByte(s[j])
			}
		} else {
			for j, k := i, 2*(numRows-1)-i; j < len(s) || k < len(s); j, k = j+2*(numRows-1), k+2*(numRows-1) {
				if j < len(s) {
					builder.WriteByte(s[j])
				}
				if k < len(s) {
					builder.WriteByte(s[k])
				}
			}
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

// 4
// 0     6        12  2 * (n - 1)
// 1   5 7     11     2 * (n - 2)   2 * (n - 3)
// 2 4   8  10        2 * (n - 3)   2 * (n - 2)
// 3     9            2 * (n - 1)

// 6
// 0 0         10
// 1 1       9 11
// 2 2     8   12
// 3 3   7     13
// 4 4 6       14
// 5 5         15
