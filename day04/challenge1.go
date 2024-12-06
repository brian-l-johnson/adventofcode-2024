package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Board[T any] struct {
	grid [][]T
}

func (b *Board[T]) print() {
	for _, row := range b.grid {
		fmt.Println(row)
	}
}
func (b *Board[T]) getNeighbours(row int, col int) []coords {
	var neighbours []coords
	for _, dir := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
		if row+dir[0] >= 0 && row+dir[0] < len(b.grid) && col+dir[1] >= 0 && col+dir[1] < len(b.grid[row]) {
			neighbours = append(neighbours, coords{row: row + dir[0], col: col + dir[1]})
		}
	}
	return neighbours
}

type coords struct {
	row int
	col int
}

func isXMAS(board [][]string, target []string, x int, y int, xdir int, ydir int, pos int) bool {
	if pos == len(target) {
		return true
	}
	if x < 0 || y < 0 || x >= len(board[0]) || y >= len(board) {
		return false
	}
	if board[y][x] != target[pos] {
		return false
	}
	return isXMAS(board, target, x+xdir, y+ydir, xdir, ydir, pos+1)
}

func main() {
	//read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var board [][]string

	target := [4]string{"X", "M", "A", "S"}

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, strings.Split(line, ""))
	}
	fmt.Println("board: ", board)

	b := Board[string]{grid: board}
	count := 0
	for row := 0; row < len(b.grid); row++ {
		for col := 0; col < len(b.grid[row]); col++ {
			for _, dir := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
				if isXMAS(board, target[:], col, row, dir[0], dir[1], 0) {
					count += 1
				}
			}
		}
	}

	fmt.Println("count: ", count)
}
