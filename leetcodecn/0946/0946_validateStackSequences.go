package main

func validateStackSequences(pushed []int, popped []int) bool {
	ints := make([]int, 0, 1002)
	i, j := 0, 0
	for i < len(pushed) {
		ints = append(ints, pushed[i])
		i++
		for j < len(popped) && len(ints) > 0 && popped[j] == ints[len(ints)-1] {
			ints = ints[:len(ints)-1]
			j++
		}
	}
	return j == len(popped)
}
