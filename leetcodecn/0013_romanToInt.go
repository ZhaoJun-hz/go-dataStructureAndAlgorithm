package main

import "fmt"

func romanToInt(s string) int {
	result := 0
	value := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && value[s[i]] < value[s[i+1]] {
			result -= value[s[i]]
		} else {
			result += value[s[i]]
		}
	}

	return result
}

func main() {
	//fmt.Println(romanToInt("III"))
	//fmt.Println(romanToInt("IV"))
	//fmt.Println(romanToInt("IX"))
	//fmt.Println(romanToInt("LVIII"))
	//fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("DCXXI"))
}
