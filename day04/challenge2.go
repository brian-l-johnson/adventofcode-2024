package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isX_MAS(board [][]string, x int, y int) bool {
	if strings.Contains("MS", board[x-1][y-1]) &&
		strings.Contains("MS", board[x+1][y-1]) &&
		strings.Contains("MS", board[x-1][y+1]) &&
		strings.Contains("MS", board[x+1][y+1]) {
		if board[x-1][y-1] != board[x+1][y+1] &&
			board[x-1][y+1] != board[x+1][y-1] {
			return true
		}
	}
	return false
}

func main() {
	//read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var board [][]string

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}
	fmt.Println("board: ", board)
	count := 0

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] == "A" {
				if isX_MAS(board, i, j) {
					count++
				}
			}
		}
	}

	fmt.Println("count: ", count)
}
