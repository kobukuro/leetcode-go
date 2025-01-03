package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}
func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums)) // 1
}
