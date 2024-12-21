package main

import (
	"fmt"
	"strconv"
)

func dfs(r int, c int, grid [][]byte, visited map[string]bool) bool {
	if r < 0 || r == len(grid) || c < 0 || c == len(grid[0]) || grid[r][c] == '0' {
		return false
	}
	key := strconv.Itoa(r) + ";" + strconv.Itoa(c)
	val := visited[key]
	if val {
		return false
	}
	visited[key] = true
	dfs(r+1, c, grid, visited)
	dfs(r-1, c, grid, visited)
	dfs(r, c+1, grid, visited)
	dfs(r, c-1, grid, visited)
	return true
}
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[string]bool)
	res := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, grid, visited) {
				res++
			}
		}
	}
	return res
}
func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid)) // 1
}
