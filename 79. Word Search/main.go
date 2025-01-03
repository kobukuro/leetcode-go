package main

import (
	"fmt"
	"strconv"
)

func dfs(r int, c int, board [][]byte, word string, visited map[string]bool, index int) bool {
	key := strconv.Itoa(r) + ";" + strconv.Itoa(c)
	if index == len(word) {
		return true
	}
	if r < 0 || r == len(board) || c < 0 || c == len(board[0]) || visited[key] || board[r][c] != word[index] {
		return false
	}
	visited[key] = true
	if dfs(r+1, c, board, word, visited, index+1) || dfs(r-1, c, board, word, visited, index+1) || dfs(r, c+1, board, word, visited, index+1) || dfs(r, c-1, board, word, visited, index+1) {
		return true
	}
	visited[key] = false
	return false

}
func exist(board [][]byte, word string) bool {
	visited := make(map[string]bool)
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if dfs(r, c, board, word, visited, 0) {
				return true
			}
		}
	}
	return false
}
func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // true
}
