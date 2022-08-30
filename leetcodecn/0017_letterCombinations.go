package main

import "fmt"

var strs = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
var ans = []string{}

func letterCombinations(digits string) []string {
	ans = make([]string, 0, 10000)
	if len(digits) == 0 {
		return ans
	}
	dfs(digits, 0, "")
	return ans
}

func dfs(digits string, u int, path string) {
	if u == len(digits) {
		ans = append(ans, path)
	} else {
		index := int(digits[u] - '0')
		for _, element := range strs[index] {
			dfs(digits, u+1, path+string(element))
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
