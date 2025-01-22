package main

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	string1 := ""
	for _, word := range word1 {
		string1 += word
	}
	string2 := ""
	for _, word := range word2 {
		string2 += word
	}
	return string1 == string2
}
func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2)) // true

}
