package main

import "fmt"

func removeVowels(s string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	res := ""
	for _, c := range s {
		if !vowels[c] {
			res += string(c)
		}
	}
	return res
}
func main() {
	s := "leetcodeisacommunityforcoders"
	res := removeVowels(s)
	fmt.Println(res) // ltcdscmmntyfrcdrs
}
