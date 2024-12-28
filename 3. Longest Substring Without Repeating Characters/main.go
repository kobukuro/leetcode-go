package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	lookup := make(map[byte]bool)
	l := 0
	for r := 0; r < len(s); r++ {
		for lookup[s[r]] {
			lookup[s[l]] = false
			l++
		}
		lookup[s[r]] = true
		res = max(res, r-l+1)
	}
	return res
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // 3
}
