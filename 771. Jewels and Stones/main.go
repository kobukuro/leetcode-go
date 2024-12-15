package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	lookup := map[rune]bool{}
	res := 0
	for _, jewel := range jewels {
		lookup[jewel] = true
	}
	for _, stone := range stones {
		if lookup[stone] {
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
