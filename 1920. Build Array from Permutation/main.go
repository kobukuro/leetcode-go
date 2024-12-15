package main

import "fmt"

func buildArray(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = nums[nums[i]]
	}
	return res
}
func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	result := buildArray(nums)
	fmt.Println(result) // [0 1 2 4 5 3]
}
