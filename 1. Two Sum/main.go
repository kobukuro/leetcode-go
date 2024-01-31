package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, exists := lookup[target-nums[i]]; exists {
			return []int{i, value}
		} else {
			lookup[nums[i]] = i
		}
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
