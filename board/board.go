package board

import "fmt"

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
