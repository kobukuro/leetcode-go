package main

import "fmt"

func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, c := range s {
		if val, ok := pairs[c]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
func main() {
	s := "()"
	fmt.Println(isValid(s)) // true
}
