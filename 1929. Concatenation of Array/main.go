package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
func main() {
	nums := []int{1, 2, 1}
	result := getConcatenation(nums)
	fmt.Println(result) // [1 2 1 1 2 1]
}
