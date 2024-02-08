package main

import "fmt"

func climbStairsRecursive(n int, memo map[int]int) int {
	if value, exists := memo[n]; exists {
		return value
	}
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	memo[n] = climbStairsRecursive(n-1, memo) + climbStairsRecursive(n-2, memo)
	return memo[n]
}
func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbStairsRecursive(n, memo)
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
