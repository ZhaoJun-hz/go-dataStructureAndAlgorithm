package main

import "fmt"

func isValid(s string) bool {
	// 括号对字典
	bracketsMap := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	// 传入字符串为空则返回false（leetcode认为空字符串应该返回true，这里注意）
	if s == "" {
		return true
	}
	// 初始化栈
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		// 如果栈中有数据，则进行比对，如果栈顶元素和当前元素匹配，则弹出，其他情况向栈中压入元素
		if len(stack) > 0 {
			if c, ok := bracketsMap[stack[len(stack)-1]]; ok && c == s[i] {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	// 到最后如果栈不为空，则说明未完全匹配掉（完全匹配的话，栈只能为空）
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid(""))       // true
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
