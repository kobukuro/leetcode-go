package main

import "fmt"

func containsDuplicate(nums []int) bool {
	lookup := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, exists := lookup[nums[i]]; exists {
			return true
		} else {
			lookup[nums[i]] = true
		}
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println(result)
}
