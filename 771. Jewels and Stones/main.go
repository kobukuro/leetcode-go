package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	lookup := make(map[rune]bool)
	res := 0
	for _, jewel := range jewels {
		if _, exists := lookup[jewel]; !exists {
			lookup[jewel] = true
		}
	}
	for _, stone := range stones {
		if _, exists := lookup[stone]; exists {
			res++
		}
	}
	return res
}
func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}
