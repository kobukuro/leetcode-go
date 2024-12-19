package main

import "fmt"

func climbStairsRecursion(n int, memo map[int]int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if value, exists := memo[n]; exists {
		return value
	}
	memo[n] = climbStairsRecursion(n-1, memo) + climbStairsRecursion(n-2, memo)
	return memo[n]
}
func climbStairs(n int) int {
	memo := map[int]int{}
	return climbStairsRecursion(n, memo)
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
