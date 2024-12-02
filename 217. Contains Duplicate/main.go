package main

import "fmt"

func containsDuplicate(nums []int) bool {
	lookup := map[int]bool{}
	for _, num := range nums {
		if _, ok := lookup[num]; ok {
			return true
		}
		lookup[num] = true
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println(result) // true
}
