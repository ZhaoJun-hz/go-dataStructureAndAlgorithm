package main

import "fmt"

var ans = []string{}

// 任意前缀中, 左括号数量大于等于右括号数量
// 左右括号数量相等
func generateParenthesis(n int) []string {
	ans = make([]string, 0, 100)
	dfs(n, n, "")
	return ans
}

func dfs(lc int, rc int, path string) {

	if lc == 0 && rc == 0 {
		ans = append(ans, path)
	} else {
		if lc < rc {
			dfs(lc, rc-1, path+")")
		}
		if lc > 0 {
			dfs(lc-1, rc, path+"(")
		}
	}
}

func main() {
	//fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
