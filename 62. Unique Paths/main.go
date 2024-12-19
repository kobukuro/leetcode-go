package main

import (
	"fmt"
	"strconv"
)

func uniquePathsRecursive(m int, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + ";" + strconv.Itoa(n)
	if value, exists := memo[key]; exists {
		return value
	}
	if m == 1 || n == 1 {
		return 1
	}
	memo[key] = uniquePathsRecursive(m-1, n, memo) + uniquePathsRecursive(m, n-1, memo)
	return memo[key]
}
func uniquePaths(m int, n int) int {
	memo := make(map[string]int)
	return uniquePathsRecursive(m, n, memo)
}
func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n)) // 28
}
