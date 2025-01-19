package main

import (
	"fmt"
	"strconv"
)

func explore(r int, c int, grid [][]int, visited map[string]bool) int {
	key := strconv.Itoa(r) + ";" + strconv.Itoa(c)
	if r < 0 || r == len(grid) || c < 0 || c == len(grid[0]) || grid[r][c] == 0 || visited[key] {
		return 0
	}
	visited[key] = true
	return 1 + explore(r+1, c, grid, visited) + explore(r-1, c, grid, visited) +
		explore(r, c+1, grid, visited) + explore(r, c-1, grid, visited)
}
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	rows := len(grid)
	cols := len(grid[0])
	visited := map[string]bool{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			res = max(res, explore(r, c, grid, visited))
		}
	}
	return res
}
func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid)) // 6
}
