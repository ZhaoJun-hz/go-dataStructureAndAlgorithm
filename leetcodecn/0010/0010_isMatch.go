package main

import "fmt"

// dp[i][j], 所有s[1:i]和p[1:j]的匹配方案, 是否存在匹配方案(bool)
// dp[i][j] = p[j] 不是 ?*  s[i] == p[j] || p[j] == . && dp[i - 1][j - 1]
//            p[j] 是 ?*
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	s = " " + s
	p = " " + p
	dp := [25][35]bool{}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i > 0 && p[j] != '*' {
				dp[i][j] = (s[i] == p[j] || p[j] == '.') && dp[i-1][j-1]
			} else if p[j] == '*' {
				dp[i][j] = dp[i][j-2] || i > 0 && dp[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[n][m]
}

func main() {
	//fmt.Println(isMatch("ab", ".*"))
	//fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", "a"))
}
