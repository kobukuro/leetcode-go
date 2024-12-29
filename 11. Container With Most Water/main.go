package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // 49
}
