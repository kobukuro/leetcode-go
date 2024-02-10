package main

import "fmt"

func fibRecursive(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if value, exists := memo[n]; exists {
		return value
	}
	memo[n] = fibRecursive(n-1, memo) + fibRecursive(n-2, memo)
	return memo[n]
}
func fib(n int) int {
	memo := make(map[int]int)
	return fibRecursive(n, memo)
}
func main() {
	n := 3
	res := fib(n)
	fmt.Println(res) // 2
}
