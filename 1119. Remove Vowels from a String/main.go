package main

import "fmt"

func removeVowels(s string) string {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	res := ""
	for _, char := range s {
		if _, exists := vowels[string(char)]; !exists {
			res += string(char)
		}
	}
	return res
}
func main() {
	s := "leetcodeisacommunityforcoders"
	res := removeVowels(s)
	fmt.Println(res)
}
