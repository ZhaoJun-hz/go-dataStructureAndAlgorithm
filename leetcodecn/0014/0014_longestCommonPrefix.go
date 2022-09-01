package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	builder := strings.Builder{}
	for i := 0; i < len(strs[0]); i++ {
		if i < len(strs[0]) {
			temp := strs[0][i]
			for k := 0; k < len(strs); k++ {
				if i >= len(strs[k]) {
					return builder.String()
				}
				if strs[k][i] != temp {
					return builder.String()
				}
			}
			builder.WriteByte(temp)
		}
	}

	return builder.String()
}

func main() {
	//fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))

}
