package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lookup := map[int]int{}
	for idx, num := range nums {
		complement := target - num
		if complementIndex, exists := lookup[complement]; exists {
			return []int{idx, complementIndex}
		} else {
			lookup[num] = idx
		}
	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // [1 0]
}
